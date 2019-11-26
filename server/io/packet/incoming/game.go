package incoming

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"log"
)

var packetLengths = []int8{
	0, 0, 0, 1, -1, 0, 0, 0, 0, 0,  // 0
	0, 0, 0, 0, 8, 0, 6, 2, 2, 0,   // 10
	0, 2, 0, 6, 0, 12, 0, 0, 0, 0,  // 20
	0, 0, 0, 0, 0, 8, 4, 0, 0, 2,   // 30
	2, 6, 0, 6, 0, -1, 0, 0, 0, 0,  // 40
	0, 0, 0, 12, 0, 0, 0, 0, 8, 0,  // 50
	0, 8, 0, 0, 0, 0, 0, 0, 0, 0,   // 60
	6, 0, 2, 2, 8, 6, 0, -1, 0, 6,  // 70
	0, 0, 0, 0, 0, 1, 4, 6, 0, 0,   // 80
	0, 0, 0, 0, 0, 3, 0, 0, -1, 0,  // 90
	0, 13, 0, -1, 0, 0, 0, 0, 0, 0, // 100
	0, 0, 0, 0, 0, 0, 0, 6, 0, 0,   // 110
	1, 0, 6, 0, 0, 0, -1, 0, 2, 6,  // 120
	0, 4, 6, 8, 0, 6, 0, 0, 0, 2,   // 130
	0, 0, 0, 0, 0, 6, 0, 0, 0, 0,   // 140
	0, 0, 1, 2, 0, 2, 6, 0, 0, 0,   // 150
	0, 0, 0, 0, -1, -1, 0, 0, 0, 0, // 160
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,   // 170
	0, 8, 0, 3, 0, 2, 0, 0, 8, 1,   // 180
	0, 0, 12, 0, 0, 0, 0, 0, 0, 0,  // 190
	2, 0, 0, 0, 0, 0, 0, 0, 4, 0,   // 200
	4, 0, 0, 0, 7, 8, 0, 0, 10, 0,  // 210
	0, 0, 0, 0, 0, 0, -1, 0, 6, 0,  // 220
	1, 0, 0, 0, 6, 0, 6, 8, 1, 0,   // 230
	0, 4, 0, 0, 0, 0, -1, 0, -1, 4, // 240
	0, 0, 6, 6, 0, 0, 0,            // 250
}

var packetDecoderMap = map[byte]packet.IncomingPacket{
	0: NewHandshake(),
}

type Game struct {

}

func NewGame() *Game {
	return &Game{}
}

func (game *Game) Decode(b *io.Buffer, c *io.Client, _ byte, _ int8) ([]packet.PacketData, error) {
	var opcode = byte(uint32(b.ReadByte()) - c.IsaacDecryptor.NextValue())
	var length = packetLengths[ opcode ]

	// variable in length
	if length == -1 {

		// has at least one byte that we can read to fetch the length
		if b.Remaining() > 0 {
			length = int8(b.ReadByte())

			if uint32(length) < b.Remaining() {
				log.Print("Not enough data to read entire packet ", length, b.Remaining())

				b.Compact()
				return nil, nil
			}
		}
	}

	// still have no idea what the hell the length is, so discard everything
	if length == -1 {
		return nil, nil
	}

	var decoder, exists = packetDecoderMap[opcode]
	if !exists {
		return nil, nil
	}

	return decoder.Decode(b, c, opcode, length)
}
