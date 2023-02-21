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
func NewMenuForEvenGroup(user *entities.User, groupID int) tgbotapi.InlineKeyboardMarkup {
	createNewTaskKeyTitle, err := messages.ReturnMessageByLanguage(messages.CreateNewTaskKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	createNewTaskKeyData := keys.CreateNewTaskKeyData + strconv.Itoa(groupID)

	showTasksKeyTitle, err := messages.ReturnMessageByLanguage(messages.ShowAllTasksFromGroupKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	showTasksKeyData := keys.ShowAllTasksByGorupIDForChief + strconv.Itoa(groupID)

	changeNameTitle, err := messages.ReturnMessageByLanguage(messages.RenameGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	changeNameKey := keys.RenameGroupWithId + strconv.Itoa(groupID)

	showEmployeesTitle, err := messages.ReturnMessageByLanguage(messages.ShowAllEmployeesFromGroupWithIdKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	showEmployeesKey := keys.ShowAllEmployeesFromGroupWithId + strconv.Itoa(groupID)

	deleteGroupTitle, err := messages.ReturnMessageByLanguage(messages.DeleteGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	deleteGroupKey := keys.DeleteGroup + strconv.Itoa(groupID)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(createNewTaskKeyTitle, createNewTaskKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(showTasksKeyTitle, showTasksKeyData),
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
		leaveKey = "Error key title parsing"
	}
	leaveData := keys.LeaveGroupID + strconv.Itoa(int(groupID))

	showTasksKeyTitle, err := messages.ReturnMessageByLanguage(messages.ShowAllTasksFromGroupKeyTitle, lng)
	if err != nil {
		log.Println(err)
		showTasksKeyTitle = "Error key title parsing"
	}
	showTaskData := keys.ShowAllTasksByGorupIDForEmployee + strconv.Itoa(int(groupID))
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(showTasksKeyTitle, showTaskData),
		),
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

func NewExpectingDescriptionOrSkipKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	skipKeyTitle, err := messages.ReturnMessageByLanguage(messages.SkipKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(skipKeyTitle, keys.SkipDescritionEntering),
		),
	)
	return keyboard
}

func NewAssignToEntireGroupAllSomeEmployees(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	toWholeGroupKeyTitle, err := messages.ReturnMessageByLanguage(messages.AddWholeGroupToTaskKeyTitle, user.Language)
	if err != nil {
		log.Println()
	}
	toSomeEmployees, err := messages.ReturnMessageByLanguage(messages.AddSomeEmployeesToTaskKeyTitle, user.Language)
	if err != nil {
		log.Println()
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(toWholeGroupKeyTitle, keys.AddWholeGroupToTaskKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(toSomeEmployees, keys.AddSomeEmployeesToTaskKeyData),
		),
	)
	return keyboard
}

func NewAssignKeyboard(user *entities.User, employeeID int) tgbotapi.InlineKeyboardMarkup {
	assignKeyTitle, err := messages.ReturnMessageByLanguage(messages.ToAssignKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	assignKeyData := keys.AssignToEmployeeWithID + strconv.Itoa(employeeID)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(assignKeyTitle, assignKeyData),
		), //TODO: Implement functionality
	)
	return keyboard
}

func SeeTaskDetailsForEmployee(user *entities.User, taskID int) tgbotapi.InlineKeyboardMarkup {
	toViewTaskDetailsKeyTitle, err := messages.ReturnMessageByLanguage(messages.ToSeeTaskDetailsKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	toViewTaskDetailsKeyData := keys.ToSeeTaskDetailsTaskIDForEmployee + strconv.Itoa(taskID)
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(toViewTaskDetailsKeyTitle, toViewTaskDetailsKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)
	return keyboard
}

func SeeTaskDetailsForChief(user *entities.User, taskID int) tgbotapi.InlineKeyboardMarkup {
	toViewTaskDetailsKeyTitle, err := messages.ReturnMessageByLanguage(messages.ToSeeTaskDetailsKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	toViewTaskDetailsKeyData := keys.ToSeeTaskDetailsTaskIDForChief + strconv.Itoa(taskID)
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(toViewTaskDetailsKeyTitle, toViewTaskDetailsKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)
	return keyboard
}

//It`ll be needed when implementing time control
// func NewTimeZonesKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
// 	keyboard := tgbotapi.NewInlineKeyboardMarkup(
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("-12", keys.TimeZone+"-1200"),
// 			tgbotapi.NewInlineKeyboardButtonData("-11", keys.TimeZone+"-1100"),
// 			tgbotapi.NewInlineKeyboardButtonData("-10", keys.TimeZone+"-1000"),
// 			tgbotapi.NewInlineKeyboardButtonData("-9:30", keys.TimeZone+"-0930"),
// 			tgbotapi.NewInlineKeyboardButtonData("-9", keys.TimeZone+"-0900"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("-8", keys.TimeZone+"-0800"),
// 			tgbotapi.NewInlineKeyboardButtonData("-7", keys.TimeZone+"-0700"),
// 			tgbotapi.NewInlineKeyboardButtonData("-6", keys.TimeZone+"-0600"),
// 			tgbotapi.NewInlineKeyboardButtonData("-5", keys.TimeZone+"-0500"),
// 			tgbotapi.NewInlineKeyboardButtonData("-4:30", keys.TimeZone+"-0430"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("-4", keys.TimeZone+"-0400"),
// 			tgbotapi.NewInlineKeyboardButtonData("-3:30", keys.TimeZone+"-0330"),
// 			tgbotapi.NewInlineKeyboardButtonData("-3", keys.TimeZone+"-0300"),
// 			tgbotapi.NewInlineKeyboardButtonData("-2", keys.TimeZone+"-0200"),
// 			tgbotapi.NewInlineKeyboardButtonData("-1", keys.TimeZone+"-0100"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("0", keys.TimeZone+"+0000"),
// 			tgbotapi.NewInlineKeyboardButtonData("+1", keys.TimeZone+"+0100"),
// 			tgbotapi.NewInlineKeyboardButtonData("+2", keys.TimeZone+"+0200"),
// 			tgbotapi.NewInlineKeyboardButtonData("+3", keys.TimeZone+"+0300"),
// 			tgbotapi.NewInlineKeyboardButtonData("+3:30", keys.TimeZone+"+0330"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("+4", keys.TimeZone+"+0400"),
// 			tgbotapi.NewInlineKeyboardButtonData("+4:30", keys.TimeZone+"+0430"),
// 			tgbotapi.NewInlineKeyboardButtonData("+5", keys.TimeZone+"+0500"),
// 			tgbotapi.NewInlineKeyboardButtonData("+5:30", keys.TimeZone+"+0530"),
// 			tgbotapi.NewInlineKeyboardButtonData("+5:45", keys.TimeZone+"+0545"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("+6", keys.TimeZone+"+0600"),
// 			tgbotapi.NewInlineKeyboardButtonData("+6:30", keys.TimeZone+"+0630"),
// 			tgbotapi.NewInlineKeyboardButtonData("+7", keys.TimeZone+"+0700"),
// 			tgbotapi.NewInlineKeyboardButtonData("+8", keys.TimeZone+"+0800"),
// 			tgbotapi.NewInlineKeyboardButtonData("+9", keys.TimeZone+"+0900"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("+9:30", keys.TimeZone+"+0930"),
// 			tgbotapi.NewInlineKeyboardButtonData("+10", keys.TimeZone+"+1000"),
// 			tgbotapi.NewInlineKeyboardButtonData("+10:30", keys.TimeZone+"+1030"),
// 			tgbotapi.NewInlineKeyboardButtonData("+11", keys.TimeZone+"+1100"),
// 			tgbotapi.NewInlineKeyboardButtonData("+11:30", keys.TimeZone+"+1130"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("+12", keys.TimeZone+"+1200"),
// 			tgbotapi.NewInlineKeyboardButtonData("+12:45", keys.TimeZone+"+1245"),
// 			tgbotapi.NewInlineKeyboardButtonData("+13", keys.TimeZone+"+1300"),
// 			tgbotapi.NewInlineKeyboardButtonData("+13:45", keys.TimeZone+"+1345"),
// 			tgbotapi.NewInlineKeyboardButtonData("+14", keys.TimeZone+"+1400"),
// 		),
// 	)
// 	return keyboard
// }

func CreateNewTaskKeyboard(user *entities.User, groupID int) tgbotapi.InlineKeyboardMarkup {
	createNewTaskKeyTitle, err := messages.ReturnMessageByLanguage(messages.CreateNewTaskKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	createNewTaskKeyData := keys.CreateNewTaskKeyData + strconv.Itoa(groupID)
	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(createNewTaskKeyTitle, createNewTaskKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)
	return keyboard
}

func TaskDetailsKeyboardForChief(user *entities.User, taskID int) tgbotapi.InlineKeyboardMarkup {
	markAsDoneKeyTitle, err := messages.ReturnMessageByLanguage(messages.MarkTaskAsClosedForChiefKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	markAsDoneKeyData := keys.MarkTaskAsClosedForChiefKeyData + strconv.Itoa(taskID)

	deleteKeyTitle, err := messages.ReturnMessageByLanguage(messages.DeleteTaskForChiefKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	deleteKeyData := keys.DeleteTaskForChiefKeyData + strconv.Itoa(taskID)

	addExecutorKeyTitle, err := messages.ReturnMessageByLanguage(messages.AddExecutorForChiefKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	addExecutorKeyData := keys.AddExecutorForChiefKeyData + strconv.Itoa(taskID)

	removeExecutorKeyTitle, err := messages.ReturnMessageByLanguage(messages.RemoveExecutorForChiefKeyTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	removeExecutorKeyData := keys.RemoveExecutorForChiefKeyData + strconv.Itoa(taskID)

	mainMenuKeyTitle, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(markAsDoneKeyTitle, markAsDoneKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(deleteKeyTitle, deleteKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(addExecutorKeyTitle, addExecutorKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(removeExecutorKeyTitle, removeExecutorKeyData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKeyTitle, keys.GoToMainMenu),
		),
	)
	return keyboard
}
