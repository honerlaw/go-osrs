package login

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"log"
)

type LoginRequest struct {
	requestType            byte
	magicId                byte
	version                uint16
	memoryType             byte
	crcKeys                []uint32 // can toss on the client
	encryptedOpcodeSuccess byte
	decryptIsaacSeed       []uint32
	encryptIsaacSeed       []uint32
	clientId               uint32
	username               string
	password               string
}

func (l *LoginRequest) Handle(c *io.Client) []packet.Packet {
	if l.requestType != 16 && l.requestType != 18 {
		log.Print("Invalid login opcode ", l.requestType)

		c.Close()
		return nil
	}

	if l.magicId != 255 || l.version != 317 {
		log.Print("Invalid magic id or version", l.magicId, l.version)

		c.Close()
		return nil
	}

	if l.encryptedOpcodeSuccess != 10 {
		log.Print("Invalid decrypted success opcode", l.encryptedOpcodeSuccess)

		c.Close()
		return nil
	}

	log.Println(l.decryptIsaacSeed,l.encryptIsaacSeed);

	c.SetIsaacSeeds(l.decryptIsaacSeed, l.encryptIsaacSeed)
	c.Player.SetLoginInformation(l.memoryType, l.clientId, l.username, l.password)

	c.State = packet.PACKET_STATE_GAME

	return []packet.Packet{NewLoginResponse(2, 0)}
}

func (h *LoginRequest) Encode() *io.Buffer {
	return nil
}
