package handshake

import (
	"log"
	"osrs/io"
	"osrs/io/packet"
)

func Decoder(b *io.Buffer, _ *io.Client) ([]packet.Packet, error) {
	if b.Length() != 2 {
		log.Print("HANDSHAKE: 2 bytes should have been read! ", b.Length())

		b.Compact()
		return nil, nil
	}

	return []packet.Packet{NewHandshakeRequest(b.ReadByte(), b.ReadByte())}, nil
}
