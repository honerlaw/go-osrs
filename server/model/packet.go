package model

import (
	"github.com/honerlaw/go-osrs/io"
)

type Packet interface {
	Handle(*Client) []Packet
	Encode() *io.Buffer
}