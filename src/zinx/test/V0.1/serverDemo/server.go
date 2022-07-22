package main

import (
	"Zserver/src/zinx/servernet"
	"fmt"
)

func main() {
	s := servernet.NewServer("[V01]")
	fmt.Println(s)
	s.Serve()
}
