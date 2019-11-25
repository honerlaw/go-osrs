package codec

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/packet"
)

const (
	CODEC_STATE_HANDSHAKE = 0
	CODEC_STATE_LOGIN     = 1
	CODEC_STATE_GAME      = 2
)

type Codec interface {
	Decode(*io.Buffer, *io.Client) ([]packet.Packet, error)
	Encode(*io.Buffer, *io.Client) *io.Buffer
}
