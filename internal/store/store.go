package nba

type NBA interface {
	CreateGame()
	UpdateScore()
}
 
type API struct {
	Db NBA
	Game Game
}

type Game struct {
	IsStarted bool
	Duration  int
	DoneChan  chan struct{}
}

