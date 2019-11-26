package outgoing

import (
	"github.com/honerlaw/go-osrs/io"
)

type LoginStatus struct {
	status    byte
	privilege byte
}

func NewLoginStatus(status byte, privilege byte) *LoginStatus {
	return &LoginStatus{
		status:    status,
		privilege: privilege,
	}
}

func (h *LoginStatus) Encode(_ *io.Client) *io.Buffer {
	var buf = io.NewBuffer(3)
	buf.WriteByte(h.status)
	buf.WriteByte(h.privilege)
	buf.WriteByte(0)
	return buf
}
