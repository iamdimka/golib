package udp

import (
	"encoding/binary"
	"math"
	"net"
	"sync"
)

const (
	flagReset byte = 1 << iota
	flagChunked
	flagClose
)

const (
	ackBits = 32

	mtu           = 1000
	headerSize    = 9
	chunkDataSize = mtu - headerSize - 4
)

var (
	Protocol uint32 = 0xfaf1fbfc
)

type Connection struct {
	sync.Mutex
	conn   net.PacketConn
	remote *net.UDPAddr

	seq    uint32
	offset uint32
	buf    []byte

	outFirst uint32
	outQueue [][]byte

	inFirst uint32 // next accept seq
	inLast  uint32 // next accept seq + ackBits
	inQueue [ackBits][]byte

	chData  chan []byte
	chClose chan struct{}
}

func NewConnection(addr *net.UDPAddr) *Connection {
	return &Connection{
		remote:  addr,
		inFirst: 1,
		chData:  make(chan []byte, 1),
		chClose: make(chan struct{}, 1),
	}
}

func (c *Connection) receive(b []byte) {
	if len(b) < headerSize {
		return
	}

	if !checkProtocol(b) {
		return
	}

	id := binary.LittleEndian.Uint32(b[4:])
	if id < c.inFirst || id > c.inFirst+ackBits {
		return
	}

	flag := b[8]
	ack := binary.LittleEndian.Uint32(b[9:])
	if c.outFirst > 0 && ack >= c.outFirst {
		left := ack - c.outFirst
		c.outQueue = c.outQueue[left+1:]
		c.outFirst = ack + 1
	}

	if id > c.inLast {
		c.inLast = id
	}

	c.inQueue[id-c.inFirst] = b

	var done byte

	for i := c.inLast - c.inFirst; i >= 0; i-- {
		chunk := c.inQueue[done]
		if chunk == nil {
			break
		}

		if flag&flagChunked == flagChunked {
			if c.offset == 0 {
				size := binary.LittleEndian.Uint32(chunk[9:])
				c.buf = make([]byte, size)
			}

			c.offset += uint32(copy(c.buf[c.offset:], chunk[13:]))
		} else if c.offset != 0 {
			copy(c.buf[c.offset:], chunk[13:])
			buf := c.buf
			c.buf = nil
			c.offset = 0
			c.chData <- buf
		} else {
			c.chData <- chunk[9:]
		}

		done++
	}

	if done > 0 {
		c.inFirst += uint32(done)
		moveLeft(c.inQueue, done)
	}
}

func (c *Connection) sender() {
	c.Lock()

	for i := 0; i < len(c.outQueue); i++ {
		c.conn.WriteTo(c.outQueue[i], c.remote)
	}

	c.Unlock()
}

func (c *Connection) Send(b []byte) {
	c.Lock()
	l := len(b)

	if l <= chunkDataSize {
		c.seq++
		p := bufferPool.Get().([]byte)[:l+headerSize]
		binary.LittleEndian.PutUint32(p, Protocol)
		binary.LittleEndian.PutUint32(p[4:], c.seq)
		copy(p[headerSize:], b)
		c.outQueue = append(c.outQueue, p)
	} else {
		chunks := int(math.Ceil(float64(l) / chunkDataSize))

		offset := 0
		end := chunkDataSize
		for {
			c.seq++
			p := bufferPool.Get().([]byte)[:l+headerSize]
			binary.LittleEndian.PutUint32(p, Protocol)
			binary.LittleEndian.PutUint32(p[4:], c.seq)
			binary.LittleEndian.PutUint32(p[9:], uint32(l))
			chunks--
			if chunks == 0 {
				p[8] = 0
				copy(p[13:], b[offset:l])
				c.outQueue = append(c.outQueue, p)
				break
			}

			p[8] = flagChunked
			copy(p[13:], b[offset:end])
			c.outQueue = append(c.outQueue, p)
			offset = end
			end += chunkDataSize
		}
	}

	c.Unlock()
}

func moveLeft(b [ackBits][]byte, size byte) {
	for i := size; i < ackBits; i++ {
		b[i-size] = b[i]
		b[i] = nil
	}
}

// func (c *Connection){}

func checkProtocol(b []byte) bool {
	return b[0] == byte(Protocol) && b[1] == byte(Protocol>>8) && b[2] == byte(Protocol>>16) && b[3] == byte(Protocol>>24)
}
