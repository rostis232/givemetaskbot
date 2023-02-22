package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/repository"
)

type Authorisation interface {
	GetUser(chatId int64) (entities.User, error)
	NewUserRegistration(chatId int64) error
	SelectLanguage(user *entities.User, callbackQueryData string) error
	SetUserName(user *entities.User, message *tgbotapi.Message) error
	ShowChatId(user *entities.User) error
	UserNameChanging(user *entities.User) error
	MainMenu(user *entities.User) error
	UserMenuSettings(user *entities.User) error
	ChangeLanguageMenu(user *entities.User) error
	AskingForNewGroupTitle(user *entities.User) error
	CreatingNewGroup(user *entities.User, message *tgbotapi.Message) error
	AddingEmployeeToGroup(user *entities.User, message *tgbotapi.Message) error
	GroupsMenu(user *entities.User) error
	ShowAllChiefsGroups(user *entities.User) error
	ShowAllEmployeeGroups(user *entities.User) error
	AskingForUpdatedGroupName(user *entities.User, callbackQueryData string) error
	UpdateGroupName(user *entities.User, newGroupName string) error
	ShowAllEmploysFromGroup(user *entities.User, callbackQueryData string) error
	WarningBeforeDeletingEmployeeFromGroup(user *entities.User, callbackQueryData string) error
	DeleteEmployeeFromGroup(user *entities.User, receivedData string) error
	CopyEmployeeToAnotherGroup(user *entities.User, callbackQueryData string) error
	ConfirmCopyEmployeeToAnotherGroup(user *entities.User, callbackQueryData string) error
	WarningBeforeGroupDeleting(user *entities.User, callbackQueryData string) error
	DeleteGroup(user *entities.User, callbackQueryData string) error
	LeaveGroupBeforeConfirmation(user *entities.User, callbackQueryData string) error
	LeaveGroupWithConfirmation(user *entities.User) error
	CreateNewTaskAskingTitle(user *entities.User, callbackQueryData string) error
	CreateNewTaskGotTitle(user *entities.User, taskTitle string) error
	UpdateTaskDescription(user *entities.User, taskDesc string) error
	AddAllEmployeesToTask(user *entities.User) error
	AddSomeEmployeesToTaskShowList(user *entities.User, callbackQueryData string) error
	AssignEmployee(user *entities.User, callbackQueryData string) error
	SkipDescritionEntering(user *entities.User) error
	ShowTaskDetailsForEmployee(user *entities.User, callbackQueryData string) error
	ShowTaskDetailsForChief(user *entities.User, callbackQueryData string) error
	ShowAllGroupTasksForChief(user *entities.User, callbackQueryData string) error
	ShowAllGroupTasksForEmployee(user *entities.User, callbackQueryData string) error
	MarkTaskAsComplete(user *entities.User, callbackQueryData string) error
	DeleteTask(user *entities.User, callbackQueryData string) error 
	DeleteExecutors(user *entities.User, callbackQueryData string) error
}

type Service struct {
	Authorisation
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
	}
}
