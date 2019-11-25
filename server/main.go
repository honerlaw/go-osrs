package main

import (
	"github.com/honerlaw/go-osrs/codec"
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/packet"
	"log"
	"net"
)

var codecMap = map[byte]codec.Codec{
	codec.CODEC_STATE_GAME:      codec.NewGameCodec(),
	codec.CODEC_STATE_LOGIN:     codec.NewLoginCodec(),
	codec.CODEC_STATE_HANDSHAKE: codec.NewHandshakeCodec(),
}

func main() {
	var listener, err = net.Listen("tcp", "0.0.0.0:43594")

	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}

	log.Print("Listening for new connections")

	for {
		var conn, err = listener.Accept()

		if err != nil {
			log.Print("Failed to accept new connection", err)
			continue
		}

		log.Print("Accepted new connection", conn.LocalAddr())

		var client = io.NewClient(conn)

		go listen(client)
	}
}

/*

The flow of this is basically

1. call a state codec to decode the initial data into a set of packets to handle
2. handle those packets in some way (maybe we should use eventing to event out that the packet occurred instead?)
2.1. handling will return new packets to send out
3. encode the new packets to send out using the individual packet encoders
4. encode the packet using the game encoder
5. send the packet back out

Future Flow
1. state codec decodes the initial data into a set of packets
2. packet data is evented out to be handled by whatever is listening for the different packets
3. outgoing packets are written to an "out" buffer as they are handled by listeners (we only write full packets at a time)
4. we periodically flush the buffer
*/
func listen(c *io.Client) {
	for {
		var err = c.Buffer.Read(c.Reader)
		if err != nil {
			log.Print("Failed to read data", err)
			return
		}

		if c.Buffer.Length() == 0 {
			continue
		}

		var codec = codecMap[c.State]

		// decode the buffer into a packets
		var decodedPackets, decodeErr = codec.Decode(c.Buffer, c)

		// something really bad happened if we have an error from the decoder, so just disconnect
		if decodeErr != nil {
			c.Close()
			return
		}

		// no packet was given packet, so do nothing
		if decodedPackets == nil {
			continue
		}

		// handle the decoded packets and get the responses to send out
		var allResponsePackets [][]packet.Packet = nil
		for _, decodedPacket := range decodedPackets {
			var packets = decodedPacket.Handle(c)
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

				var encodedBuf = codec.Encode(responsePacket.Encode(), c)
				c.Write(encodedBuf, true)
			}
		}
	}
}
