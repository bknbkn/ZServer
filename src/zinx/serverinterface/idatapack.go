package serverinterface

type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	UnPackHead([]byte) (IMessage, error)
}
