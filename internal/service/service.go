package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/repository"
)

type Authorisation interface {
	GetUser(chatId int64) (entities.User, error)
	NewUserRegistration(chatId int64) (tgbotapi.MessageConfig, error)
	SelectLanguage(user *entities.User, lng messages.Language) (tgbotapi.MessageConfig, error)
	SetUserName(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	ShowChatId(user *entities.User) (tgbotapi.MessageConfig, error)
	UserNameChanging(user *entities.User) (tgbotapi.MessageConfig, error)
	MainMenu(user *entities.User) (tgbotapi.MessageConfig, error)
	AskingForNewGroupTitle(user *entities.User) (tgbotapi.MessageConfig, error)
	CreatingNewGroup(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	AddingEmployeeToGroup(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	GroupsMenu(user *entities.User) (tgbotapi.MessageConfig, error)
	ShowAllChiefsGroups(user *entities.User) (tgbotapi.MessageConfig, error)
	AskingForUpdatedGroupName(user *entities.User, groupId int) (tgbotapi.MessageConfig, error)
	UpdateGroupName(user *entities.User, newGroupName string) (tgbotapi.MessageConfig, error)
	ShowAllEmploysFromGroup(user *entities.User, groupId int) (tgbotapi.MessageConfig, error)
}

type Service struct {
	Authorisation
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
	}
}
