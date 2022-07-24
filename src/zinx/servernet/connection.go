package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/utils"
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

type Connection struct {
	TcpServer     serverinterface.IServer
	Conn          *net.TCPConn
	ConnID        uint32
	isClosed      bool
	ExitChan      chan bool
	MessageHandle serverinterface.IMessageHandle

	// 读写goroutine之间通信
	msgChan chan []byte

	// 链接属性
	property     map[string]interface{}
	propertyLock sync.RWMutex
}

func NewConnection(server serverinterface.IServer, conn *net.TCPConn, connID uint32, handle serverinterface.IMessageHandle) *Connection {
	c := &Connection{
		Conn:          conn,
		ConnID:        connID,
		MessageHandle: handle,
		isClosed:      false,
		ExitChan:      make(chan bool, 1),
		msgChan:       make(chan []byte),
		TcpServer:     server,
		property:      make(map[string]interface{}),
	}
	if err := c.TcpServer.GetConnMgr().AddConn(c); err != nil {
		log.Println("Add connection err: ", err)
		return nil
	}
	return c
}

func (conn *Connection) StartReader() {
	log.Println("[Reader Goroutine is Running...]")
	defer log.Printf("[Conn %v reader exit, remote addr is %v]\n",
		conn.ConnID, conn.RemoteAddr())
	defer conn.Stop()

	for {
		dp := NewDataPack()
		msg, err := dp.UnpackRead(conn.GetTCPConnection())
		if err != nil {
			log.Println("read conn msg err: ", err)
			break
		}
		req := &Request{
			conn:     conn,
			IMessage: msg,
		}
		// 建立新的goroutine，并行处理耗时的任务，这里由于是TCP链接，只负责数据传输
		// 不需要像HTTP1.1一样等待上一个请求响应后，才能处理下一个请求
		// V0.8之后交给工作池处理
		if utils.GlobalConfig.WorkerPoolSize > 0 {
			conn.MessageHandle.AddMessageToQueue(req)
		} else {
			go conn.MessageHandle.DoMsgHandler(req)
		}

	}
}

func (conn *Connection) StartWriter() {
	log.Println("[Writer Goroutine is Running...]")
	defer log.Printf("[Conn %v writer exit, remote addr is %v]\n",
		conn.ConnID, conn.RemoteAddr())

	for {
		select {
		case data := <-conn.msgChan:
			if _, err := conn.Conn.Write(data); err != nil {
				log.Println("send conn msg err: ", err)
				break
			}
		case <-conn.ExitChan:
			return
		}
	}
}

func (conn *Connection) Start() {
	log.Println("Conn Start.... ID is :", conn.ConnID)
	go conn.StartReader()
	go conn.StartWriter()
	conn.TcpServer.CallOnConnStart(conn)
}
func (conn *Connection) Stop() {
	if conn.isClosed {
		log.Printf("%v has Stopped\n", conn.ConnID)
		return
	}
	conn.isClosed = true
	log.Println("Connection Stop, ConnID :", conn.ConnID)
	if err := conn.TcpServer.GetConnMgr().RemoveConn(conn.ConnID); err != nil {
		log.Println("remove connection err: ", err)
	}
	conn.TcpServer.CallOnConnStop(conn)
	for err := conn.Conn.Close(); err != nil; err = conn.Conn.Close() {
		log.Printf("Close Conn %v failed err : %v, try again...", conn.ConnID, err)
		time.Sleep(time.Second)
	}
	close(conn.ExitChan)
	close(conn.msgChan)
}
func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}
func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}
func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}
func (conn *Connection) SendMsg(msgId uint32, data []byte) error {
	if conn.isClosed {
		return errors.New("connection is closed when send message")
	}

	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		log.Println("Pack err, msgId :", msgId, err)
		return errors.New("pack msg err")
	}

	//if _, err := conn.GetTCPConnection().Write(binaryMsg); err != nil {
	//	log.Println("Write msg err, msgId: ", msgId, err)
	//	return errors.New("write msg err")
	//}
	conn.msgChan <- binaryMsg
	return nil
}

func (conn *Connection) SetProperty(key string, value interface{}) {
	conn.propertyLock.Lock()
	defer conn.propertyLock.Unlock()
	conn.property[key] = value
}
func (conn *Connection) GetProperty(key string) (interface{}, error) {
	conn.propertyLock.RLock()
	defer conn.propertyLock.RUnlock()
	if value, ok := conn.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found, get fail")
	}

}
func (conn *Connection) RemoveProperty(key string) error {
	conn.propertyLock.Lock()
	defer conn.propertyLock.Unlock()
	if _, ok := conn.property[key]; ok {
		delete(conn.property, key)
		return nil
	} else {
		return errors.New("no property found, remove fail")
	}
}
