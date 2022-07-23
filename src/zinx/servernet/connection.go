package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/utils"
	"log"
	"net"
	"time"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	ExitChan chan bool
	Router   serverinterface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router serverinterface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
}

func (conn *Connection) StartReader() {
	log.Println("Reader is running...")
	defer log.Printf("Conn %v reader exit, remote addr is %v\n",
		conn.ConnID, conn.RemoteAddr())
	defer conn.Stop()

	for {
		buf := make([]byte, utils.GlobalConfig.MaxPackageSize)
		cnt, err := conn.Conn.Read(buf)
		if err != nil {
			log.Println("recv buf err :", err)
			continue
		}

		req := Request{
			conn: conn,
			data: buf[:cnt],
		}

		// 建立新的goroutine，并行处理耗时的任务，这里由于是TCP链接，只负责数据传输
		// 不需要像HTTP1.1一样等待上一个请求响应后，才能处理下一个请求
		go func(request serverinterface.IRequest) {
			conn.Router.BeforeHandle(&req)
			conn.Router.Handle(&req)
			conn.Router.AfterHandle(&req)
		}(&req)
		//if err := conn.handler(conn.Conn, buf, cnt); err != nil {
		//	log.Printf("Conn %v handle is err: %v\n", conn.ConnID, err)
		//	continue
		//}
	}
}

func (conn *Connection) StartWriter() {

}

func (conn *Connection) Start() {
	log.Println("Conn Start.... ID is :", conn.ConnID)
	go conn.StartReader()
}
func (conn *Connection) Stop() {
	if conn.isClosed {
		log.Printf("%v has Stopped\n", conn.ConnID)
		return
	}
	log.Println("Connection Stop, ConnID :", conn.ConnID)
	conn.isClosed = true
	for err := conn.Conn.Close(); err != nil; err = conn.Conn.Close() {
		log.Printf("Close Conn %v failed err : %v, try again...", conn.ConnID, err)
		time.Sleep(time.Second)
	}
	close(conn.ExitChan)
}
func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}
func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}
func (conn *Connection) RemoteAddr() net.Addr {
	return nil
}
func (conn *Connection) Send(data []byte) error {
	return nil
}
