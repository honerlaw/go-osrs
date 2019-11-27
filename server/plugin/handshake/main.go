package main

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"github.com/honerlaw/go-osrs/io/packet/event"
	"log"
)

// export the login plugin so its accessible
var PluginMain handshakePlugin

type handshakePlugin struct {}

func (plug *handshakePlugin) Init() error {
	return nil
}

func (plug *handshakePlugin) EventCodes() []int32 {
	return []int32{event.HandshakeEventCode}
}

func (plug *handshakePlugin) Handle(client *io.Client, event packet.PacketEvent) error {

	log.Println(event)

	return nil
}


