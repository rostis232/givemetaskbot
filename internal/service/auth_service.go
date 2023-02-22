package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/keyboards"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/repository"
	"github.com/rostis232/givemetaskbot/internal/state_service"
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

func (u *AuthService) NewUserRegistration(chatId int64) error {
	user := entities.User{
		ChatId: chatId,
		Status: state_service.Expecting_language,
	}
	if err := u.repository.NewUserRegistration(&user); err != nil {
		log.Printf("Error while new user registration: %s", err)
		return err
	}
	msg := tgbotapi.NewMessage(chatId, "Choose your language:\nОберіть мову:")
	msg.ReplyMarkup = keyboards.LanguageKeyboard
	MsgChan <- msg
	return nil
}

// //It`ll be  need to be used when implementing time control
func (u *AuthService) SelectLanguage(user *entities.User, callbackQueryData string) error {
	msg := tgbotapi.NewMessage(user.ChatId, "")
	var err error

	_, lng, ok := strings.Cut(callbackQueryData, ":")
	if !ok {
		log.Println("error while cutting string with language")
	}

	switch {
	case user.Status == state_service.Expecting_language:
		user.Status = state_service.Expecting_new_user_name

		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageAfterFirstLanguageSelection, messages.Language(lng))
		if err != nil {
			log.Println(err)
		}
	default:
		user.Status = state_service.MainMenu

		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageAfterLanguageUpdate, messages.Language(lng))
		if err != nil {
			log.Println(err)
		}
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	}
	user.Language = messages.Language(lng)
	if err := u.repository.UpdateLanguage(user); err != nil {
		log.Println(err)
		return err
	}

	MsgChan <- msg
	return nil
}

func (u *AuthService) SetUserName(user *entities.User, message *tgbotapi.Message) error {
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
		return err
	}
	MsgChan <- msg
	return err
}

func (u *AuthService) ShowChatId(user *entities.User) error {
	text, err := messages.ReturnMessageByLanguage(messages.MessageWithChatId, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg := tgbotapi.NewMessage(user.ChatId, text)
	MsgChan <- msg

	code := keys.ChatIdPrefix + strconv.Itoa(int(user.ChatId+keys.ChatIdSuffix))
	msg = tgbotapi.NewMessage(user.ChatId, code)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)

	MsgChan <- msg

	return nil
}

func (u *AuthService) UserNameChanging(user *entities.User) error {
	title, err := messages.ReturnMessageByLanguage(messages.MessageBeforeNameChanging, user.Language)
	if err != nil {
		log.Println(err)
	}
	user.Status = state_service.Changing_user_name
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}
	msg := tgbotapi.NewMessage(user.ChatId, title)
	MsgChan <- msg
	return nil
}

func (u *AuthService) MainMenu(user *entities.User) error {
	title, err := messages.ReturnMessageByLanguage(messages.MainMenuTitle, user.Language)
	if err != nil {
		log.Println(err)
	}

	user.Status = state_service.MainMenu
	user.ActiveGroup = 0
	user.ActiveTask = 0

	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}

	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.NewMainMenuKeyboard(user)

	MsgChan <- msg

	return nil
}

func (u *AuthService) UserMenuSettings(user *entities.User) error {
	title, err := messages.ReturnMessageByLanguage(messages.UserSettingsMenuTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.NewUserSettingsMenuKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) ChangeLanguageMenu(user *entities.User) error {
	title, err := messages.ReturnMessageByLanguage(messages.ChangeLanguageKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.LanguageKeyboard
	MsgChan <- msg
	return nil
}

func (u *AuthService) AskingForNewGroupTitle(user *entities.User) error {
	title, err := messages.ReturnMessageByLanguage(messages.MessageEnterNewGroupName, user.Language)
	if err != nil {
		log.Println(err)
	}
	user.Status = state_service.Expecting_new_group_name

	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}

	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) CreatingNewGroup(user *entities.User, message *tgbotapi.Message) error {
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
		return err
	}
	user.Status = state_service.Adding_employee_to_group
	user.ActiveGroup = groupId

	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}

	msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(title, group.GroupName))
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) AddingEmployeeToGroup(user *entities.User, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(user.ChatId, messages.UnknownError)
	var err error
	//Перевірка чи має код префікс
	_, codeWithoutPrefix, cutting := strings.Cut(message.Text, keys.ChatIdPrefix)
	if !cutting {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageEmployeeCodeBroken, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			MsgChan <- msg
			return err
		}
		MsgChan <- msg
		return nil
	}
	//Конвертація коду
	codeWithSuffix, err := strconv.Atoi(codeWithoutPrefix)
	if err != nil {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageEmployeeCodeBroken, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			MsgChan <- msg
			return err
		}
		MsgChan <- msg
		return nil
	}
	cleanCode := codeWithSuffix - keys.ChatIdSuffix
	employee, err := u.repository.GetUserByChatId(int64(cleanCode))
	//Перевірка, чи є користувачі з таким кодом
	if err == sql.ErrNoRows {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageNoEmployeeWithThisCode, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			MsgChan <- msg
			return err
		}
		MsgChan <- msg
		return err
	}
	if err != nil {
		MsgChan <- msg
		return err
	}
	//Перевірка чи не ввів користувач свій код
	if user.ChatId == employee.ChatId {
		msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageEmployeeCodeEqualsChiefCode, user.Language)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		if err != nil {
			log.Println(err)
			MsgChan <- msg
			return err
		}
		MsgChan <- msg
		return nil
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
	if err := u.repository.AddEmployeeToGroup(user.ActiveGroup, int(employee.UserId)); err != nil {
		log.Println(err)
		return err
	}

	text, err := messages.ReturnMessageByLanguage(messages.MessageEmployeeAddingSuccess, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg.Text = fmt.Sprintf(text, employee.UserName)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) GroupsMenu(user *entities.User) error {
	title, err := messages.ReturnMessageByLanguage(messages.GroupMenuTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg := tgbotapi.NewMessage(user.ChatId, title)
	msg.ReplyMarkup = keyboards.NewGroupMenuKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) ShowAllChiefsGroups(user *entities.User) error {
	user.Status = state_service.Expecting_new_group_name
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}
	allGroups, err := u.repository.GetAllChiefsGroups(user)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(allGroups) == 0 {
		//Якщо немає груп пишемо відповідне повідомлення
		text, err := messages.ReturnMessageByLanguage(messages.NoGroups, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, text)
		MsgChan <- msg
	} else {
		for _, g := range allGroups {
			msg := tgbotapi.NewMessage(user.ChatId, g.GroupName)
			msg.ReplyMarkup = keyboards.NewMenuForEvenGroup(user, int(g.Id))
			MsgChan <- msg
		}
	}
	//Пропозиція додати нову групу чи повернутись до головного меню
	text, err := messages.ReturnMessageByLanguage(messages.AddNewGroupFromGroupListMenu, user.Language)
	if err != nil {
		log.Println(err)
	}
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return err
}

func (u *AuthService) ShowAllEmployeeGroups(user *entities.User) error {
	allGroups, err := u.repository.GetAllEmployeeGroups(user)
	if err != nil {
		log.Println(err)
		return err
	}
	if len(allGroups) == 0 {
		//Якщо немає груп пишемо відповідне повідомлення
		text, err := messages.ReturnMessageByLanguage(messages.NoGroups, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, text)
		MsgChan <- msg
	} else {
		for _, g := range allGroups {
			msg := tgbotapi.NewMessage(user.ChatId, g.GroupName)
			msg.ReplyMarkup = keyboards.NewMenuForEvenGroupForEmployee(user.Language, g.Id)
			MsgChan <- msg
		}
	}
	return nil
}

func (u *AuthService) AskingForUpdatedGroupName(user *entities.User, callbackQueryData string) error {
	_, groupId, ok := strings.Cut(callbackQueryData, keys.RenameGroupWithId)
	if !ok {
		log.Println("Error while getting group ID")
	}
	groupIdInt, err := strconv.Atoi(groupId)
	if err != nil {
		log.Printf("Error while converting group ID from string to int: %s", err)
	}

	text, err := messages.ReturnMessageByLanguage(messages.RenameGroupTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	user.ActiveGroup = groupIdInt
	user.Status = state_service.Changing_group_name
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return err
}

func (u *AuthService) UpdateGroupName(user *entities.User, newGroupName string) error {
	oldGroup, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		log.Println(err)
		return err
	}

	allEmployees, err := u.repository.ShowAllEmploysFromGroup(user.ActiveGroup)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := u.repository.UpdateGroupName(user, newGroupName); err != nil {
		log.Println(err)
		return err
	}

	text, err := messages.ReturnMessageByLanguage(messages.MessageNewGroupNameAccepted, user.Language)
	if err != nil {
		log.Println(err)
	}
	text = fmt.Sprintf(text, oldGroup.GroupName, newGroupName)
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg

	if len(allEmployees) != 0 {
		for _, employee := range allEmployees {
			messageForEmployee, err := messages.ReturnMessageByLanguage(messages.MessageNewGroupNameAccepted, employee.Language)
			if err != nil {
				log.Println(err)
			}
			messageForEmployee = fmt.Sprintf(messageForEmployee, oldGroup.GroupName, newGroupName)
			msgForEmployee := tgbotapi.NewMessage(employee.ChatId, messageForEmployee)
			msgForEmployee.ReplyMarkup = keyboards.NewToMainMenuKeyboard(&employee)
			MsgChan <- msgForEmployee
		}
	}
	return nil
}

func (u *AuthService) ShowAllEmploysFromGroup(user *entities.User, callbackQueryData string) error {
	_, groupId, ok := strings.Cut(callbackQueryData, keys.ShowAllEmployeesFromGroupWithId)
	if !ok {
		log.Println("Помилка отримання ID групи")
	}
	groupIdInt, err := strconv.Atoi(groupId)
	if err != nil {
		log.Println("Помилка отримання ID групи")
	}
	user.Status = state_service.Adding_employee_to_group
	user.ActiveGroup = groupIdInt
	allEmployees, err := u.repository.ShowAllEmploysFromGroup(user.ActiveGroup)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}

	group, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		log.Println(err)
	}

	if allEmployees == nil {
		//In case if there are no employees in the group
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

		//In other case show all employees
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
	return err
}

func (u *AuthService) WarningBeforeDeletingEmployeeFromGroup(user *entities.User, callbackQueryData string) error {
	_, employeeUserIdString, ok := strings.Cut(callbackQueryData, keys.EmployeeIDtoDeleteFromGroup)
	if !ok {
		log.Println("Помилка отримання chatID працівника")
		return errors.New("error while getting ID of group member")
	}
	group, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		log.Println(err)
		return err
	}
	employeeUserIdInt, err := strconv.Atoi(employeeUserIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	employee, err := u.repository.GetUserByUserId(int64(employeeUserIdInt))
	if err != nil {
		log.Println("Помилка отримання даних працівника з БД")
		return err
	}

	messageToChief, err := messages.ReturnMessageByLanguage(messages.WarningMessageBeforeDeletingEmployeeFromGroup, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(messageToChief, employee.UserName, group.GroupName))
	msgToChief.ReplyMarkup = keyboards.NewConfirmDeletingEmployeeFromGroupKeyboard(user, employeeUserIdInt)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) DeleteEmployeeFromGroup(user *entities.User, receivedData string) error {
	_, employeeUserIdString, ok := strings.Cut(receivedData, keys.ConfirmDeletingEmployeeId)
	if !ok {
		log.Println("Помилка отримання chatID працівника")
		return errors.New("error while getting ID of group member")
	}
	group, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		log.Println(err)
		return err
	}
	employeeUserIdInt, err := strconv.Atoi(employeeUserIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	employee, err := u.repository.GetUserByUserId(int64(employeeUserIdInt))
	if err != nil {
		log.Println("Помилка отримання даних працівника з БД")
		return err
	}
	if err := u.repository.DeleteEmployeeFromGroup(&employee, &group); err != nil {
		log.Printf("Error while deleting user from group: %s", err)
		return err
	}
	messageToEmployee, err := messages.ReturnMessageByLanguage(messages.YouAreDeletedFromGroup, employee.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToEmployee := tgbotapi.NewMessage(employee.ChatId, fmt.Sprintf(messageToEmployee, group.GroupName))
	MsgChan <- msgToEmployee
	messageToChief, err := messages.ReturnMessageByLanguage(messages.EmployeeHaveBeenDeletedFromGroup, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(messageToChief, employee.UserName, group.GroupName))
	msgToChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) CopyEmployeeToAnotherGroup(user *entities.User, callbackQueryData string) error {
	_, employeeUserIdString, ok := strings.Cut(callbackQueryData, keys.EmployeeIDtoCopyToANotherGroup)
	if !ok {
		log.Println("Помилка отримання chatID працівника")
	}
	employeeUserIdInt, err := strconv.Atoi(employeeUserIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	groupsWithoutThisEmployee, err := u.repository.GetGroupsWithoutSelectedEmployee(employeeUserIdInt)
	if err != nil {
		return err
	}

	employee, err := u.repository.GetUserByUserId(int64(employeeUserIdInt))
	if err != nil {
		log.Println("Помилка отримання даних працівника з БД")
		return err
	}
	for _, v := range groupsWithoutThisEmployee {
		text, err := messages.ReturnMessageByLanguage(messages.MessageWhenCopyEmployeeToAnotherGroup, user.Language)
		if err != nil {
			log.Println("Помилка отримання тексту повідомлення")
		}
		msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(text, employee.UserName, v.GroupName))
		msg.ReplyMarkup = keyboards.NewCopyEmployeeKeyboard(user, int(v.Id), int(employee.UserId))
		MsgChan <- msg
	}
	text, err := messages.ReturnMessageByLanguage(messages.GoToMainMenu, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) ConfirmCopyEmployeeToAnotherGroup(user *entities.User, callbackQueryData string) error {
	groupString, employeeString, ok := strings.Cut(callbackQueryData, ":")
	if !ok {
		err := errors.New("error while cutting whole string")
		log.Println(err)
		return err
	}
	_, groupCode, ok := strings.Cut(groupString, keys.CopyEmployeeGroupID)
	if !ok {
		err := errors.New("error while cutting group info string")
		log.Println(err)
		return err
	}
	_, employeeCode, ok := strings.Cut(employeeString, keys.CopyEmployeeEmployeeID)
	if !ok {
		err := errors.New("error while cutting employee info string")
		log.Println(err)
		return err
	}
	groupIdInt, err := strconv.Atoi(groupCode)
	if err != nil {
		log.Println("Помилка визначення ID групи")
		return err
	}
	employeeIdInt, err := strconv.Atoi(employeeCode)
	if err != nil {
		log.Println("Помилка визначення ID учасника")
		return err
	}
	group, err := u.repository.GetGroupById(groupIdInt)
	if err != nil {
		return err
	}

	employee, err := u.repository.GetUserByUserId(int64(employeeIdInt))
	if err != nil {
		return err
	}

	if err := u.repository.AddEmployeeToGroup(groupIdInt, employeeIdInt); err != nil {
		log.Println(err)
		return err
	}
	//send message to chief
	messageToChief, err := messages.ReturnMessageByLanguage(messages.MessageToChiefUserCopied, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(messageToChief, employee.UserName, group.GroupName))
	msgToChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msgToChief

	//send message to employee
	messageToEmployee, err := messages.ReturnMessageByLanguage(messages.MessageToEmployeeCopied, employee.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToEmployee := tgbotapi.NewMessage(employee.ChatId, fmt.Sprintf(messageToEmployee, group.GroupName))
	msgToEmployee.ReplyMarkup = keyboards.NewToMainMenuKeyboard(&employee)
	MsgChan <- msgToEmployee
	return nil
}

func (u *AuthService) WarningBeforeGroupDeleting(user *entities.User, callbackQueryData string) error {
	_, groupIdString, ok := strings.Cut(callbackQueryData, keys.DeleteGroup)
	if !ok {
		log.Println("Помилка отримання group ID")
		return errors.New("error while getting group ID")
	}
	groupIdInt, err := strconv.Atoi(groupIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	group, err := u.repository.GetGroupById(groupIdInt)
	if err != nil {
		log.Println(err)
		return err
	}
	user.ActiveGroup = groupIdInt
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}
	messageToChief, err := messages.ReturnMessageByLanguage(messages.WarningBeforeGroupDeleting, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(messageToChief, group.GroupName))
	msgToChief.ReplyMarkup = keyboards.NewConfirmDeletingGroupKeyboard(user)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) DeleteGroup(user *entities.User, callbackQueryData string) error {
	//Get group information
	group, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		return err
	}

	// Get all employees from group
	allEmployees, err := u.repository.ShowAllEmploysFromGroup(user.ActiveGroup)
	if err != nil {
		return err
	}

	//Delete group from repository
	if err := u.repository.DeleteGroup(user.ActiveGroup); err != nil {
		log.Printf("Error while deleting user from group: %s", err)
		return err
	}

	//Send message to chief
	messageToChief, err := messages.ReturnMessageByLanguage(messages.GroupIsDeleted, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(messageToChief, group.GroupName))
	msgToChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msgToChief

	//Send message to employees
	if len(allEmployees) != 0 {
		for _, v := range allEmployees {
			messageToEmployee, err := messages.ReturnMessageByLanguage(messages.GroupIsDeleted, v.Language)
			if err != nil {
				log.Println("Помилка отримання тексту повідомлення")
			}
			msgToEmployee := tgbotapi.NewMessage(v.ChatId, fmt.Sprintf(messageToEmployee, group.GroupName))
			msgToEmployee.ReplyMarkup = keyboards.NewToMainMenuKeyboard(&v)
			MsgChan <- msgToEmployee
		}
	}
	return nil
}

func (u *AuthService) LeaveGroupBeforeConfirmation(user *entities.User, callbackQueryData string) error {
	_, groupIDString, ok := strings.Cut(callbackQueryData, keys.LeaveGroupID)
	if !ok {
		log.Println("Помилка отримання group ID")
		return errors.New("error while getting group ID")
	}
	groupIDInt, err := strconv.Atoi(groupIDString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	group, err := u.repository.GetGroupById(groupIDInt)
	if err != nil {
		log.Println(err)
		return err
	}
	user.ActiveGroup = groupIDInt
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}
	text, err := messages.ReturnMessageByLanguage(messages.LeaveGroupWarning, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	text = fmt.Sprintf(text, group.GroupName)
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewConfirmLeavingGroupKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) LeaveGroupWithConfirmation(user *entities.User) error {
	group, err := u.repository.GetGroupById(user.ActiveGroup)
	if err != nil {
		return err
	}

	chief, err := u.repository.GetUserByUserId(group.ChiefUserId)
	if err != nil {
		return err
	}

	if err := u.repository.LeaveGroup(int(user.UserId), user.ActiveGroup); err != nil {
		return err
	}

	textForEmployee, err := messages.ReturnMessageByLanguage(messages.MessageForEmployeeWhoLeavedGroup, user.Language)
	if err != nil {
		log.Println(err)
	}
	textForEmployee = fmt.Sprintf(textForEmployee, group.GroupName)
	msgForEmployee := tgbotapi.NewMessage(user.ChatId, textForEmployee)
	msgForEmployee.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msgForEmployee

	textForChief, err := messages.ReturnMessageByLanguage(messages.MessageForChiefAboutEmployeeLeftGroup, chief.Language)
	if err != nil {
		log.Println(err)
	}
	textForChief = fmt.Sprintf(textForChief, user.UserName, group.GroupName)
	msgForChief := tgbotapi.NewMessage(chief.ChatId, textForChief)
	msgForChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(&chief)
	MsgChan <- msgForChief
	return nil
}

func (u *AuthService) CreateNewTaskAskingTitle(user *entities.User, callbackQueryData string) error {
	_, groupIdString, ok := strings.Cut(callbackQueryData, keys.CreateNewTaskKeyData)
	if !ok {
		log.Println("Помилка отримання group ID")
		return errors.New("error while getting group ID")
	}
	groupIdInt, err := strconv.Atoi(groupIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	user.ActiveGroup = groupIdInt
	user.Status = state_service.Expecting_new_task_title
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}
	messageToChief, err := messages.ReturnMessageByLanguage(messages.ExpectingNewTaskTitleMessage, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, messageToChief)
	msgToChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) CreateNewTaskGotTitle(user *entities.User, taskTitle string) error {
	taskID, err := u.repository.CreateNewTask(taskTitle, user.ActiveGroup)
	if err != nil {
		return err
	}

	user.ActiveTask = taskID
	user.ActiveGroup = 0
	user.Status = state_service.Expecting_new_task_description
	if err := u.repository.UpdateStatus(user); err != nil {
		log.Println(err)
		return err
	}

	messageToChief, err := messages.ReturnMessageByLanguage(messages.NewTaskTitleAcceptedExpectingNewTaskDesc, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf(messageToChief, taskTitle))
	msgToChief.ReplyMarkup = keyboards.NewExpectingDescriptionOrSkipKeyboard(user)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) UpdateTaskDescription(user *entities.User, taskDesc string) error {
	if err := u.repository.UpdateTaskDescription(taskDesc, user.ActiveTask); err != nil {
		return err
	}
	text, err := messages.ReturnMessageByLanguage(messages.TaskDescriptionAccepted, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	text = fmt.Sprintf(text, taskDesc)
	msgToChief := tgbotapi.NewMessage(user.ChatId, text)
	msgToChief.ReplyMarkup = keyboards.NewAssignToEntireGroupAllSomeEmployees(user)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) AddAllEmployeesToTask(user *entities.User) error {
	task, err := u.repository.GetTaskByID(user.ActiveTask)
	if err != nil {
		return err
	}
	group, err := u.repository.GetGroupById(int(task.GroupId))
	if err != nil {
		return err
	}
	allEmployees, err := u.repository.ShowAllEmploysFromGroup(int(task.GroupId))
	if err != nil {
		return err
	}
	if len(allEmployees) == 0 {
		// Message to chief in case when there aren`t employees in the group
		text, err := messages.ReturnMessageByLanguage(messages.NoEmployeesInTheGroup, user.Language)
		if err != nil {
			log.Println(err)
		}
		text = fmt.Sprintf(text, group.GroupName)
		msgToChief := tgbotapi.NewMessage(user.ChatId, text)
		msgToChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		MsgChan <- msgToChief
		return nil
	}
	for _, employee := range allEmployees {
		if err := u.repository.AddEmployeeToTask(int(task.Id), int(employee.UserId)); err != nil {
			log.Println(err)
		}
		//Message to employee
		textToEmployee, err := messages.ReturnMessageByLanguage(messages.YouHaveBeenAssignedTheTask, employee.Language)
		if err != nil {
			log.Println(err)
		}
		textToEmployee = fmt.Sprintf(textToEmployee, task.TaskName)
		msgToEmployee := tgbotapi.NewMessage(employee.ChatId, textToEmployee)
		msgToEmployee.ReplyMarkup = keyboards.SeeTaskDetailsForEmployee(&employee, int(task.Id))
		MsgChan <- msgToEmployee
	}
	text, err := messages.ReturnMessageByLanguage(messages.AllEmployeesFromGroupSuccessfullyAddedToTask, user.Language)
	if err != nil {
		log.Println(err)
	}
	text = fmt.Sprintf(text, task.TaskName, group.GroupName)
	msgToChief := tgbotapi.NewMessage(user.ChatId, text)
	msgToChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) AddSomeEmployeesToTaskShowList(user *entities.User, callbackQueryData string) error {
	var taskID int
	if user.ActiveTask != 0 {
		taskID = user.ActiveTask
	} else if callbackQueryData != "" {
		_, taskIdString, ok := strings.Cut(callbackQueryData, keys.AddExecutorForChiefKeyData)
		if !ok {
			log.Println("Помилка отримання group ID")
			return errors.New("error while getting group ID")
		}
		taskIdInt, err := strconv.Atoi(taskIdString)
		if err != nil {
			log.Println("Помилка визначення ID")
			return err
		}
		taskID = taskIdInt
	} else {
		return fmt.Errorf("cannot parse task ID")
	}

	allEmployees, err := u.repository.GetEmployeesWhichAreInTheGroupButNotAssignedToTheTask(taskID)
	if err != nil {
		return err
	}
	if len(allEmployees) == 0 {
		//In case if there are no employees
		text, err := messages.ReturnMessageByLanguage(messages.NoEmployeesToAssignToTheTask, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, text)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		MsgChan <- msg
	} else {
		for _, e := range allEmployees {
			msg := tgbotapi.NewMessage(user.ChatId, e.UserName)
			msg.ReplyMarkup = keyboards.NewAssignKeyboard(user, int(e.UserId), taskID)
			MsgChan <- msg
		}
	}
	return nil
}

func (u *AuthService) AssignEmployee(user *entities.User, callbackQueryData string) error {
	taskIDString, employeeIDString, ok := strings.Cut(callbackQueryData, keys.AssignToEmployeeWithID)
	if !ok {
		return errors.New("error while getting group ID")
	}
	taskIdInt, err := strconv.Atoi(taskIDString)
	if err != nil {
		return err
	}

	task, err := u.repository.GetTaskByID(taskIdInt)
	if err != nil {
		return err
	}

	employeeIDInt, err := strconv.Atoi(employeeIDString)
	if err != nil {
		return err
	}

	employee, err := u.repository.GetUserByUserId(int64(employeeIDInt))
	if err != nil {
		return err
	}

	err = u.repository.AddEmployeeToTask(taskIdInt, employeeIDInt)
	if err != nil {
		return err
	}
	//make message to chief
	text, err := messages.ReturnMessageByLanguage(messages.EmployeeAddedToTheTaskMessage, user.Language)
	if err != nil {
		log.Println(err)
	}
	text = fmt.Sprintf(text, task.TaskName, employee.UserName)
	msgToChief := tgbotapi.NewMessage(user.ChatId, text)
	msgToChief.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msgToChief

	//make message to employee
	textToEmployee, err := messages.ReturnMessageByLanguage(messages.YouHaveBeenAssignedTheTask, employee.Language)
	if err != nil {
		log.Println(err)
	}
	textToEmployee = fmt.Sprintf(textToEmployee, task.TaskName)
	msgToEmployee := tgbotapi.NewMessage(employee.ChatId, textToEmployee)
	msgToEmployee.ReplyMarkup = keyboards.SeeTaskDetailsForEmployee(&employee, int(task.Id))
	MsgChan <- msgToEmployee
	return nil
}

func (u *AuthService) SkipDescritionEntering(user *entities.User) error {
	text, err := messages.ReturnMessageByLanguage(messages.TaskDescriptionSkipped, user.Language)
	if err != nil {
		log.Println("Помилка отримання тексту повідомлення")
	}
	msgToChief := tgbotapi.NewMessage(user.ChatId, text)
	msgToChief.ReplyMarkup = keyboards.NewAssignToEntireGroupAllSomeEmployees(user)
	MsgChan <- msgToChief
	return nil
}

func (u *AuthService) ShowTaskDetailsForEmployee(user *entities.User, callbackQueryData string) error {
	//Gettin Task ID from callbasc data
	_, taskIDString, ok := strings.Cut(callbackQueryData, keys.ToSeeTaskDetailsTaskIDForEmployee)
	if !ok {
		log.Println("Помилка отримання group ID")
		return errors.New("error while getting group ID")
	}
	taskIdInt, err := strconv.Atoi(taskIDString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}

	//Getting task entitie from repository
	task, err := u.repository.GetTaskByID(taskIdInt)
	if err != nil {
		return err
	}
	taskDesc, err := messages.ReturnMessageByLanguage(messages.NoTaskDescription, user.Language)
	if err != nil {
		return err
	}
	if task.TaskDescription != "no_description" {
		taskDesc = task.TaskDescription
	}

	//Geting group entitie from repository
	group, err := u.repository.GetGroupById(int(task.GroupId))
	if err != nil {
		return err
	}

	//Geting slice of executors from repository and formating string of their names
	allExecutors, err := u.repository.GetAllExecutors(int(task.Id))
	if err != nil {
		return err
	}
	executorsString := ""
	for index, executor := range allExecutors {
		executorsString += executor.UserName
		if index != len(allExecutors)-1 {
			executorsString += "\n"
		}
	}

	//Formating creating date
	createdTime, err := time.Parse(time.RFC3339, task.CreatingTime)
	if err != nil {
		log.Println(err)
	}
	createdString := createdTime.Format("02.01.2006")
	if err != nil {
		createdString = "parsing time error"
	}

	text, err := messages.ReturnMessageByLanguage(messages.TaskDetailsForEmployees, user.Language)
	if err != nil {
		log.Println(err)
	}
	text = fmt.Sprintf(text, task.TaskName, group.GroupName, createdString, taskDesc, executorsString)
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
	MsgChan <- msg
	return nil
}

func (u *AuthService) ShowTaskDetailsForChief(user *entities.User, callbackQueryData string) error {
	//Gettin Task ID from callback data
	_, taskIDString, ok := strings.Cut(callbackQueryData, keys.ToSeeTaskDetailsTaskIDForChief)
	if !ok {
		log.Println("Помилка отримання group ID")
		return errors.New("error while getting group ID")
	}
	taskIdInt, err := strconv.Atoi(taskIDString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}

	//Getting task entitie from repository
	task, err := u.repository.GetTaskByID(taskIdInt)
	if err != nil {
		return err
	}
	taskDesc, err := messages.ReturnMessageByLanguage(messages.NoTaskDescription, user.Language)
	if err != nil {
		return err
	}
	if task.TaskDescription != "no_description" {
		taskDesc = task.TaskDescription
	}

	//Geting group entitie from repository
	group, err := u.repository.GetGroupById(int(task.GroupId))
	if err != nil {
		return err
	}

	//Geting slice of executors from repository and formating string of their names
	allExecutors, err := u.repository.GetAllExecutors(int(task.Id))
	if err != nil {
		return err
	}
	executorsString := ""
	for index, executor := range allExecutors {
		executorsString += executor.UserName
		if index != len(allExecutors)-1 {
			executorsString += "\n"
		}
	}

	//Formating creating date
	createdTime, err := time.Parse(time.RFC3339, task.CreatingTime)
	if err != nil {
		log.Println(err)
	}
	createdString := createdTime.Format("02.01.2006")
	if err != nil {
		createdString = "parsing time error"
	}

	text, err := messages.ReturnMessageByLanguage(messages.TaskDetailsForEmployees, user.Language)
	if err != nil {
		log.Println(err)
	}
	text = fmt.Sprintf(text, task.TaskName, group.GroupName, createdString, taskDesc, executorsString)
	msg := tgbotapi.NewMessage(user.ChatId, text)
	msg.ReplyMarkup = keyboards.TaskDetailsKeyboardForChief(user, taskIdInt)
	MsgChan <- msg
	return nil
}

// ShowAllGroupTasksForChief shows all tasks from selected group for Chief
func (u *AuthService) ShowAllGroupTasksForChief(user *entities.User, callbackQueryData string) error {
	_, groupIdString, ok := strings.Cut(callbackQueryData, keys.ShowAllTasksByGorupIDForChief)
	if !ok {
		log.Println("Помилка отримання group ID")
		return errors.New("error while getting group ID")
	}
	groupIdInt, err := strconv.Atoi(groupIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	tasks, err := u.repository.GetAllTasksByGroupID(int64(groupIdInt))
	if err != nil {
		return err
	}

	if len(tasks) != 0 {
		for _, task := range tasks {
			text := fmt.Sprintf(task.TaskName)
			msg := tgbotapi.NewMessage(user.ChatId, text)
			msg.ReplyMarkup = keyboards.SeeTaskDetailsForChief(user, int(task.Id))
			MsgChan <- msg
		}
	} else {
		text, err := messages.ReturnMessageByLanguage(messages.NoTasksForGroup, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, text)
		//TODO:This keyboard is only for Chief, but shows to employyes too!
		msg.ReplyMarkup = keyboards.CreateNewTaskKeyboard(user, groupIdInt)
		MsgChan <- msg
	}

	return nil
}

// ShowAllGroupTasksForEmployee shows all tasks from selected group for Employees
func (u *AuthService) ShowAllGroupTasksForEmployee(user *entities.User, callbackQueryData string) error {
	_, groupIdString, ok := strings.Cut(callbackQueryData, keys.ShowAllTasksByGorupIDForEmployee)
	if !ok {
		log.Println("Помилка отримання group ID")
		return errors.New("error while getting group ID")
	}
	groupIdInt, err := strconv.Atoi(groupIdString)
	if err != nil {
		log.Println("Помилка визначення ID")
		return err
	}
	tasks, err := u.repository.GetAllTasksByGroupID(int64(groupIdInt))
	if err != nil {
		return err
	}

	if len(tasks) != 0 {
		for _, task := range tasks {
			text := fmt.Sprintf(task.TaskName)
			msg := tgbotapi.NewMessage(user.ChatId, text)
			msg.ReplyMarkup = keyboards.SeeTaskDetailsForEmployee(user, int(task.Id))
			MsgChan <- msg
		}
	} else {
		text, err := messages.ReturnMessageByLanguage(messages.NoTasksForGroup, user.Language)
		if err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(user.ChatId, text)
		msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(user)
		MsgChan <- msg
	}
	return nil
}

func (u *AuthService) MarkTaskAsComplete(user *entities.User, callbackQueryData string) error {
	// TODO: Implement
	return nil
}

func (u *AuthService) DeleteTask(user *entities.User, callbackQueryData string) error {
	// TODO: Implement
	return nil
}

func (u *AuthService) DeleteExecutors(user *entities.User, callbackQueryData string) error {
	// TODO: Implement
	return nil
}
