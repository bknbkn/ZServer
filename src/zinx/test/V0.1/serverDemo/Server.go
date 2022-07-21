package main

import (
	"Zserver/src/zinx/serverNet"
	"fmt"
)

func main() {
	s := serverNet.NewServer("V01")
	fmt.Println(s)
	s.Serve()
}
