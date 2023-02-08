package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"log"
	"strconv"
)

var (
	// LanguageKeyboard hold keys for language choosing
	LanguageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Українська", messages.LanguageKey+string(messages.UA)),
			tgbotapi.NewInlineKeyboardButtonData("English", messages.LanguageKey+string(messages.EN)),
		),
	)
	// RemoveKeyboard removes any keyboards
	RemoveKeyboard = tgbotapi.NewRemoveKeyboard(true)
)

// NewToMainMenuKeyboard returns new inline keyboard with only one button to MainMenu
func NewToMainMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	keyText, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	toMainMenuKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyText, keys.GoToMainMenu),
		),
	)
	return toMainMenuKeyboard
}

// NewKeyboardChooseCreateOrJoinGroup returns new inline keyboard with two keys: create new
// group or join to exist group
func NewKeyboardChooseCreateOrJoinGroup(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	keyCreate, err := messages.ReturnMessageByLanguage(messages.CreateNewGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyJoin, err := messages.ReturnMessageByLanguage(messages.JoinToGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyCreate, keys.CreateNewGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyJoin, keys.JoinTheExistGroup),
		),
	)
	return keyboard
}

// NewMainMenuKeyboard returns new inline keyboard for Main Menu
func NewMainMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	userMenuTitle, err := messages.ReturnMessageByLanguage(messages.UserSettingsMenuTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	groupMenuTitle, err := messages.ReturnMessageByLanguage(messages.GroupMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(userMenuTitle, keys.GoToUserMenuSettings),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(groupMenuTitle, keys.GoToGroupsMenu),
		),
	)
	return keyboard

}

// NewUserSettingsMenuKeyboard returns new inline keyboard for User Settings Menu
func NewUserSettingsMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	changeLanguageKeyTitle, err := messages.ReturnMessageByLanguage(messages.ChangeLanguageKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	changeUserNameKeyTitle, err := messages.ReturnMessageByLanguage(messages.ChangeUserName, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyJoin, err := messages.ReturnMessageByLanguage(messages.JoinToGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyText, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(changeLanguageKeyTitle, keys.GoToChangeLanguageMenu),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(changeUserNameKeyTitle, keys.GoToChangeUserName),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyJoin, keys.JoinTheExistGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyText, keys.GoToMainMenu),
		),
	)
	return keyboard
}

// NewGroupMenuKeyboard returns new inline keyboard for Groups Menu
func NewGroupMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	createNewGroupKey, err := messages.ReturnMessageByLanguage(messages.CreateNewGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	showAllChiefsGroups, err := messages.ReturnMessageByLanguage(messages.ShowAllChiefsGroups, user.Language)
	if err != nil {
		log.Println(err)
	}
	showAllEmployeeGroups, err := messages.ReturnMessageByLanguage(messages.ShowAllEmployeeGroups, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyJoin, err := messages.ReturnMessageByLanguage(messages.JoinToGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(createNewGroupKey, keys.CreateNewGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(showAllChiefsGroups, keys.ShowAllChiefsGroups),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(showAllEmployeeGroups, keys.ShowAllEmployeeGroups),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyJoin, keys.JoinTheExistGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)

	return keyboard
}

// NewMenuForEvenGroup returns new inline keyboard for even group after showing all groups in Group Menu
func NewMenuForEvenGroup(user *entities.User, group *entities.Group) tgbotapi.InlineKeyboardMarkup {
	changeNameTitle, err := messages.ReturnMessageByLanguage(messages.RenameGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	changeNameKey := keys.RenameGroupWithId + strconv.Itoa(int(group.Id))

	showEmployeesTitle, err := messages.ReturnMessageByLanguage(messages.ShowAllEmployeesFromGroupWithIdKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	showEmployeesKey := keys.ShowAllEmployeesFromGroupWithId + strconv.Itoa(int(group.Id))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(changeNameTitle, changeNameKey),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(showEmployeesTitle, showEmployeesKey),
		),
	)
	return keyboard
}

func NewMenuForEvenEmployee(user, employee *entities.User) tgbotapi.InlineKeyboardMarkup {
	deleteTitle, err := messages.ReturnMessageByLanguage(messages.DeleteEmployeeFromGroupKeyText, user.Language)
	if err != nil {
		log.Println(err)
	}
	copyTitle, err := messages.ReturnMessageByLanguage(messages.CopyEmployeeToAnotherGroupKeyText, user.Language)
	if err != nil {
		log.Println(err)
	}
	moveTitle, err := messages.ReturnMessageByLanguage(messages.MoveEmployeeToAnotherGroupKeyText, user.Language)
	if err != nil {
		log.Println(err)
	}
	deleteKey := keys.EmployeeIDtoDeleteFromGroup + strconv.Itoa(int(employee.UserId))
	copyKey := keys.EmployeeIDtoCopyToANotherGroup + strconv.Itoa(int(employee.UserId))
	moveKey := keys.EmployeeIDtoMoveToANotherGroup + strconv.Itoa(int(employee.UserId))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(deleteTitle, deleteKey),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(copyTitle, copyKey),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(moveTitle, moveKey),
		),
	)
	return keyboard
}

func NewGroupCreatingKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	createNewGroupKey, err := messages.ReturnMessageByLanguage(messages.CreateNewGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyJoin, err := messages.ReturnMessageByLanguage(messages.JoinToGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(createNewGroupKey, keys.CreateNewGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyJoin, keys.JoinTheExistGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)

	return keyboard
}
