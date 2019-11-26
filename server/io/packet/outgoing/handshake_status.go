package outgoing

import (
	"github.com/honerlaw/go-osrs/io"
)

type HandshakeStatus struct {
	random         int64
	status         byte
	serverIsaacKey int64
}

func NewClientStatus(status byte, serverIsaacKey int64) *HandshakeStatus {
	return &HandshakeStatus{
		random:         0,
		status:         status,
		serverIsaacKey: serverIsaacKey,
	}
}

func (h *HandshakeStatus) Encode(_ *io.Client) *io.Buffer {
	var buf = io.NewBuffer(17)
	buf.WriteLong(h.random)
	buf.WriteByte(h.status)
	buf.WriteLong(h.serverIsaacKey)
	return buf
}
