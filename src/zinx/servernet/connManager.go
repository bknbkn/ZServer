package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/utils"
	"errors"
	"log"
	"sync"
)

type ConnManager struct {
	connections map[uint32]serverinterface.IConnection
	connLock    sync.RWMutex
	cid         int
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]serverinterface.IConnection),
	}
}

func (c *ConnManager) AddConn(conn serverinterface.IConnection) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if c.Len() >= utils.GlobalConfig.MaxConn {
		conn.Stop()
		return errors.New("too many connections")
	}
	c.connections[conn.GetConnID()] = conn
	log.Println("connection add to ConnManager successfully conn num = ", c.Len())
	return nil
}

func (c *ConnManager) RemoveConn(ConnId uint32) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if _, ok := c.connections[ConnId]; !ok {
		return errors.New("remove connection err: not found")
	}
	delete(c.connections, ConnId)
	log.Println("connection ", ConnId, "remove from ConnManager successfully conn num = ", c.Len())
	return nil
}

func (c *ConnManager) GetConn(ConnId uint32) (serverinterface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[ConnId]; !ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Unlock()
	defer c.connLock.Unlock()

	for connId, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connId)
	}
	log.Println("clear connection manager successfully")
}
