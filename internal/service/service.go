package service

import (
	"github.com/rostis232/givemetaskbot/internal/repository"
)

type Authorisation interface {
	CheckIfUserIsRegistered(chatId int64) bool
	InsertUserToUserStates(chatId int64)
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
