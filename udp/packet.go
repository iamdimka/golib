package udp

// Packet
type Packet []byte

func NewPacket(protocol uint16) Packet {
	packet := make(Packet, headerPacketSize)

	packet[0] = byte(protocol)
	packet[1] = byte(protocol >> 8)

	return packet
}

func (p Packet) Protocol() uint16 {
	return uint16(p[0]) | uint16(p[1])<<8
}

func (p Packet) Sequence() uint32 {
	return uint32(p[2]) | uint32(p[3])<<8 | uint32(p[4])<<16 | uint32(p[5])<<24
}

func (p Packet) SetSequence(sequence uint32) {
	p[2] = byte(sequence)
	p[3] = byte(sequence >> 8)
	p[4] = byte(sequence >> 16)
	p[5] = byte(sequence >> 24)
}

func (p Packet) Ack() uint32 {
	return uint32(p[6]) | uint32(p[7])<<8 | uint32(p[8])<<16 | uint32(p[9])<<24
}

func (p Packet) SetAck(ack uint32) {
	p[6] = byte(ack)
	p[7] = byte(ack >> 8)
	p[8] = byte(ack >> 16)
	p[9] = byte(ack >> 24)
}

func (p Packet) AckBits() uint32 {
	return uint32(p[10]) | uint32(p[11])<<8 | uint32(p[12])<<16 | uint32(p[13])<<24
}

func (p Packet) SetAckBits(ack uint32) {
	p[10] = byte(ack)
	p[11] = byte(ack >> 8)
	p[12] = byte(ack >> 16)
	p[13] = byte(ack >> 24)
}

func (p Packet) Payload() []byte {
	return p[headerPacketSize:]
}

func (p Packet) HeaderSize() int {
	return headerPacketSize
}

func (p Packet) PayloadSize() int {
	return len(p) - headerPacketSize
}

func (p Packet) SetPayload(data []byte) []byte {
	return append(p[:headerPacketSize], data...)
}
