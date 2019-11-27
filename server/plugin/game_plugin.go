package plugin

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
)

type GamePlugin interface {

	Init() error // general initialization

	EventCodes() []int32 // what event codes to listen for

	Handle(*io.Client, packet.PacketEvent) error

}
