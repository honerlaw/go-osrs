package incoming

import (
	"errors"
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	data2 "github.com/honerlaw/go-osrs/io/packet/data"
	"log"
)

type Login struct {
}

func NewLogin() *Login {
	return &Login{
	}
}

func (req *Login) Decode(buffer *io.Buffer, client *io.Client, _ byte, _ int8) ([]packet.PacketData, error) {
	if buffer.Length() < 2 {
		log.Print("LOGIN: 2 bytes should have been read! ", buffer.Length())

		buffer.Compact()
		return nil, nil
	}

	var requestType = buffer.ReadByte()
	if requestType != 16 && requestType != 18 {
		log.Print("Invalid login opcode ", requestType)

		client.Close()
		return nil, nil
	}

	var expectedBlockLength = uint32(buffer.ReadByte())

	// check if enough data is in the buffer, if not enough wait for more
	if expectedBlockLength < buffer.Remaining() {
		log.Print("LOGIN: not enough data in buffer", expectedBlockLength, buffer.Remaining())

		buffer.Compact()
		return nil, nil
	}

	// so we know we have enough, now make sure they  match and we don't have additional junk data
	if expectedBlockLength != buffer.Remaining() {
		log.Print("LOGIN: block length does not match buffer remaining", expectedBlockLength, buffer.Remaining())
		return nil, errors.New("Invalid buffer leength")
	}

	var magicId = buffer.ReadByte()
	var version = buffer.ReadShort()
	if magicId != 255 || version != 317 {
		log.Print("Invalid magic id or version", magicId, version)

		client.Close()
		return nil, nil
	}

	var memoryType = buffer.ReadByte()

	// we can ignore these
	var crcKeys = make([]uint32, 9)
	for i := 0; i < cap(crcKeys); i++ {
		crcKeys[i] = buffer.ReadInt()
	}

	var encryptedBlockLength = uint32(buffer.ReadByte())
	if encryptedBlockLength != buffer.Remaining() {
		log.Print("LOGIN: encrypted block length does not match remaining buffer", encryptedBlockLength, buffer.Remaining())
		return nil, errors.New("Invalid buffer length")
	}

	var encryptedOpcodeSuccess = buffer.ReadByte()
	if encryptedOpcodeSuccess != 10 {
		log.Print("Invalid decrypted success opcode", encryptedOpcodeSuccess)

		client.Close()
		return nil, nil
	}

	var decryptIsaacSeed = []uint32{
		buffer.ReadBEInt(),
		buffer.ReadBEInt(),
		buffer.ReadBEInt(),
		buffer.ReadBEInt(),
	}
	var encryptIsaacSeed = make([]uint32, 4)
	for index, seed := range decryptIsaacSeed {
		encryptIsaacSeed[index] = seed + 50
	}
	var clientId = buffer.ReadInt()
	var username, _ = buffer.ReadRSString()
	var password, _ = buffer.ReadRSString()

	return []packet.PacketData{
		data2.NewLoginData(
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
		),
	}, nil
}
