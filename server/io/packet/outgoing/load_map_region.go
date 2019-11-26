package outgoing

import (
	"github.com/honerlaw/go-osrs/io"
)

type LoadMapRegion struct {
}

func NewLoadMapRegion() *LoadMapRegion {
	return &LoadMapRegion{}
}

func (h *LoadMapRegion) Encode(client *io.Client) *io.Buffer {
	var buf = io.NewBuffer(5)
	buf.WriteOpcode(73, client.IsaacEncryptor.NextValue())
	buf.WriteBEShortA(int16(client.Player.Position.RegionX() + 6));
	buf.WriteBEShort(int16(client.Player.Position.RegionY() + 6));
	return buf
}
