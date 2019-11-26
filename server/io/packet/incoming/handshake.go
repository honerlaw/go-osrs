package incoming

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	data2 "github.com/honerlaw/go-osrs/io/packet/data"
	"github.com/honerlaw/go-osrs/model"
	"log"
	"math/rand"
)

type Handshake struct {

}

func NewHandshake() *Handshake {
	return &Handshake{}
}

func (req *Handshake) Decode(buffer *io.Buffer, client *io.Client, _ byte, _ int8) ([]packet.PacketData, error) {
	if buffer.Length() != 2 {
		log.Print("HANDSHAKE: 2 bytes should have been read! ", buffer.Length())

		buffer.Compact()
		return nil, nil
	}

	var opcode = buffer.ReadByte()
	var namehash = buffer.ReadByte()

	client.Player = model.NewPlayer(namehash)
	switch opcode {
	case 14:
		client.MoveToNextCodecState()

		return []packet.PacketData{data2.NewHandshakeData(0, rand.Int63())}, nil
	}
	return nil, nil
}
