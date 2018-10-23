package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
)

const (
	CLIENT_STATE_HANDSHAKE = 0
	CLIENT_STATE_LOGIN     = 1
	CLIENT_STATE_GAME      = 2
)

type Client struct {
	connection net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
	buffer     *Buffer
	state      byte
	player     *Player
}

func (c *Client) Listen() {
	for {
		var err = c.buffer.Read(c.reader)
		if err != nil {
			log.Print("Failed to read data", err)
			return;
		}

		if c.buffer.Length() == 0 {
			continue
		}

		switch c.state {
		case CLIENT_STATE_HANDSHAKE:
			c.handshake()
		case CLIENT_STATE_LOGIN:
			c.login()
		case CLIENT_STATE_GAME:
			c.game()
		}
	}
}

func (c *Client) Write(buf *Buffer, flush bool) {
	var _, err = c.writer.Write(buf.AsByteArray())
	if err != nil {
		log.Print("Failed to write buffer to stream", err)
	}
	if (flush) {
		var flushErr = c.writer.Flush()
		if flushErr != nil {
			log.Print("Failed to flush buffer", flushErr)
		}
	}
}

func (c *Client) handshake() {
	if c.buffer.Length() != 2 {
		log.Print("HANDSHAKE: 2 bytes should have been read! ", c.buffer.Length())

		c.buffer.Compact()
		return
	}

	var opcode = c.buffer.ReadByte()
	var namehash = c.buffer.ReadByte()

	c.player = &Player{
		namehash: namehash,
	};

	// @todo handle other handshake opcodes
	switch opcode {
	case 14:
		var buf = NewBuffer(17);
		buf.WriteLong(0)
		buf.WriteByte(0)
		buf.WriteLong(rand.Int63()) // @todo need a crypto secure version of this
		c.Write(buf, true)

		c.state = CLIENT_STATE_LOGIN
	default:
		log.Print("Unknown handeshake opcode", 14)
	}
}

func (c *Client) login() {
	if c.buffer.Length() < 2 {
		log.Print("LOGIN: 2 bytes should have been read! ", c.buffer.Length())

		c.buffer.Compact()
		return
	}

	var opcode = c.buffer.ReadByte()
	var expectedBlockLength = c.buffer.ReadByte()

	if opcode != 16 && opcode != 18 {
		log.Print("Invalid login opcode ", opcode)
		return;
	}

	if c.buffer.Remaining() != int(expectedBlockLength) {
		log.Print("Not enough data remaining to read entire block ", c.buffer.Remaining(), expectedBlockLength)
		return;
	}

	var magicId = c.buffer.ReadByte()  // 255
	var version = c.buffer.ReadShort() // should be 317
	var memoryType = c.buffer.ReadByte()
	var crcKeys = make([]uint32, 9)
	for i := 0; i < cap(crcKeys); i++ {
		crcKeys[i] = c.buffer.ReadInt()
	}
	var encryptedBlockLength = c.buffer.ReadByte()
	var encryptedOpcodeSuccess = c.buffer.ReadByte() // should be 10

	var clientIsaacSeed = c.buffer.ReadLong()
	var serverIsaacSeed = c.buffer.ReadLong()

	var userId = c.buffer.ReadInt()

	var username, _ = c.buffer.ReadRSString()
	var password, _ = c.buffer.ReadRSString()

	log.Print("RANDOM", magicId, version, memoryType, crcKeys, encryptedBlockLength, encryptedOpcodeSuccess, clientIsaacSeed, serverIsaacSeed, userId)
	log.Print("USERNAME: ", username)
	log.Print("PASSWORD: ", password)

	c.connection.Close()
}

func (c *Client) game() {

}
