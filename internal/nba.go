package nba

type Player struct {
	Name          string
	Assists       int
	Points        int
	Team          Team
	PointInAttack int
}

type Statistic struct {
	HomeTeamScore int
	AwayTeamScore int
	HomeTeam      Team
	AwayTeam      Team
	TopAssist     Player
	TopScorer     Player
}

type Team struct {
	Name string
}
