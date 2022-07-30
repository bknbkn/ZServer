package main

import (
	"Zserver/src/zinx/test/protobufTest/pb"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	person := &pb.Person{
		Name:   "BKN",
		Age:    100,
		Emails: []string{"dsad@gmail.com", "fsaf@gmail.com"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "1232",
				Type:   pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number: "14214",
				Type:   pb.PhoneType_MOBILE,
			},
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalln("proto marshal err: ", err)
	}

	newData := &pb.Person{}
	err = proto.Unmarshal(data, newData)
	if err != nil {
		log.Fatalln("unmarshl err: ", err)
	}

	log.Println("raw: ", person)
	log.Println("after: ", newData)
}
