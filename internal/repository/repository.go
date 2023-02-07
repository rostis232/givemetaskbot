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
	GetUserByChatId(chatId int64) (entities.User, error)
	GetUserByUserId(userId int64) (entities.User, error)
	UpdateLanguage(user *entities.User) error
	UpdateName(user *entities.User) error
	UpdateStatus(user *entities.User) error
	CreateGroup(group *entities.Group) (int, error)
	AddEmployeeToGroup(chief, employee *entities.User) error
	GetAllChiefsGroups(user *entities.User) ([]entities.Group, error)
	UpdateGroupName(user *entities.User, newGroupName string) error
	ShowAllEmploysFromGroup(user *entities.User) ([]entities.User, error)
	GetGroupById(id int) (entities.Group, error)
	DeleteEmployeeFromGroup(employee *entities.User, group *entities.Group) error
}

type Repository struct {
	Authorisation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
