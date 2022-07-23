package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/utils"
	"fmt"
	"log"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    serverinterface.IRouter
}

func (server *Server) Start() {
	log.Printf("[Start] Server %s, IPVersion %s, Listening at %s: %d, is starting\n",
		server.Name, server.IPVersion, server.IP, server.Port)
	log.Printf("[Start] MaxConnection: %v, MaxPackageSize: %v\n",
		utils.GlobalConfig.MaxConn, utils.GlobalConfig.MaxPackageSize)

	go func() {
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
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("AcceptTCP err: ", err)
				continue
			}
			var cid uint32 = 0
			dealConn := NewConnection(conn, cid, server.Router)
			cid++
			dealConn.Start()
		}
	}()
}

func (server *Server) Stop() {

}

func (server *Server) Serve() {
	server.Start()
	select {}
}

func (server *Server) AddRouter(router serverinterface.IRouter) {
	server.Router = router
	log.Printf("Server %v Add router succ\n", server.Name)
}

func NewServer(name string) serverinterface.IServer {
	return &Server{
		Name:      utils.GlobalConfig.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalConfig.Host,
		Port:      utils.GlobalConfig.TcpPort,
		Router:    nil,
	}
}
