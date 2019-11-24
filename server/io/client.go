package io

import (
	"bufio"
	"log"
	"net"
	"github.com/honerlaw/go-osrs/model"
)

type Client struct {
	connection       net.Conn
	Reader           *bufio.Reader
	writer           *bufio.Writer
	State            byte
	Buffer           *Buffer
	Player           *model.Player
	IsaacDecryptor   *Isaac
	IsaacEncryptor   *Isaac
	DecryptIsaacSeed []uint32
	EncryptIsaacSeed []uint32
}

func NewClient(connection net.Conn) *Client {
	return &Client{
		connection: connection,
		Reader:     bufio.NewReader(connection),
		writer:     bufio.NewWriter(connection),
		Buffer:     NewBuffer(256),
	}
}

func (c *Client) Write(buf *Buffer, flush bool) {
	var _, err = c.writer.Write(buf.AsByteArray())
	if err != nil {
		log.Print("Failed to write buffer to stream", err)
	}
	if flush {
		var flushErr = c.writer.Flush()
		if flushErr != nil {
			log.Print("Failed to flush buffer", flushErr)
		}
	}
}

func (c *Client) Close() {
	var err = c.connection.Close()
	if err != nil {
		log.Print("Failed to close connection", err)
	}
}

func (c *Client) SetIsaacSeeds(decryptIsaacSeed []uint32, encryptIsaacSeed []uint32) {
	c.IsaacDecryptor = NewIsaac(decryptIsaacSeed)
	c.IsaacEncryptor = NewIsaac(encryptIsaacSeed)
}
