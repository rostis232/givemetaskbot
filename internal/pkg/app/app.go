package app

import (
	"github.com/joho/godotenv"
	"github.com/rostis232/givemetaskbot/internal/repository"
	"github.com/rostis232/givemetaskbot/internal/service"
	"github.com/rostis232/givemetaskbot/internal/telegram"
	"github.com/spf13/viper"
	"os"
)

type Bot interface {
	Start() error
}

type Service interface {
	CheckIfUserIsRegistered(chatId int64) bool
}

type Repository interface {
}

type App struct {
	bot        Bot
	repository Repository
	service    Service
}

func NewApp() (*App, error) {
	if err := initConfig(); err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	a := &App{}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		return nil, err
	}

	a.repository = repository.NewRepository(db)

	a.service = service.NewService(a.repository)

	bot, err := telegram.NewBot(os.Getenv("TOKEN"), a.service)
	if err != nil {
		return nil, err
	}

	a.bot = bot

	return a, nil

}

// Run starts telegram bot using a Start method of Bot struct
func (a *App) Run() error {
	return a.bot.Start()
}

// initConfig initializes viper config file
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
