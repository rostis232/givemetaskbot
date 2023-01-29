package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rostis232/givemetaskbot/internal/entities"
)

const (
	UserTable  = "users"
	GroupTable = "groups"
	TaskTable  = "tasks"
)

type Authorisation interface {
	NewUserRegistration(user *entities.User) error
	GetUser(chatId int64) (entities.User, error)
	UpdateLanguage(user *entities.User) error
	UpdateName(user *entities.User) error
	UpdateStatus(user *entities.User) error
}

type Group interface {
}

type Task interface {
}

type Repository struct {
	Authorisation
	Group
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		Group:         NewGroupPostgres(db),
		Task:          NewTaskPostgres(db),
	}
}
