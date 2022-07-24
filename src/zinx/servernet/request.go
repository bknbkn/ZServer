package servernet

import "Zserver/src/zinx/serverinterface"

// Request 绑定数据和链接上下文
type Request struct {
	conn serverinterface.IConnection
	serverinterface.IMessage
}

func (r *Request) GetConnection() serverinterface.IConnection {
	return r.conn
}
