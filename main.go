package main

import (
	"Zserver/test"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	user := test.User{}
	user.Address = "sadsda"
	user.Password = "123456"
	user.UserName = "仙士可"
	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))
	//序列化user结构体数据
	out, err := proto.Marshal(&user)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	//反序列化user结构体
	user2 := test.User{}
	err = proto.Unmarshal(out, &user2)
	if err != nil {
		log.Fatalln("Failed to parse address User:", err)
	}
	bytes, _ = json.Marshal(user2)
	fmt.Println(string(bytes))
}
