package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/repository"
)

type Authorisation interface {
	GetUser(chatId int64) (entities.User, error)
	NewUserRegistration(chatId int64) (tgbotapi.MessageConfig, error)
	SelectLanguage(user *entities.User, lng messages.Language) (tgbotapi.MessageConfig, error)
	SetUserName(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
}

type Group interface {
}

type Task interface {
}

type Service struct {
	Authorisation
	Group
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
		Group:         NewGroupService(repo.Group),
		Task:          NewTaskService(repo.Task),
	}
}
