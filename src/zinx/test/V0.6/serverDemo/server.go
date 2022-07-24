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

func (p *PingRouter) Handle(request serverinterface.IRequest) {
	log.Println("Call router Handle")
	log.Println("recv from client msgId = ", request.GetMsgId(),
		"msg len = ", request.GetMsgLen(), "data = ", string(request.GetData()))

	if err := request.GetConnection().SendMsg(1, []byte("ping")); err != nil {
		log.Println("send msg err: ", err)
	}
}

type HelloRouter struct {
	servernet.BaseRouter
}

func (p *HelloRouter) Handle(request serverinterface.IRequest) {
	log.Println("Call hello Handle")
	log.Println("recv from client msgId = ", request.GetMsgId(),
		"msg len = ", request.GetMsgLen(), "data = ", string(request.GetData()))

	if err := request.GetConnection().SendMsg(2, []byte("hello")); err != nil {
		log.Println("send msg err: ", err)
	}
}

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
	s := servernet.NewServer("[V07]")
	fmt.Println(s)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
