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
	AddEmployeeToGroup(groupID, employeeID int) error
	GetAllChiefsGroups(user *entities.User) ([]entities.Group, error)
	GetAllEmployeeGroups(user *entities.User) ([]entities.Group, error)
	UpdateGroupName(user *entities.User, newGroupName string) error
	ShowAllEmploysFromGroup(groupID int) ([]entities.User, error)
	GetGroupById(id int) (entities.Group, error)
	DeleteEmployeeFromGroup(employee *entities.User, group *entities.Group) error
	GetGroupsWithoutSelectedEmployee(id int) ([]entities.Group, error)
	DeleteGroup(id int) error
	LeaveGroup(employeeID, groupID int) error
	CreateNewTask(taskTitle string, groupID int) (int, error)
	UpdateTaskDescription(taskDesc string, taskID int) error
	GetTaskByID(taskID int) (entities.Task, error)
}

type Repository struct {
	Authorisation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
