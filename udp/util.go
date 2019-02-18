package udp

import (
	"time"
	"net"
	"sync"
)

var (
	bufferPool sync.Pool
)

func init() {
	bufferPool.New = func() interface{} {
		return make([]byte, mtu)
	}
}

func getConnectionID(addr *net.UDPAddr) interface{} {
	if len(addr.IP) == 4 {
		return uint64(addr.IP[3]) | uint64(addr.IP[2])<<8 | uint64(addr.IP[1])<<16 |
			uint64(addr.IP[0])<<24 | uint64(addr.Port)<<32
	}

	return string(append(addr.IP, byte(addr.Port), byte(addr.Port>>8)))
}

func ms() uint32 {
	return uint32(time.Now().UnixNano() / int64(time.Millisecond))
}
