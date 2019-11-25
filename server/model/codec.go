package model

import (
	"github.com/honerlaw/go-osrs/io"
)

const (
	CODEC_STATE_HANDSHAKE = 0
	CODEC_STATE_LOGIN     = 1
	CODEC_STATE_GAME      = 2
)

type Codec interface {
	Decode(*io.Buffer, *Client) ([]Packet, error)
	Encode(*io.Buffer, *Client) *io.Buffer
}
