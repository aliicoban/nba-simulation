package config

import "sync"

var once sync.Once
var config *Config

type Config struct {
	App  App
	Game Game
}

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			App:  NewApp(),
			Game: NewGame(),
		}
	})

	return config
}
