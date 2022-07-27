package main

import (
	game "github.com/alicobanserver/internal/game"
	store "github.com/alicobanserver/internal/store"
	service "github.com/alicobanserver/internal/service"
)

func main() {
	db := service.NewDB()

	api := store.API{
		Db: db,
	}

	game.StartGame(api)
}
