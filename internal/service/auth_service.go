package service

import (
	"database/sql"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/keyboards"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/repository"
	"github.com/rostis232/givemetaskbot/internal/state_service"
	"log"
	"strconv"
	"strings"
)

var MsgChan = make(chan tgbotapi.MessageConfig)

type AuthService struct {
	repository repository.Authorisation
}

func NewAuthService(repository repository.Authorisation) *AuthService {
	return &AuthService{repository: repository}
}

func (u *AuthService) GetUser(chatId int64) (entities.User, error) {
	user, err := u.repository.GetUserByChatId(chatId)
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
	case user.Status == state_service.Changing_user_name:
		text, err = messages.ReturnMessageByLanguage(messages.MessageAfterUserNameUpdate, user.Language)
		msg.Text = fmt.Sprintf(text, user.UserName)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	}

	user.Status = state_service.MainMenu

	if err := u.repository.UpdateName(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	return msg, err
}

func (u *AuthService) ShowChatId(user *entities.User) (tgbotapi.MessageConfig, error) {
	text, err := messages.ReturnMessageByLanguage(messages.MessageWithChatId, user.Language)
	if err != nil {
		log.Println(err)
	}
	code := fmt.Sprintf(text, keys.ChatIdPrefix+strconv.Itoa(int(user.ChatId+keys.ChatIdSuffix)))
	msg := tgbotapi.NewMessage(user.ChatId, code)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)

	return msg, nil
}

func (u *AuthService) UserNameChanging(user *entities.User) (tgbotapi.MessageConfig, error) {
	title, err := messages.ReturnMessageByLanguage(messages.MessageBeforeNameChanging, user.Language)
	if err != nil {
		log.Println(err)
	}
	user.Status = state_service.Changing_user_name
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}
	msg := tgbotapi.NewMessage(user.ChatId, title)

	return msg, nil

}

func (u *AuthService) MainMenu(user *entities.User) (tgbotapi.MessageConfig, error) {
	title, err := messages.ReturnMessageByLanguage(messages.MainMenuTitle, user.Language)
	if err != nil {
		log.Println(err)
	}

	user.Status = state_service.MainMenu
	user.ActiveGroup = 0
	user.ActiveTask = 0

	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.NewMainMenuKeyboard(user)

	return msg, nil
}

func (u *AuthService) AskingForNewGroupTitle(user *entities.User) (tgbotapi.MessageConfig, error) {
	title, err := messages.ReturnMessageByLanguage(messages.MessageEnterNewGroupName, user.Language)
	if err != nil {
		log.Println(err)
	}
	user.Status = state_service.Expecting_new_group_name

	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)

	return msg, nil
}

func (u *AuthService) CreatingNewGroup(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	group := entities.Group{
		ChiefUserId: user.UserId,
		GroupName:   message.Text,
	}
	title, err := messages.ReturnMessageByLanguage(messages.MessageCreatedNewGroup, user.Language)
	if err != nil {
		log.Println(err)
	}
	groupId, err := u.repository.CreateGroup(&group)
	if err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}
	user.Status = state_service.Adding_employee_to_group
	user.ActiveGroup = groupId

	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(title, group.GroupName))
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)

	return msg, nil
}

func (u *AuthService) AddingEmployeeToGroup(user *entities.User, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(user.ChatId, messages.UnknownError)
	err := errors.New("")
	//Перевірка чи має код префікс
	_, codeWithoutPrefix, cutting := strings.Cut(message.Text, keys.ChatIdPrefix)
	if !cutting {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageEmployeeCodeBroken, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			return msg, err
		}
		return msg, nil
	}
	//Конвертація коду
	codeWithSuffix, err := strconv.Atoi(codeWithoutPrefix)
	if err != nil {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageEmployeeCodeBroken, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			return msg, err
		}
		return msg, nil
	}
	cleanCode := codeWithSuffix - keys.ChatIdSuffix
	employee, err := u.repository.GetUserByChatId(int64(cleanCode))
	//Перевірка, чи є користувачі з таким кодом
	if err == sql.ErrNoRows {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageNoEmployeeWithThisCode, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			return msg, err
		}
		return msg, err
	}
	if err != nil {
		return msg, err
	}
	//Перевірка чи не ввів користувач свій код
	if user.ChatId == employee.ChatId {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageEmployeeCodeEqualsChiefCode, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			return msg, err
		}
		return msg, nil
	}
	//Перевірка, чи є у користувача активна група в статусі
	if user.ActiveGroup == 0 {
		if user.ActiveGroup == 0 {
			text, err := messages.ReturnMessageByLanguage(messages.UnknownError, user.Language)
			msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
			if err != nil {
				log.Println(err)
			}
			msg.Text = fmt.Sprintf(text, "Wrong Active Group")
		}
	}

	//Передати в репозиторій
	if err := u.repository.AddEmployeeToGroup(user, &employee); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	text, err := messages.ReturnMessageByLanguage(messages.MessageEmployeeAddingSuccess, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg.Text = fmt.Sprintf(text, employee.UserName)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	return msg, nil
}

func (u *AuthService) GroupsMenu(user *entities.User) (tgbotapi.MessageConfig, error) {
	title, err := messages.ReturnMessageByLanguage(messages.GroupMenuTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.NewGroupMenuKeyboard(user)

	return msg, nil
}

func (u *AuthService) ShowAllChiefsGroups(user *entities.User) (tgbotapi.MessageConfig, error) {
	allGroups, err := u.repository.GetAllChiefsGroups(user)
	if err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	if allGroups == nil {
		//Якщо немає груп пишемо відповідне повідомлення
		text, err := messages.ReturnMessageByLanguage(messages.NoGroups, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, text)
		msg.ReplyMarkup = keyboards.NewGroupCreatingKeyboard(user)
		MsgChan <- msg
	} else {
		for _, g := range allGroups {
			msg := tgbotapi.NewMessage(user.ChatId, g.GroupName)
			msg.ReplyMarkup = keyboards.NewMenuForEvenGroup(user, &g)
			MsgChan <- msg
		}
	}

	return tgbotapi.MessageConfig{}, err
}

func (u *AuthService) AskingForUpdatedGroupName(user *entities.User, groupId int) (tgbotapi.MessageConfig, error) {
	text, err := messages.ReturnMessageByLanguage(messages.RenameGroupTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	user.ActiveGroup = groupId
	user.Status = state_service.Changing_group_name
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)

	return msg, err

}

func (u *AuthService) UpdateGroupName(user *entities.User, newGroupName string) (tgbotapi.MessageConfig, error) {
	if err := u.repository.UpdateGroupName(user, newGroupName); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	text, err := messages.ReturnMessageByLanguage(messages.MessageNewGroupNameAccepted, user.Language)
	if err != nil {
		log.Println(err)
	}
	text = fmt.Sprintf(text, newGroupName)
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)

	return msg, nil
}

func (u *AuthService) ShowAllEmploysFromGroup(user *entities.User, groupId int) (tgbotapi.MessageConfig, error) {
	user.Status = state_service.Adding_employee_to_group
	user.ActiveGroup = groupId
	allEmployees, err := u.repository.ShowAllEmploysFromGroup(user)
	if err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return tgbotapi.MessageConfig{}, err
	}

	group, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		log.Println(err)
	}

	if allEmployees == nil {
		//Якщо немає користувачів пишемо відповідне повідомлення
		text, err := messages.ReturnMessageByLanguage(messages.NoEmployeesInTheGroup, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(text, group.GroupName))
		MsgChan <- msg
	} else {
		text, err := messages.ReturnMessageByLanguage(messages.ShownMembersOfTheGroupWithName, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(text, group.GroupName))
		MsgChan <- msg

		//В іншому випадку направляємо перелік користувачів
		for _, e := range allEmployees {
			msg := tgbotapi.NewMessage(user.ChatId, e.UserName)
			msg.ReplyMarkup = keyboards.NewMenuForEvenEmployee(user, &e)
			MsgChan <- msg
		}
	}

	//надсилаємо останнє повідомлення, в якому пропонуємо ввести нового користувача або повернутись до головного меню
	text, err := messages.ReturnMessageByLanguage(messages.AddNewEmployeeToExistingGroup, user.Language)
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(text, group.GroupName))
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return tgbotapi.MessageConfig{}, err
}

func (u *AuthService) DeleteEmployeeFromGroup(user *entities.User, receivedData string) (tgbotapi.MessageConfig, error) {
	_, employeeUserIdString, ok := strings.Cut(receivedData, keys.EmployeeIDtoDeleteFromGroup)
	if !ok {
		log.Println("Помилка отримання chatID працівника")
	}
	group, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		log.Println(err)
	}
	employeeUserIdInt, err := strconv.Atoi(employeeUserIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
	}
	employee, err := u.repository.GetUserByUserId(int64(employeeUserIdInt))
	if err != nil {
		log.Println("Помилка отримання даних працівника з БД")
	}

	if err := u.repository.DeleteEmployeeFromGroup(&employee, &group); err != nil {
		log.Printf("Error while deleting user from group: %s", err)
		return tgbotapi.MessageConfig{}, err
	}
	messageToEmployee, err := messages.ReturnMessageByLanguage(messages.YouAreDeletedFromGroup, employee.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToEmployee := tgbotapi.NewMessage(employee.ChatId, fmt.Sprintf(messageToEmployee, group.GroupName))
	MsgChan <- msgToEmployee
	messageToChief, err := messages.ReturnMessageByLanguage(messages.EmployeeHaveBeenDeletedFromGroup, employee.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(messageToChief, employee.UserName, group.GroupName))
	MsgChan <- msgToChief
	return tgbotapi.MessageConfig{}, nil
}
