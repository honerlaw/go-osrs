package packet

import (
	"github.com/honerlaw/go-osrs/io"
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

func NewLoginRequest(requestType byte, magicId byte, version uint16, memoryType byte,
	crcKeys []uint32, encryptedOpcodeSuccess byte, decryptIsaacSeed []uint32,
	encryptIsaacSeed []uint32, clientId uint32, username string, password string) *LoginRequest {
	return &LoginRequest{
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
	}
}

func (l *LoginRequest) Handle(c *io.Client) []Packet {
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

	c.SetDecryptor(io.NewIsaac(l.decryptIsaacSeed))
	c.SetEncryptor(io.NewIsaac(l.encryptIsaacSeed))
	c.Player.SetLoginInformation(l.memoryType, l.clientId, l.username, l.password)

	c.MoveToNextCodecState()

	return []Packet{NewLoginResponse(2, 0)}
}

func (h *LoginRequest) Encode() *io.Buffer {
	return nil
}
