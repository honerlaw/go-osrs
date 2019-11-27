package plugin

import (
	"errors"
	"github.com/honerlaw/go-osrs/io"
	"github.com/honerlaw/go-osrs/io/packet"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

type PluginLoader struct {
	plugins []GamePlugin
}

func NewPluginLoader() *PluginLoader {
	return &PluginLoader{
		plugins: make([]GamePlugin, 0),
	}
}

func (loader *PluginLoader) Load() {
	err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, ".so") {
			return nil
		}

		plug, err := plugin.Open(path)
		if err != nil {
			return err
		}

		symPlug, err := plug.Lookup("PluginMain")
		if err != nil {
			return err
		}

		gamePlugin, ok := symPlug.(GamePlugin)
		if !ok {
			return errors.New("plugin not type of GamePlugin")
		}

		err = gamePlugin.Init()
		if err != nil {
			return err
		}

		loader.plugins = append(loader.plugins, gamePlugin)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (loader *PluginLoader) Register(client *io.Client, observer *packet.PacketEventObserver) {
	for _, gamePlugin := range loader.plugins {
		var channel = make(chan packet.PacketEvent) // only the specific plugin events go to this channel

		// use the same channel to subscribe to all required events
		for _, code := range gamePlugin.EventCodes() {
			_, err := observer.Subscribe(code, channel)
			if err != nil {
				log.Fatal(err)
			}
		}

		// okay next up we need to basically listen for incoming events, and call handle whenever we get one
		go func(ch chan packet.PacketEvent) {
			for {
				event := <- ch

				err := gamePlugin.Handle(client, event)
				if err != nil {
					log.Print(err)
				}
			}
		}(channel)
	}
}
