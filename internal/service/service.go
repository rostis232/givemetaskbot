package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/repository"
)

type Authorisation interface {
	GetUser(chatId int64) (entities.User, error)
	NewUserRegistration(chatId int64) (tgbotapi.MessageConfig, error)
	SelectLanguage(user *entities.User, callbackQueryData string) error
	SetUserName(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	ShowChatId(user *entities.User) error
	UserNameChanging(user *entities.User) error
	MainMenu(user *entities.User) error
	UserMenuSettings(user *entities.User) error
	ChangeLanguageMenu(user *entities.User) error
	AskingForNewGroupTitle(user *entities.User) error
	CreatingNewGroup(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	AddingEmployeeToGroup(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	GroupsMenu(user *entities.User) error
	ShowAllChiefsGroups(user *entities.User) error
	ShowAllEmployeeGroups(user *entities.User) error
	AskingForUpdatedGroupName(user *entities.User, callbackQueryData string) error
	UpdateGroupName(user *entities.User, newGroupName string) (tgbotapi.MessageConfig, error)
	ShowAllEmploysFromGroup(user *entities.User, callbackQueryData string) error
	DeleteEmployeeFromGroup(user *entities.User, receivedData string) error
}

type Service struct {
	Authorisation
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
	}
}
