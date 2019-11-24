package main

import (
	"log"
	"net"
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"github.com/honerlaw/go-osrs/io/packet/game"
	"github.com/honerlaw/go-osrs/io/packet/handshake"
	"github.com/honerlaw/go-osrs/io/packet/login"
)

var packetDecoderMap = map[byte]packet.PacketStateDecoder{
	packet.PACKET_STATE_HANDSHAKE: handshake.Decoder,
	packet.PACKET_STATE_LOGIN:     login.Decoder,
	packet.PACKET_STATE_GAME:      game.Decoder,
}

var packetEncoderMap = map[byte]packet.PacketStateEncoder{
	packet.PACKET_STATE_GAME: game.Encoder,
}

func main() {
	var listener, err = net.Listen("tcp", ":43594")

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

		var decoder = packetDecoderMap[c.State]
		var encoder = packetEncoderMap[c.State]

		// decode the buffer into a packet
		var decodedPackets, decodeErr = decoder(c.Buffer, c)

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

				var encodedBuf= responsePacket.Encode()

				// if additional encoding is needed, use it
				if encoder != nil {
					encodedBuf = encoder(encodedBuf, c)
				}
				c.Write(encodedBuf, true)
			}
		}
	}
}
