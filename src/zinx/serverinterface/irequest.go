package serverinterface

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	IMessage
}
