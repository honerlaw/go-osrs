package main

import (
	"github.com/honerlaw/go-osrs/game"
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"github.com/honerlaw/go-osrs/io/packet/incoming"
	"github.com/honerlaw/go-osrs/plugin"
	"log"
	"net"
)

func main() {
	var listener, err = net.Listen("tcp", "0.0.0.0:43594")

	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}

	var pluginLoader = plugin.NewPluginLoader()
	var tickHandler = game.NewTickHandler()

	pluginLoader.Load()

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
		var packetEventObserver = packet.NewPacketEventObserver()
		var handler = incoming.NewPacketHandler(client, packetEventObserver)
		pluginLoader.Register(client, packetEventObserver)

		go handler.Listen()
	}
}
