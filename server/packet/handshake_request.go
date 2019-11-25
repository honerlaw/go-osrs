package packet

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/model"
	"math/rand"
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

func (h *HandshakeRequest) Handle(client *model.Client) []model.Packet {
	client.Player = model.NewPlayer(h.namehash)
	switch h.opcode {
	case 14:
		client.MoveToNextCodecState()
		return []model.Packet{NewHandshakeResponse(0, rand.Int63())}
	}
	return nil
}

func (h *HandshakeRequest) Encode() *io.Buffer {
	return nil
}