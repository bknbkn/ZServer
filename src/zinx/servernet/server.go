package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"errors"
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

func ClientHandler(conn *net.TCPConn, data []byte, cnt int) error {
	log.Println("[Conn Handle]....")
	if _, err := conn.Write(data[:cnt]); err != nil {
		log.Println("Write back err: ", err)
		return errors.New("call back to client error")
	}
	return nil
}

func (server *Server) Start() {
	log.Printf("[Start] Server %s, IPVersion %s, Listening at %s: %d, is starting",
		server.Name, server.IPVersion, server.IP, server.Port)
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
			dealConn := NewConnection(conn, cid, ClientHandler)
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

func (server *Server) AddRouter() {

}

func NewServer(name string) serverinterface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9999,
	}
}