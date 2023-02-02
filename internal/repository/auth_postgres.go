package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rostis232/givemetaskbot/internal/entities"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) NewUserRegistration(user *entities.User) error {
	query := fmt.Sprintf("INSERT INTO %s (chat_id, status) VALUES (%d, %d);", UserTable, user.ChatId, user.Status)
	row := a.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return errors.New(fmt.Sprintf("error while creating new user: %s", err))
	}
	return nil
}

func (a *AuthPostgres) GetUser(chatId int64) (entities.User, error) {
	user := entities.User{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE chat_id = %d", UserTable, chatId)
	if err := a.db.Get(&user, query); err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (a *AuthPostgres) UpdateLanguage(user *entities.User) error {
	query := fmt.Sprintf("UPDATE %s SET language = '%s', status = '%d' WHERE chat_id = %d;", UserTable, user.Language, user.Status, user.ChatId)
	row := a.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) UpdateName(user *entities.User) error {
	query := fmt.Sprintf("UPDATE %s SET user_name = '%s', status = '%d' WHERE chat_id = %d;", UserTable, user.UserName, user.Status, user.ChatId)
	row := a.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) UpdateStatus(user *entities.User) error {
	query := fmt.Sprintf("UPDATE %s SET status = %d, active_group = %d, active_task = %d WHERE chat_id = %d;", UserTable, user.Status, user.ActiveGroup, user.ActiveTask, user.ChatId)
	row := a.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) CreateGroup(group *entities.Group) (int, error) {
	var groupId int

	query := fmt.Sprintf("INSERT INTO %s (chief_user_id, group_name) VALUES (%d, '%s')  RETURNING group_id;", GroupTable, group.ChiefUserId, group.GroupName)
	row := a.db.QueryRow(query)
	if err := row.Scan(&groupId); err != nil {
		return 0, err
	}

	return groupId, nil
}

func (a *AuthPostgres) AddEmployeeToGroup(chief, employee *entities.User) error {
	query := fmt.Sprintf("INSERT INTO %s (group_id, employee_user_id) VALUES (%d, %d);", GroupEmployeeTable, chief.ActiveGroup, employee.UserId)
	row := a.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) GetAllChiefsGroups(user *entities.User) ([]entities.Group, error) {
	var allGroups []entities.Group
	query := fmt.Sprintf("SELECT * FROM %s WHERE chief_user_id = %d;", GroupTable, user.UserId)
	err := a.db.Select(&allGroups, query)
	return allGroups, err
}

func (a *AuthPostgres) UpdateGroupName(user *entities.User) error {
	return nil
}
