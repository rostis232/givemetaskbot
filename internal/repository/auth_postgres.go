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

func (a *AuthPostgres) NewUserRegistration(user entities.User) error {
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
