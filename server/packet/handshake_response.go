package packet

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/model"
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

func (h *HandshakeResponse) Handle(client *model.Client) []model.Packet {
	return nil
}

func (h *HandshakeResponse) Encode() *io.Buffer {
	var buf = io.NewBuffer(17)
	buf.WriteLong(h.random)
	buf.WriteByte(h.status)
	buf.WriteLong(h.serverIsaacKey)
	return buf
}
