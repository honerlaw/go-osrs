package incoming

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"log"
)

type PacketHandler struct {
	client   *io.Client
	observer *packet.PacketEventObserver
}

func NewPacketHandler(client *io.Client, observer *packet.PacketEventObserver) *PacketHandler {
	return &PacketHandler{
		client: client,
		observer: observer,
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

		// fire off the data to any listeners so the listeners can decide what to do with it
		for _, datum := range data {
			handler.observer.Publish(datum)
		}
	}
}

func (handler *PacketHandler) decode() ([]packet.PacketEvent, error) {
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
