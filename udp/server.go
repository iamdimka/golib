package udp

import (
	"bytes"
	"encoding/binary"
	"net"
	"time"
)

const headerPacketSize = 14

type temporary interface {
	Temporary() bool
}

type Server struct {
	protocol    uint16
	connections map[interface{}]*Connection
	conn        *net.UDPConn
}

func NewServer(protocol uint16) *Server {
	return &Server{
		protocol:    protocol,
		connections: make(map[interface{}]*Connection),
	}
}

func (s *Server) Listen(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	s.conn, err = net.ListenUDP(udpAddr.Network(), udpAddr)
	if err != nil {
		return err
	}

	go s.handleConnections()
	return nil
}

func (s *Server) Send(addr *net.UDPAddr, data []byte) (err error) {
	_, err = s.conn.WriteTo(data, addr)
	return
}

func (s *Server) handleConnections() {
	var (
		id         interface{}
		sequence   uint32
		ack        uint32
		connection *Connection
		packet     Packet
		buf        = make([]byte, BufferSize)
		reader     = bytes.NewReader(buf)
		l          int
		from       *net.UDPAddr
		err        error
		ok         bool
	)

	for {
		l, from, err = s.conn.ReadFromUDP(buf)

		if err != nil {
			if err, ok := err.(temporary); ok {
				if err.Temporary() {
					time.Sleep(time.Millisecond * 5)
					continue
				}
			}

			break
		}

		if l < headerPacketSize {
			continue
		}

		packet = Packet(buf[:l])

		if packet.Protocol() != s.protocol {
			continue
		}

		id = getConnectionID(from)
		connection, ok = s.connections[id]
		if !ok {
			connection = incomingConnection(from, s)
			s.connections[id] = connection
		}

		reader.Reset(buf)
		binary.Read(reader, binary.LittleEndian, &protocol)
	}
}
