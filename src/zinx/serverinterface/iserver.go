package serverinterface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter()
}
