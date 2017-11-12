package udp

// Packet
type Packet struct {
	Protocol    uint16 // 0 bit for syn
	Sequence    uint32
	Ack         uint32
	AckPrevious uint32
}
