package handshake

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"math/rand"
	"github.com/honerlaw/go-osrs/model"
)

type HandshakeRequest struct {
	opcode   byte
	namehash byte
}

func NewHandshakeRequest(opcode byte, namehash byte) *HandshakeRequest {
	return &HandshakeRequest{
		opcode,
		namehash,
	}
}

func (h *HandshakeRequest) Handle(client *io.Client) []packet.Packet {
	client.Player = model.NewPlayer(h.namehash)
	switch h.opcode {
	case 14:
		client.State = packet.PACKET_STATE_LOGIN
		return []packet.Packet{NewHandshakeResponse(0, rand.Int63())}
	}
	return nil
}

func (h *HandshakeRequest) Encode() *io.Buffer {
	return nil
}