package main

import (
	"Zserver/src/zinx/servernet"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("Client Start...")
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Println("Client err: ", err)
		return
	}
	for {
		dp := servernet.NewDataPack()
		binMsg, err := dp.Pack(servernet.NewMessage(0, []byte("V0.5 test0")))
		if err != nil {
			log.Fatalln("pack err: ", err)
		}
		if _, err := conn.Write(binMsg); err != nil {
			log.Fatalln("write msg err: ", err)
		}

		msg, err := dp.UnpackRead(conn.(*net.TCPConn))
		log.Printf("msg id = %v, msg length = %v, msg data = %v",
			msg.GetMsgId(), msg.GetMsgLen(), string(msg.GetData()))
		time.Sleep(time.Second)
	}
}
