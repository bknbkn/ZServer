package servernet

import "Zserver/src/zinx/serverinterface"

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
