package udp

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Server struct {
	protocol uint16
	connections map[string]
}

func NewServer(protocol uint16) *Server {
	return &Server{
		protocol: protocol,
	}
}

func (s *Server) Listen(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP(udpAddr.Network(), udpAddr)
	if err != nil {
		return err
	}

	go s.handleConnections(conn)
	return nil
}

func (s *Server) handleConnections(conn *net.UDPConn) {

	var (
		protocol uint16
		sequence uint32
		ack uint32
		buf      = make([]byte, 1000)
		reader   = bytes.NewReader(buf)
		l        int
		from     *net.UDPAddr
		err      error
	)

	for {
		l, from, err = conn.ReadFromUDP(buf)

		if err != nil {
			break
		}

		if l < 14 {
			continue
		}

		protocol = binary.LittleEndian.Uint16(buf)
		sequence = binary.LittleEndian.Uint32(buf[2:])
		ack = binary.LittleEndian.Uint32(buf[4:])
		
		reader.Reset(buf)
		binary.Read(reader, binary.LittleEndian, &protocol)
		if protocol != s.protocol {
			continue
		}


	}
}
