package main

import (
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"github.com/honerlaw/go-osrs/io/packet/event"
	"log"
)

// export the login plugin so its accessible
var PluginMain loginPlugin

type loginPlugin struct {}

func (plug *loginPlugin) Init() error {
	return nil
}

func (plug *loginPlugin) EventCodes() []int32 {
	return []int32{event.LoginEventCode}
}

func (plug *loginPlugin) Handle(client *io.Client, event packet.PacketEvent) error {

	log.Println(event)

	return nil
}


