package main

import (
	"github.com/honerlaw/go-osrs/codec"
	"github.com/honerlaw/go-osrs/game"
	"github.com/honerlaw/go-osrs/model"
	"log"
	"net"
)

var codecMap = map[byte]model.Codec{
	model.CODEC_STATE_GAME:      codec.NewGameCodec(),
	model.CODEC_STATE_LOGIN:     codec.NewLoginCodec(),
	model.CODEC_STATE_HANDSHAKE: codec.NewHandshakeCodec(),
}

func main() {
	var listener, err = net.Listen("tcp", "0.0.0.0:43594")

	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}

	var gameState = model.NewGameState()
	var playerUpdate = game.NewPlayerUpdate(gameState)

	// start the player update cycle
	go playerUpdate.Cycle()

	log.Print("Listening for new connections")
	for {
		var conn, err = listener.Accept()

		if err != nil {
			log.Print("Failed to accept new connection", err)
			continue
		}

		log.Print("Accepted new connection", conn.LocalAddr())

		var client = model.NewClient(conn)
		var handler = game.NewRequestHandler(client, codecMap)

		gameState.AddClient(client)

		go handler.Listen()
	}
}
