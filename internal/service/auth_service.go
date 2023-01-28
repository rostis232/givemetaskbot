package service

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/keyboards"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/repository"
	"github.com/rostis232/givemetaskbot/internal/state_service"
	"log"
)

type AuthService struct {
	repository repository.Authorisation
}

func NewAuthService(repository repository.Authorisation) *AuthService {
	return &AuthService{repository: repository}
}

func (u *AuthService) GetUser(chatId int64) (entities.User, error) {
	user, err := u.repository.GetUser(chatId)
	return user, err
}

func (u *AuthService) NewUserRegistration(chatId int64) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(chatId, "")
	user := entities.User{
		ChatId: chatId,
		Status: state_service.Expecting_language,
	}
	if err := u.repository.NewUserRegistration(user); err != nil {
		return msg, err
	}

	msg.Text = "Choose your language:"
	msg.ReplyMarkup = keyboards.LanguageKeyboard

	return msg, nil
}

func (u *AuthService) SelectLanguage(user entities.User, lng messages.Language) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(user.ChatId, "")
	err := errors.New("")

	switch {
	case user.Status == state_service.Expecting_language:
		user.Status = state_service.Expecting_new_user_name
		user.Language = lng
		// user to repo
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageAfterFirstLanguageSelection, user.Language)
		if err != nil && msg.Text != "" {
			log.Println(err)
			return msg, nil
		} else if err != nil {
			return tgbotapi.MessageConfig{}, err
		} else {
			return msg, nil
		}

	}

	return msg, nil
}
