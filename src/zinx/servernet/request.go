package servernet

import "Zserver/src/zinx/serverinterface"

// Request 绑定数据和链接上下文
type Request struct {
	conn serverinterface.IConnection
	data []byte
}

func (r *Request) GetConnection() serverinterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
