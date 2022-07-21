package serverInterface

type IServer interface {
	Start()
	Stop()
	Serve()
}
