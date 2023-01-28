package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/repository"
	"github.com/rostis232/givemetaskbot/internal/user_state"
)

type Authorisation interface {
	Registration(chatId int64) tgbotapi.MessageConfig
}

type Group interface {
}

type Task interface {
}

type StateService interface {
}

type Service struct {
	Authorisation
	Group
	Task
	StateService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
		Group:         NewGroupService(repo.Group),
		Task:          NewTaskService(repo.Task),
		StateService:  user_state.NewStateService(),
	}
}
