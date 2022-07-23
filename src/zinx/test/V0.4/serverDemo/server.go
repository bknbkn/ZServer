package main

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/servernet"
	"fmt"
	"log"
	"os"
)

type PingRouter struct {
	servernet.BaseRouter
}

func (p *PingRouter) BeforeHandle(request serverinterface.IRequest) {
	log.Println("Call router BeforeHandle")
	if _, err := request.
		GetConnection().
		GetTCPConnection().
		Write([]byte("before ping...\n")); err != nil {
		log.Println("call before ping err: ", err)
	}

}

func (p *PingRouter) Handle(request serverinterface.IRequest) {
	log.Println("Call router Handle")
	if _, err := request.
		GetConnection().
		GetTCPConnection().
		Write([]byte("call ping...\n")); err != nil {
		log.Println("call ping err: ", err)
	}
}

func (p *PingRouter) AfterHandle(request serverinterface.IRequest) {
	log.Println("Call router AfterHandle")
	if _, err := request.
		GetConnection().
		GetTCPConnection().
		Write([]byte("call after ping...\n")); err != nil {
		log.Println("call after ping err: ", err)
	}
}

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
	s := servernet.NewServer("[V03]")
	fmt.Println(s)

	s.AddRouter(&PingRouter{})
	s.Serve()
}
