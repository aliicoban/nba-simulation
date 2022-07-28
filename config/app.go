package config

type App struct {
	Name    string
	Env     string
	Port    string
	BaseUrl string
}

func NewApp() App {
	return App{
		Name:    Getenv("APP_NAME", "nba"),
		Env:     Getenv("APP_ENV", "local"),
		Port:    Getenv("PORT", "4444"),
		BaseUrl: Getenv("BASE_URL", "http://localhost"),
	}
}
