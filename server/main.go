package main

import (
	"github.com/honerlaw/go-osrs/game"
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet/incoming"
	"log"
	"net"
)

func main() {
	var listener, err = net.Listen("tcp", "0.0.0.0:43594")

	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}

	var tickHandler = game.NewTickHandler()

	// start the player update cycle
	go tickHandler.Cycle()

	log.Print("Listening for new connections")
	for {
		var conn, err = listener.Accept()

		if err != nil {
			log.Print("Failed to accept new connection", err)
			continue
		}

		log.Print("Accepted new connection", conn.LocalAddr())

		var client = io.NewClient(conn)
		var handler = incoming.NewPacketHandler(client)

		go handler.Listen()
	}
}
