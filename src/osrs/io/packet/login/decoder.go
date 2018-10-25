package login

import (
	"osrs/io"
	"osrs/io/packet"
	"errors"
	"log"
)

func Decoder(b *io.Buffer, _ *io.Client) ([]packet.Packet, error) {
	if b.Length() < 2 {
		log.Print("LOGIN: 2 bytes should have been read! ", b.Length())

		b.Compact()
		return nil, nil
	}

	var requestType = b.ReadByte()
	var expectedBlockLength = int(b.ReadByte())

	// check if enough data is in the buffer, if not enough wait for more
	if expectedBlockLength < b.Remaining() {
		log.Print("LOGIN: not enough data in buffer", expectedBlockLength, b.Remaining())

		b.Compact()
		return nil, nil
	}

	// so we know we have enough, now make sure they  match and we don't have additional junk data
	if expectedBlockLength != b.Remaining() {
		log.Print("LOGIN: block length does not match buffer remaining", expectedBlockLength, b.Remaining())
		return nil, errors.New("Invalid buffer leength")
	}

	var magicId = b.ReadByte()
	var version = b.ReadShort()
	var memoryType = b.ReadByte()

	// we can ignore these
	var crcKeys = make([]uint32, 9)
	for i := 0; i < cap(crcKeys); i++ {
		crcKeys[i] = b.ReadInt()
	}

	var encryptedBlockLength = int(b.ReadByte())
	if encryptedBlockLength != b.Remaining() {
		log.Print("LOGIN: encrypted block length does not match remaining buffer", encryptedBlockLength, b.Remaining())
		return nil, errors.New("Invalid buffer length")
	}

	var encryptedOpcodeSuccess = b.ReadByte()
	var decryptIsaacSeed = []uint32{
		b.ReadInt(),
		b.ReadInt(),
		b.ReadInt(),
		b.ReadInt(),
	}
	var encryptIsaacSeed = make([]uint32, 4)
	for index, seed := range decryptIsaacSeed {
		encryptIsaacSeed[index] = seed + 50
	}
	var clientId = b.ReadInt()
	var username, _ = b.ReadRSString()
	var password, _ = b.ReadRSString()

	return []packet.Packet{&LoginRequest{
		requestType,
		magicId,
		version,
		memoryType,
		crcKeys,
		encryptedOpcodeSuccess,
		decryptIsaacSeed,
		encryptIsaacSeed,
		clientId,
		username,
		password,
	}}, nil
}
