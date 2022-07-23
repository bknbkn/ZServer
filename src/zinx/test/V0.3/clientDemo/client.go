package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("Client Start...")
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp4", "0.0.0.0:9999")
	if err != nil {
		log.Println("Client err: ", err)
		return
	}
	for {
		_, err := conn.Write([]byte("Hello World..."))
		if err != nil {
			log.Println("Write conn err: ", err)
			continue
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Println("Read conn err: ", err)
			continue
		}
		log.Printf("Server send is %s\n", buf[:cnt])
		time.Sleep(time.Second)
	}
}
