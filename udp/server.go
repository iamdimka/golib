package udp

import (
	"bytes"
	"encoding/binary"
	"net"
)

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
			break
		}

		if l < 14 {
			continue
		}

		packet = Packet(buf[:l])

		if packet.Protocol() != s.protocol {
			continue
		}

		id = getConnectionID(from)
		connection, ok = s.connections[id]
		if !ok {
			connection = newConnection(from, s)
			s.connections[id] = connection
		}

		//ackbitfields = [10,11,12,13]

		reader.Reset(buf)
		binary.Read(reader, binary.LittleEndian, &protocol)
	}
}

func getConnectionID(addr *net.UDPAddr) interface{} {
	if len(addr.IP) == 4 {
		return uint64(addr.IP[3]) | uint64(addr.IP[2])<<8 | uint64(addr.IP[1])<<16 |
			uint64(addr.IP[0])<<24 | uint64(addr.Port)<<32
	}

	return string(append(addr.IP, byte(addr.Port), byte(addr.Port>>8)))
}
