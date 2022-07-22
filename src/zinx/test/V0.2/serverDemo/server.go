package main

import (
	"Zserver/src/zinx/servernet"
	"fmt"
)

func main() {
	s := servernet.NewServer("[V02]")
	fmt.Println(s)
	s.Serve()
}
