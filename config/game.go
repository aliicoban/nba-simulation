package config

type Game struct {
	TimeInterval string
}

func NewGame() Game {
	return Game{
		TimeInterval: Getenv("TIME_INTERVAL", "5"),
	}
}
