package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/utils"
	"fmt"
	"log"
	"net"
)

type Server struct {
	Name          string
	IPVersion     string
	IP            string
	Port          int
	MessageHandle serverinterface.IMessageHandle
	ConnManager   serverinterface.IConnManager

	OnConnStart func(conn serverinterface.IConnection)
	OnConnStop  func(conn serverinterface.IConnection)
}

func (server *Server) Start() {
	log.Printf("[Start] Server %s, IPVersion %s, Listening at %s: %d, is starting\n",
		server.Name, server.IPVersion, server.IP, server.Port)
	log.Printf("[Start] MaxConnection: %v, MaxPackageSize: %v\n",
		utils.GlobalConfig.MaxConn, utils.GlobalConfig.MaxPackageSize)

	go func() {
		// 0) Start Worker Pool
		server.MessageHandle.StartWorkerPool()

		// 1) Get TCP Addr
		addr, err := net.ResolveTCPAddr(server.IPVersion,
			fmt.Sprintf("%s:%d", server.IP, server.Port))
		if err != nil {
			log.Println("Resolve tcp addr error: ", err)
			return
		}

		// 2) Listen Server Addr
		listener, err := net.ListenTCP(server.IPVersion, addr)
		if err != nil {
			log.Printf("Get %s listener error: %v\n", server.IPVersion, err)
			return
		}
		log.Printf("Start Server %s successfully. Now listening...\n", server.Name)

		// 3) Block for client connect
		var cid uint32 = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("AcceptTCP err: ", err)
				continue
			}
			if server.ConnManager.Len() >= utils.GlobalConfig.MaxConn {
				log.Println("too many connection, max = ", utils.GlobalConfig.MaxConn)
				if err := conn.Close(); err != nil {
					log.Println("close overflow conn err: ", err)
				}
				continue
			}
			dealConn := NewConnection(server, conn, cid, server.MessageHandle)
			cid++
			dealConn.Start()
		}
	}()
}

func (server *Server) Stop() {
	log.Println("[STOP] Server: ", server.Name)
	server.ConnManager.ClearConn()
}

func (server *Server) Serve() {
	server.Start()
	select {}
}

func (server *Server) AddRouter(msgId uint32, router serverinterface.IRouter) {
	server.MessageHandle.AddRouter(msgId, router)
	log.Printf("Server %v Add router succ\n", server.Name)
}

func (server *Server) GetConnMgr() serverinterface.IConnManager {
	return server.ConnManager
}
func (server *Server) SetOnConnStart(hook func(serverinterface.IConnection)) {
	server.OnConnStart = hook
}
func (server *Server) SetOnConnStop(hook func(serverinterface.IConnection)) {
	server.OnConnStop = hook
}
func (server *Server) CallOnConnStart(conn serverinterface.IConnection) {
	log.Println("----> call onConnStart()")
	server.OnConnStart(conn)
}
func (server *Server) CallOnConnStop(conn serverinterface.IConnection) {
	log.Println("----> call onConnStop()")
	server.OnConnStop(conn)
}
func NewServer(name string) serverinterface.IServer {
	return &Server{
		Name:          utils.GlobalConfig.Name,
		IPVersion:     "tcp4",
		IP:            utils.GlobalConfig.Host,
		Port:          utils.GlobalConfig.TcpPort,
		MessageHandle: NewMessageHandle(),
		ConnManager:   NewConnManager(),
		OnConnStart:   func(conn serverinterface.IConnection) {},
		OnConnStop:    func(conn serverinterface.IConnection) {},
	}
}
