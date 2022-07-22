package main

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
	//fmt.Println("heel")
	var de De
	//dd := DD{}
	//de = &dd
	de = &DD{}
	de.A()
	de.B()
	//http.ListenAndServe()

}
