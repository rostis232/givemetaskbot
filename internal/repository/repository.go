package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rostis232/givemetaskbot/internal/entities"
)

const (
	UserTable          = "users"
	GroupTable         = "groups"
	TaskTable          = "tasks"
	GroupEmployeeTable = "group_employee"
)

type Authorisation interface {
	NewUserRegistration(user *entities.User) error
	GetUser(chatId int64) (entities.User, error)
	UpdateLanguage(user *entities.User) error
	UpdateName(user *entities.User) error
	UpdateStatus(user *entities.User) error
	CreateGroup(group *entities.Group) (int, error)
	AddEmployeeToGroup(chief, employee *entities.User) error
}

type Repository struct {
	Authorisation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
