package serverNet

import (
	"Zserver/src/zinx/serverInterface"
	"fmt"
	"log"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IPString  string
	Port      int
}

func (server *Server) Start() {
	log.Printf("[Start] Server %s, IPVersion %s, Listening at %s: %d, is starting",
		server.Name, server.IPVersion, server.IPString, server.Port)
	go func() {
		// 1) Get TCP Addr
		addr, err := net.ResolveTCPAddr(server.IPVersion,
			fmt.Sprintf("%s:%d", server.IPString, server.Port))
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
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						log.Println("Receive client stream err: ", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						log.Println("Write back to client err: ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (server *Server) Stop() {

}

func (server *Server) Serve() {
	server.Start()
	select {}
}

func NewServer(name string) serverInterface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IPString:  "0.0.0.0",
		Port:      9999,
	}
}
