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

func ConnStart(conn serverinterface.IConnection) {

	if err := conn.SendMsg(200, []byte("on start")); err != nil {
		log.Println("send err: ", err)
	}

	log.Println("set conn property")
	conn.SetProperty("Name", "aaa")
	conn.SetProperty("Home", "bbb")
}

func ConnStop(conn serverinterface.IConnection) {
	log.Println("====> connection lost Id = ", conn.GetConnID())
	log.Println("get conn property")
	if name, err := conn.GetProperty("Name"); err == nil {
		log.Println("Name ", name)
	}
	if home, err := conn.GetProperty("Home"); err == nil {
		log.Println("Home ", home)
	}
}

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
	s := servernet.NewServer("[V06]")
	fmt.Println(s)

	s.SetOnConnStart(ConnStart)
	s.SetOnConnStop(ConnStop)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	s.Serve()
}
