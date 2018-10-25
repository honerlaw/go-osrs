package packet

import (
	"osrs/io"
)

type PacketStateDecoder func(*io.Buffer, *io.Client) ([]Packet, error)
type PacketStateEncoder func(*io.Buffer, *io.Client) *io.Buffer // used for further encoding of a packet

type PacketDecoder func(uint32, int8, *io.Buffer) []Packet

type Packet interface {
	Handle(*io.Client) []Packet
	Encode() *io.Buffer
}

const (
	PACKET_STATE_HANDSHAKE = 0
	PACKET_STATE_LOGIN     = 1
	PACKET_STATE_GAME      = 2
)
