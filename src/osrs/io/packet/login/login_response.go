package login

import (
	"osrs/io"
	"osrs/io/packet"
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

func (h *LoginResponse) Handle(client *io.Client) []packet.Packet {
	return nil
}

func (h *LoginResponse) Encode() *io.Buffer {
	var buf = io.NewBuffer(3)
	buf.WriteByte(h.status)
	buf.WriteByte(h.privilege)
	buf.WriteByte(0)
	return buf
}
