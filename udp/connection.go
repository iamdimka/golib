package udp

import "net"
import "sync"

const (
	allAccepted = ^uint32(0)
)

var (
	BufferSize = 1000
)

type Connection struct {
	sync.Mutex
	Addr           *net.UDPAddr
	server         *Server
	outSeq         uint32
	acceptedOutSeq uint32
	inSeq          uint32
	in32           uint32     //last32 received messages
	out32          [32][]byte //last32 sent messages
	pending32      [32][]byte // 0 is inSeq, 1 = inSeq-1
}

func newConnection(addr *net.UDPAddr, server *Server) *Connection {
	return &Connection{
		Addr:   addr,
		server: server,
	}
}

func (c *Connection) prepend(payload []byte) {
	for i := 0; i < 31; i++ {
		c.out32[i+1] = c.out32[i]
	}

	c.out32[0] = payload
}

func (c *Connection) Receive(p Packet) {
	seq := p.Sequence()
	diff := seq - c.inSeq

	if seq <= c.inSeq {
		return
	}

	if diff > 31 {
		return
	}

}

func (c *Connection) Send(data []byte) error {
	c.Lock()
	c.outSeq++
	id := c.outSeq
	p := NewPacket(c.server.protocol)
	p.SetSequence(id)
	p.SetAck(c.inSeq)
	p.SetAckBits(c.in32)
	b := p.SetPayload(data)
	c.prepend(b)
	c.Unlock()

	return c.server.Send(c.Addr, b)
}
