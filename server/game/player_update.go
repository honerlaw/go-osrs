package game

import (
	"github.com/honerlaw/go-osrs/model"
	"log"
	"time"
)

type PlayerUpdate struct {
	ticker    *time.Ticker
	gameState *model.GameState
}

func NewPlayerUpdate(gameState *model.GameState) *PlayerUpdate {
	return &PlayerUpdate{
		ticker:    time.NewTicker(600 * time.Millisecond),
		gameState: gameState,
	}
}

func (update *PlayerUpdate) Cycle() {
	for ; true; <-update.ticker.C {
		// https://github.com/honerlaw/osrs/blob/master/server/src/main/java/org/server/osrs/PlayerUpdating.java
		// split the above into packets for each "update section", the packets then can just return the raw buffer and merge
		// into a single buffer, also need to sort out bit manipulation and crap :/

		log.Println("test")
	}
}
