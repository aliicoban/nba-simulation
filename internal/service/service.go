package service

import (
	"fmt"

	nba "github.com/alicobanserver/internal"
	game "github.com/alicobanserver/internal/game"
	store "github.com/alicobanserver/internal/store"

	"context"
	"math/rand"
)

// Service ...
type Service interface {
	Start(ctx context.Context)
}

// service ...
type service struct {
	game game.Game
}

// NewService ...
func NewService() Service {
	return &service{}
}

// Store ...
type Store struct {
	Statistics nba.Statistic
	Players    []nba.Player
}

// NewDB
func NewDB() *Store {
	return &Store{
		Statistics: nba.Statistic{},
		Players:    []nba.Player{},
	}
}

// Start ...
func (s *service) Start(ctx context.Context) {
	db := NewDB()
	api := store.API{
		Db: db,
	}
	game.Start(ctx, api)

}

var attackCountInAMinute = []int{1, 5}

const homeTeamName = "TeamA"
const awayTeamName = "TeamB"

func (api *Store) CreateGame() {
	homeTeam := nba.Team{
		Name: "TeamA",
	}
	awayTeam := nba.Team{
		Name: "TeamB",
	}

	homeTeamPlayersLineup := []nba.Player{
		{
			Name:    "PlayerA1",
			Assists: 0,
			Points:  0,
			Team:    homeTeam,
		},
		{
			Name:    "PlayerA2",
			Assists: 0,
			Points:  0,
			Team:    homeTeam,
		},
		{
			Name:    "PlayerA3",
			Assists: 0,
			Points:  0,
			Team:    homeTeam,
		},
		{
			Name:    "PlayerA4",
			Assists: 0,
			Points:  0,
			Team:    homeTeam,
		},
		{
			Name:    "PlayerA5",
			Assists: 0,
			Points:  0,
			Team:    homeTeam,
		},
	}
	awayTeamPlayersLineup := []nba.Player{
		{
			Name:    "PlayerB1",
			Assists: 0,
			Points:  0,
			Team:    awayTeam,
		},
		{
			Name:    "PlayerB2",
			Assists: 0,
			Points:  0,
			Team:    awayTeam,
		},
		{
			Name:    "PlayerB3",
			Assists: 0,
			Points:  0,
			Team:    awayTeam,
		},
		{
			Name:    "PlayerB4",
			Assists: 0,
			Points:  0,
			Team:    awayTeam,
		},
		{
			Name:    "PlayerB5",
			Assists: 0,
			Points:  0,
			Team:    awayTeam,
		},
	}

	api.Statistics = nba.Statistic{
		AwayTeam:      awayTeam,
		HomeTeam:      homeTeam,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		TopAssist:     nba.Player{},
		TopScorer:     nba.Player{},
	}

	api.Players = append(homeTeamPlayersLineup, awayTeamPlayersLineup...)

	fmt.Println(api.Statistics)
}

// Update Score ...
func (api *Store) UpdateScore() {
	possibleScores := []int{2, 3}

	for i := 0; i < rand.Intn(len(attackCountInAMinute)); i++ {
		player := api.Players[rand.Intn(len(api.Players))]
		assistPlayer := api.Players[rand.Intn(len(api.Players))]

		player.PointInAttack = possibleScores[rand.Intn(len(possibleScores))]
		assistPlayer.Assists++

		if player.Team.Name == homeTeamName {
			api.Statistics.HomeTeamScore += player.PointInAttack
		} else {
			api.Statistics.AwayTeamScore += player.PointInAttack
		}

		player.Points += player.PointInAttack

		for i, item := range api.Players {
			if item.Name == player.Name {
				api.Players[i] = player
			}

			if item.Name == assistPlayer.Name {
				api.Players[i] = assistPlayer
			}
		}

		api.Statistics.TopScorer = findTopScorer(api.Players)
		api.Statistics.TopAssist = findTopAssists(api.Players)
	}

	fmt.Printf("%s: %d, %s: %d, Top Score: %d, Top Scorer: %s, Top Assists: %d, Top Assist Player: %s",
		api.Statistics.HomeTeam,
		api.Statistics.HomeTeamScore,
		api.Statistics.AwayTeam,
		api.Statistics.AwayTeamScore,
		api.Statistics.TopScorer.Points,
		api.Statistics.TopScorer.Name,
		api.Statistics.TopAssist.Assists,
		api.Statistics.TopAssist.Name,
	)
	fmt.Println()
}

// findTopScorer ...
func findTopScorer(players []nba.Player) nba.Player {
	topScorer := players[0]
	for _, player := range players {
		if player.Points > topScorer.Points {
			topScorer = player
		}
	}
	return topScorer
}

// findTopAssists ...
func findTopAssists(players []nba.Player) nba.Player {
	topAssists := players[0]
	for _, player := range players {
		if player.Assists > topAssists.Assists {
			topAssists = player
		}
	}
	return topAssists
}
