package serverinterface

type IConnManager interface {
	AddConn(conn IConnection) error
	RemoveConn(ConnId uint32) error
	GetConn(ConnId uint32) (IConnection, error)
	Len() int
	ClearConn()
}
