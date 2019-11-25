package packet

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/model"
)

type LoginResponse struct {
	status    byte
	privilege byte
}

func NewLoginResponse(status byte, privilege byte) *LoginResponse {
	return &LoginResponse{
		status:    status,
		privilege: privilege,
	}
}

func (h *LoginResponse) Handle(client *model.Client) []model.Packet {
	return nil
}

func (h *LoginResponse) Encode() *io.Buffer {
	var buf = io.NewBuffer(3)
	buf.WriteByte(h.status)
	buf.WriteByte(h.privilege)
	buf.WriteByte(0)
	return buf
}
