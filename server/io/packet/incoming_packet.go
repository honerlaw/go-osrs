package packet

import (
	"github.com/honerlaw/go-osrs/io"
)

type IncomingPacket interface {
	Decode(buffer *io.Buffer, client *io.Client, opcode byte, length int8) ([]PacketEvent, error)
}
