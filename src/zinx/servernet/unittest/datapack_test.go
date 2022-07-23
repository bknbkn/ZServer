package unittest

import (
	"Zserver/src/zinx/servernet"
	"io"
	"log"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatalln("server listen err: ", err)
	}
	/*
		服务端
	*/
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("server accept err ", err)
			}

			go func(conn net.Conn) {
				dp := servernet.NewDataPack()
				log.Println("start connection ")
				for {
					log.Println("read headData")
					headData := make([]byte, dp.GetHeadLen())
					if _, err := io.ReadFull(conn, headData); err != nil {
						log.Fatalln("read head err: ", err)
					}
					msg, err := dp.UnPackHead(headData)
					if err != nil {
						log.Fatalln("unpack err: ", err)
					}
					if msg.GetMsgLen() > 0 {
						log.Println("msg len: ", msg.GetMsgLen())
					}
					data := make([]byte, msg.GetMsgLen())
					if _, err := io.ReadFull(conn, data); err != nil {
						log.Fatalln("read data err: ", err)
					}
					msg.SetData(data)
					log.Printf("Receive MsgID: %v, MsgLen: %d, MsgData: %c\n",
						msg.GetMsgId(), msg.GetMsgLen(), msg.GetData())
				}
			}(conn)
		}
	}()

	/*
		客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatalln("client dial err: ", err)
	}
	dp := servernet.NewDataPack()

	// 模拟粘包
	msg1 := &servernet.Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte("Hello"),
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		log.Fatalln("client pack err: ", err)
	}

	msg2 := &servernet.Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte("world"),
	}

	sendData2, err := dp.Pack(msg2)
	if err != nil {
		log.Fatalln("client pack err: ", err)
	}
	sendData := append(sendData1, sendData2...)
	conn.Write(sendData)
	select {}
}
