package incoming

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"log"
)

type PacketHandler struct {
	client      *io.Client
}

func NewPacketHandler(client *io.Client) *PacketHandler {
	return &PacketHandler{
		client:      client,
	}
}

func (handler *PacketHandler) Listen() {
	for {
		var err = handler.client.Buffer.Read(handler.client.Reader)
		if err != nil {
			log.Print("Failed to read data", err)
			return
		}

		if handler.client.Buffer.Length() == 0 {
			continue
		}

		data, err := handler.decode()
		if err != nil {
			handler.client.Close()
			return
		}

		// no packet data to emit / handle, so just continue listening
		if data == nil {
			continue
		}

		// @todo event out the incoming data to listeners
	}
}

func (handler *PacketHandler) decode() ([]packet.PacketData, error) {
	switch handler.client.State {
	case 0:
		return NewHandshake().Decode(handler.client.Buffer, handler.client, 0, -1)
	case 1:
		return NewLogin().Decode(handler.client.Buffer, handler.client, 0, -1)
	case 2:
		return NewGame().Decode(handler.client.Buffer, handler.client, 0, -1)
	default:
		return nil, nil
	}
}
