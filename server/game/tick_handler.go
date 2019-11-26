package game

import (
	"log"
	"time"
)

type TickHandler struct {
	ticker *time.Ticker
}

func NewTickHandler() *TickHandler {
	return &TickHandler{
		ticker: time.NewTicker(600 * time.Millisecond),
	}
}

func (update *TickHandler) Cycle() {
	for ; true; <-update.ticker.C {
		// https://github.com/honerlaw/osrs/blob/master/server/src/main/java/org/server/osrs/PlayerUpdating.java
		// split the above into packets for each "update section", the packets then can just return the raw buffer and merge
		// into a single buffer, also need to sort out bit manipulation and crap :/

		log.Println("test")
	}
}
