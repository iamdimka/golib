package tcp

import (
	"bufio"
	"net"
	"sync"
)

type Connection struct {
	sync.Mutex
	id      uint64
	Socket  net.Conn
	Reader  *bufio.Reader
	context map[interface{}]interface{}
}

func FromSocket(socket net.Conn, ai uint64) *Connection {
	return &Connection{
		id:      ai,
		Socket:  socket,
		Reader:  bufio.NewReader(socket),
		context: make(map[interface{}]interface{}),
	}
}

func (c *Connection) ID() uint64 {
	return c.id
}

func (c *Connection) Set(key, value interface{}) {
	c.context[key] = value
}

func (c *Connection) Has(key interface{}) bool {
	_, ok := c.context[key]
	return ok
}

func (c *Connection) Get(key interface{}) interface{} {
	return c.context[key]
}

func (c *Connection) Del(key interface{}) {
	delete(c.context, key)
}
