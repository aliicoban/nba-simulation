package nba

type NBA interface {
	CreateGame()
	UpdateScore()
}

type API struct {
	Db NBA
}
