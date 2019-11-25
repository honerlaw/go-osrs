package game

import (
	"github.com/honerlaw/go-osrs/model"
	"log"
)

type RequestHandler struct {
	client   *model.Client
	codecMap map[byte]model.Codec
}

func NewRequestHandler(client *model.Client, codecMap map[byte]model.Codec) *RequestHandler {
	return &RequestHandler{
		client:   client,
		codecMap: codecMap,
	}
}

/*

The flow of this is basically

1. call a state codec to decode the initial data into a set of packets to handle
2. handle those packets in some way (maybe we should use eventing to event out that the packet occurred instead?)
2.1. handling will return new packets to send outs
3. encode the new packets to send out using the individual packet encoders
4. encode the packet using the game encoder
5. send the packet back out

*/
func (handler *RequestHandler) Listen() {
	for {
		var err = handler.client.Buffer.Read(handler.client.Reader)
		if err != nil {
			log.Print("Failed to read data", err)
			return
		}

		if handler.client.Buffer.Length() == 0 {
			continue
		}

		var codec = handler.codecMap[handler.client.State]

		// decode the buffer into a packets
		var decodedPackets, decodeErr = codec.Decode(handler.client.Buffer, handler.client)

		// something really bad happened if we have an error from the decoder, so just disconnect
		if decodeErr != nil {
			handler.client.Close()
			return
		}

		// no packet was given packet, so do nothing
		if decodedPackets == nil {
			continue
		}

		// handle the decoded packets and get the responses to send out
		var allResponsePackets [][]model.Packet = nil
		for _, decodedPacket := range decodedPackets {
			var packets = decodedPacket.Handle(handler.client)
			if packets == nil {
				continue
			}
			allResponsePackets = append(allResponsePackets, packets)
		}

		// no response packets to send out, do nothing
		if allResponsePackets == nil {
			continue
		}

		// send out all the response packets
		for _, responsePackets := range allResponsePackets {
			for _, responsePacket := range responsePackets {

				var encodedPacket = responsePacket.Encode()
				if encodedPacket == nil {
					continue;
				}

				var encodedBuf = codec.Encode(encodedPacket, handler.client)
				handler.client.Write(encodedBuf, true)
			}
		}
	}
}
