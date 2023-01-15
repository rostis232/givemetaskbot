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
	CreateUser(user entities.User) error
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