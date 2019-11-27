package packet

import (
	"errors"
)

type PacketEventObserver struct {
	listeners map[int32][]chan PacketEvent
}

func NewPacketEventObserver() *PacketEventObserver {
	return &PacketEventObserver{
		listeners: make(map[int32][]chan PacketEvent),
	}
}

func (ob *PacketEventObserver) Publish(data PacketEvent) {
	var listeners, ok = ob.listeners[data.EventCode()]
	if !ok {
		return
	}
	for _, listener := range listeners {
		listener <- data
	}
}

func (ob *PacketEventObserver) Subscribe(eventCode int32, channel chan PacketEvent) (func(), error) {
	var listeners, ok = ob.listeners[eventCode]
	if !ok {
		ob.listeners[eventCode] = []chan PacketEvent{channel}
	} else {
		for _, listener := range listeners {
			if listener == channel {
				return nil, errors.New("channel already listening for event code")
			}
		}
		listeners = append(listeners, channel)
	}
	return func() {
		ob.Unsubscribe(eventCode, channel)
	}, nil
}

func (ob *PacketEventObserver) Unsubscribe(eventCode int32, channel chan PacketEvent) {
	if _, ok := ob.listeners[eventCode]; !ok {
		return
	}
	for i := 0; i < len(ob.listeners[eventCode]); i++ {
		if ob.listeners[eventCode][i] == channel {
			ob.listeners[eventCode] = append(ob.listeners[eventCode][:i], ob.listeners[eventCode][i+1:]...)
			break;
		}
	}
}
