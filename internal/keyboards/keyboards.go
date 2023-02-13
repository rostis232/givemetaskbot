package keyboards

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
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
	//TODO: Add key to adding new task from Main Menu
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
	//TODO: Add key to create new task from group menu
	createNewTaskKeyTitle, err := messages.ReturnMessageByLanguage(messages.CreateNewTaskKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	createNewTaskKeyData := keys.CreateNewTaskKeyData + strconv.Itoa(int(group.Id))

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

	deleteGroupTitle, err := messages.ReturnMessageByLanguage(messages.DeleteGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	deleteGroupKey := keys.DeleteGroup + strconv.Itoa(int(group.Id))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(createNewTaskKeyTitle, createNewTaskKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(changeNameTitle, changeNameKey),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(showEmployeesTitle, showEmployeesKey),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(deleteGroupTitle, deleteGroupKey),
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
	deleteKey := keys.EmployeeIDtoDeleteFromGroup + strconv.Itoa(int(employee.UserId))
	copyKey := keys.EmployeeIDtoCopyToANotherGroup + strconv.Itoa(int(employee.UserId))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(deleteTitle, deleteKey),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(copyTitle, copyKey),
		),
	)
	return keyboard
}

func NewCopyEmployeeKeyboard(user *entities.User, groupID int, employeeID int) tgbotapi.InlineKeyboardMarkup {
	copyKey, err := messages.ReturnMessageByLanguage(messages.CopeEmployeeKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	copyData := keys.CopyEmployeeGroupID + strconv.Itoa(groupID) + ":" + keys.CopyEmployeeEmployeeID + strconv.Itoa(employeeID)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(copyKey, copyData),
		),
	)
	return keyboard
}

func NewConfirmDeletingEmployeeFromGroupKeyboard(user *entities.User, employeeID int) tgbotapi.InlineKeyboardMarkup {
	confirmKey, err := messages.ReturnMessageByLanguage(messages.ConfirmationButtonTextForDeletingEmployeeFromGroup, user.Language)
	if err != nil {
		log.Println(err)
	}
	confirmData := keys.ConfirmDeletingEmployeeId + strconv.Itoa(employeeID)
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(confirmKey, confirmData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)
	return keyboard
}

func NewConfirmDeletingGroupKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	confirmKey, err := messages.ReturnMessageByLanguage(messages.ConfirmationButtonTextForDeletingGroup, user.Language)
	if err != nil {
		log.Println(err)
	}
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(confirmKey, keys.ConfirmDeletingGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)
	return keyboard
}

func NewMenuForEvenGroupForEmployee(lng messages.Language, groupID int64) tgbotapi.InlineKeyboardMarkup {
	leaveKey, err := messages.ReturnMessageByLanguage(messages.ExitFromTheGroup, lng)
	if err != nil {
		log.Println(err)
	}
	leaveData := keys.LeaveGroupID + strconv.Itoa(int(groupID))
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(leaveKey, leaveData),
		),
	)

	return keyboard
}

func NewConfirmLeavingGroupKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	confirmKey, err := messages.ReturnMessageByLanguage(messages.ConfirmLeavingGroup, user.Language)
	if err != nil {
		log.Println(err)
	}
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(confirmKey, keys.ConfirmLeavingGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)
	return keyboard
}
