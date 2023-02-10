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
	query := fmt.Sprintf("INSERT INTO %s (chat_id, status) VALUES ($1, $2);", UserTable)
	row := a.db.QueryRow(query, user.ChatId, user.Status)
	if err := row.Err(); err != nil {
		return errors.New(fmt.Sprintf("error while creating new user: %s", err))
	}
	return nil
}

func (a *AuthPostgres) GetUserByChatId(chatId int64) (entities.User, error) {
	user := entities.User{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE chat_id = $1;", UserTable)
	if err := a.db.Get(&user, query, chatId); err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (a *AuthPostgres) GetUserByUserId(userId int64) (entities.User, error) {
	user := entities.User{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1;", UserTable)
	if err := a.db.Get(&user, query, userId); err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (a *AuthPostgres) UpdateLanguage(user *entities.User) error {
	query := fmt.Sprintf("UPDATE %s SET language = $1, status = $2 WHERE chat_id = $3;", UserTable)
	row := a.db.QueryRow(query, user.Language, user.Status, user.ChatId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) UpdateName(user *entities.User) error {
	query := fmt.Sprintf("UPDATE %s SET user_name = $1, status = $2 WHERE chat_id = $3;", UserTable)
	row := a.db.QueryRow(query, user.UserName, user.Status, user.ChatId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) UpdateStatus(user *entities.User) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1, active_group = $2, active_task = $3 WHERE chat_id = $4;", UserTable)
	row := a.db.QueryRow(query, user.Status, user.ActiveGroup, user.ActiveTask, user.ChatId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) CreateGroup(group *entities.Group) (int, error) {
	var groupId int
	query := fmt.Sprintf("INSERT INTO %s (chief_user_id, group_name) VALUES ($1, $2)  RETURNING group_id;", GroupTable)
	row := a.db.QueryRow(query, group.ChiefUserId, group.GroupName)
	if err := row.Scan(&groupId); err != nil {
		return 0, err
	}

	return groupId, nil
}

func (a *AuthPostgres) AddEmployeeToGroup(groupID, employeeID int) error {
	query := fmt.Sprintf("INSERT INTO %s (group_id, employee_user_id) VALUES ($1, $2);", GroupEmployeeTable)
	row := a.db.QueryRow(query, groupID, employeeID)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) GetAllChiefsGroups(user *entities.User) ([]entities.Group, error) {
	var allGroups []entities.Group
	query := fmt.Sprintf("SELECT * FROM %s WHERE chief_user_id = $1;", GroupTable)
	err := a.db.Select(&allGroups, query, user.UserId)
	return allGroups, err
}

func (a *AuthPostgres) GetAllEmployeeGroups(user *entities.User) ([]entities.Group, error) {
	//TODO: Change SQL request!!!
	var allGroups []entities.Group
	query := fmt.Sprintf("SELECT gt.group_id, gt.chief_user_id, gt.group_name FROM %s gt INNER JOIN %s ge ON gt.group_id = ge.group_id WHERE employee_user_id = %d;", GroupTable, GroupEmployeeTable, user.UserId)
	err := a.db.Select(&allGroups, query)
	return allGroups, err
}

func (a *AuthPostgres) UpdateGroupName(user *entities.User, newGroupName string) error {
	query := fmt.Sprintf("UPDATE %s SET group_name = $1 WHERE group_id = $2;", GroupTable)
	row := a.db.QueryRow(query, newGroupName, user.ActiveGroup)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) ShowAllEmploysFromGroup(user *entities.User) ([]entities.User, error) {
	//TODO: Change SQL request!!!
	var allEmployees []entities.User
	query := fmt.Sprintf("SELECT ut.user_id, ut.chat_id, ut.user_name, ut.language, ut.status, ut.active_group, ut.active_task FROM %s ge INNER JOIN %s ut ON ge.employee_user_id = ut.user_id WHERE group_id = %d;", GroupEmployeeTable, UserTable, user.ActiveGroup)
	err := a.db.Select(&allEmployees, query)
	return allEmployees, err
}

func (a *AuthPostgres) GetGroupById(id int) (entities.Group, error) {
	group := entities.Group{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE group_id = $1;", GroupTable)
	if err := a.db.Get(&group, query, id); err != nil {
		return entities.Group{}, err
	}

	return group, nil
}

func (a *AuthPostgres) DeleteEmployeeFromGroup(employee *entities.User, group *entities.Group) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE employee_user_id = $1 AND group_id = $2;", GroupEmployeeTable)
	row := a.db.QueryRow(query, employee.UserId, group.Id)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *AuthPostgres) GetGroupsWithoutSelectedEmployee(id int) ([]entities.Group, error) {
	var allGroups []entities.Group
	query := fmt.Sprintf("SELECT * FROM %s WHERE group_id NOT IN (SELECT group_id FROM %s WHERE employee_user_id = $1 GROUP BY group_id);", GroupTable, GroupEmployeeTable)
	err := a.db.Select(&allGroups, query, id)
	return allGroups, err
}
