package handshake

import (
	"osrs/io"
	"osrs/io/packet"
)

type HandshakeResponse struct {
	random         int64
	status         byte
	serverIsaacKey int64
}

func NewHandshakeResponse(status byte, serverIsaacKey int64) *HandshakeResponse {
	return &HandshakeResponse{
		random:         0,
		status:         status,
		serverIsaacKey: serverIsaacKey,
	}
}

func (h *HandshakeResponse) Handle(client *io.Client) []packet.Packet {
	return nil
}

func (h *HandshakeResponse) Encode() *io.Buffer {
	var buf = io.NewBuffer(17)
	buf.WriteLong(h.random)
	buf.WriteByte(h.status)
	buf.WriteLong(h.serverIsaacKey)
	return buf
}
