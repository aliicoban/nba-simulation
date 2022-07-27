package game

import (
	config "github.com/alicobanserver/config"
	store "github.com/alicobanserver/internal/store"
	"strconv"
	"time"
)

type ticker struct {
	timer     *time.Timer
	totalTime time.Duration
}

const gameTime = 240

func StartGame(api store.API) {

	ticker := &ticker{}
	cfg := config.GetConfig()
	seconds, _ := strconv.Atoi(cfg.Game.TimeInterval)
	ticker.updateTimer(seconds)
	api.Db.CreateGame()
	for {
		<-ticker.timer.C
		api.Db.UpdateScore()
		ticker.updateTimer(seconds)
		ticker.totalTime += time.Second * time.Duration(seconds) 

		if ticker.totalTime == time.Second*time.Duration(gameTime) {
			ticker.timer.Stop()
		}

	}
}

func (t *ticker) updateTimer(seconds int) {
	diff := time.Second * time.Duration(seconds)

	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}
