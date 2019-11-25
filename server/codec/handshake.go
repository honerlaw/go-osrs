package codec

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/packet"
	"log"
)

type HandshakeCodec struct {

}

func NewHandshakeCodec() *HandshakeCodec {
	return &HandshakeCodec{}
}

func (codec *HandshakeCodec) Decode(b *io.Buffer, _ *io.Client) ([]packet.Packet, error) {
	if b.Length() != 2 {
		log.Print("HANDSHAKE: 2 bytes should have been read! ", b.Length())

		b.Compact()
		return nil, nil
	}

	return []packet.Packet{packet.NewHandshakeRequest(b.ReadByte(), b.ReadByte())}, nil
}

func (codec *HandshakeCodec) Encode(b *io.Buffer, _ *io.Client) *io.Buffer {
	return b
}
