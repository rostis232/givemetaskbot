package service

import (
	"errors"
	"fmt"
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
	if err := u.repository.NewUserRegistration(&user); err != nil {
		return msg, err
	}

	msg.Text = "Choose your language:"
	msg.ReplyMarkup = keyboards.LanguageKeyboard

	return msg, nil
}

func (u *AuthService) SelectLanguage(user *entities.User, lng messages.Language) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(user.ChatId, "")
	err := errors.New("")

	switch {
	case user.Status == state_service.Expecting_language:
		user.Status = state_service.Expecting_new_user_name

		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageAfterFirstLanguageSelection, lng)
		if err != nil {
			log.Println(err)
		}
	default:
		user.Status = state_service.MainMenu

		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageAfterLanguageUpdate, lng)
		if err != nil {
			log.Println(err)
		}

	}
	user.Language = lng
	if err := u.repository.UpdateLanguage(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	return msg, nil
}

func (u *AuthService) SetUserName(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(user.ChatId, "")
	err := errors.New("")
	text := ""
	user.UserName = message.Text

	switch {
	case user.Status == state_service.Expecting_new_user_name:
		text, err = messages.ReturnMessageByLanguage(messages.MessageAfterFirstNameEntering, user.Language)
		msg.Text = fmt.Sprintf(text, user.UserName)
		msg.ReplyMarkup = keyboards.NewKeyboardChooseCreateOrJoinGroup(user)
	default:
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageAfterLanguageUpdate, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	}

	user.Status = state_service.MainMenu

	if err := u.repository.UpdateName(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	return msg, err
}
