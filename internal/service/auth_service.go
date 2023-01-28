package service

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/keyboards"
	"github.com/rostis232/givemetaskbot/internal/repository"
	"log"
)

type AuthService struct {
	repository repository.Authorisation
}

func NewAuthService(repository repository.Authorisation) *AuthService {
	return &AuthService{repository: repository}
}

func (u *AuthService) Registration(chatId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "")

	user, err := u.repository.GetUser(chatId)
	switch {
	case err == sql.ErrNoRows:
		//If user isn`t in DB, we add his chatId with unregistered status into user state machine
		newUser := entities.User{ChatId: chatId, Status: entities.UserUnregistered}
		if err := u.repository.CreateUser(newUser); err != nil {
			log.Println(err)
		}
		msg.Text = "Для реєстрації введіть ім'я, яке буде видно пов'язаним учасникам цього боту"
		msg.ReplyMarkup = keyboards.RemoveKeyboard
		return msg
	case user.UserName == "" && user.Status == entities.UserUnregistered:
		// If user is in DB, then we add his chatId, user struct with status don`t need
		//registration to user state machine
		msg.Text = "Для реєстрації введіть ім'я, яке буде видно пов'язаним учасникам цього боту"
		msg.ReplyMarkup = keyboards.RemoveKeyboard
		return msg
	default:
		msg.Text = "Ви вже зареєстрований користувач, або виникла непередбачувана помилка"
		msg.ReplyMarkup = keyboards.RemoveKeyboard
		return msg
	}
}
