package model

import "sync"

type GameState struct {
	mutex    *sync.Mutex
	Clients []Client
}

func NewGameState() *GameState {
	return &GameState{
		mutex:    &sync.Mutex{},
		Clients: make([]Client, 0),
	}
}

func (state *GameState) AddClient(client *Client) {
	state.mutex.Lock()

	state.Clients = append(state.Clients, *client)

	state.mutex.Unlock()
}
