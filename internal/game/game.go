package game

import (
	"context"
	config "github.com/alicobanserver/config"
	store "github.com/alicobanserver/internal/store"
	"strconv"
	"time"
)

type Game interface {
	Start(ctx context.Context, api store.API)
}

type ticker struct {
	timer     *time.Timer
	totalTime time.Duration
}

const (
	gameTime  = 48 * 60 // seconds
	gameSpeed = 12      // for simulation
)

// Start ..
func Start(ctx context.Context, api store.API) {
	ticker := &ticker{}
	cfg := config.GetConfig()
	seconds, _ := strconv.Atoi(cfg.Game.TimeInterval)
	ticker.updateTimer(seconds)
	api.Db.CreateGame()

	for {
		<-ticker.timer.C
		api.Db.UpdateScore()
		ticker.updateTimer(seconds)
		ticker.totalTime += time.Second * time.Duration(seconds) * gameSpeed

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
