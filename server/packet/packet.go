package packet

import "github.com/honerlaw/go-osrs/io"

type Packet interface {
	Handle(*io.Client) []Packet
	Encode() *io.Buffer
}