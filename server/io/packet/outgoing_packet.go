package packet

import (
	"github.com/honerlaw/go-osrs/io"
)

type OutgoingPacket interface {
	Encode(client *io.Client) *io.Buffer
}
