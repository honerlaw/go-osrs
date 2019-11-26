package outgoing

import (
	"github.com/honerlaw/go-osrs/io"
)

type PlayerUpdate struct {

}

func NewPlayerUpate() *PlayerUpdate {
	return &PlayerUpdate{}
}

func (h *PlayerUpdate) Encode(client *io.Client) *io.Buffer {
	var buf = io.NewBuffer(256)
	buf.WriteVariableShortOpcode(81, client.IsaacEncryptor.NextValue())
	buf.EnableBitMode()

	// no update required
	buf.WriteBits(1, 0)

	// no other players to update
	buf.WriteBits(8, 0)


	buf.DisableBitMode()
	buf.WriteVariableShortLength()
	return buf
}
