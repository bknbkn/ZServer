package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type De interface {
	A()
	B()
}
type DD struct {
}

func (dd *DD) A() {

}

func (dd DD) B() {

}

//func NewA() De {
//	return DD{}
//}
func main() {
	var a uint32 = 18
	b := bytes.NewBuffer([]byte{})

	binary.Write(b, binary.BigEndian, a)
	fmt.Printf("%b", b.Bytes())

}
