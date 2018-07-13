package tcp

import (
	"log"
	"net"
	"sync"
	"time"
)

const maxCap = 1024

type Handler interface {
	HandleTCP(net.Conn)
}

type HandlerFunc func(net.Conn)

func (m HandlerFunc) HandleTCP(c net.Conn) {
	m(c)
}

type Server struct {
	sync.Mutex
	Listener net.Listener
	handler  func(net.Conn)
}

func New(handler Handler) *Server {
	return &Server{
		handler: handler.HandleTCP,
	}
}

func (s *Server) Listen(addr string) error {
	ln, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}

	var tempDelay time.Duration
	s.Listener = ln

	for {
		socket, e := ln.Accept()

		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}

				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}

				log.Printf("Server: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}

			return e
		}

		go s.handler(socket)
	}
}
