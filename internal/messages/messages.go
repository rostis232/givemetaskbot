package messages

import (
	"fmt"
)

const (
	UnknownError                                       = "Unknown error: %s. Contact with administrator @rostis232"
	UnknownCommand                                     = "Unknown command"
	MessageForUnregisteredUsers                        = "MessageForUnregisteredUsers"
	MessageIfUserAlreadyExists                         = "MessageIfUserAlreadyExists"
	MessageAfterFirstLanguageSelection                 = "MessageAfterFirstLanguageSelection"
	MessageAfterLanguageUpdate                         = "MessageAfterLanguageUpdate"
	MessageAfterFirstNameEntering                      = "MessageAfterFirstNameEntering"
	MessageAfterUserNameUpdate                         = "MessageAfterUserNameUpdate"
	ToMainMenuKey                                      = "ToMainMenuKey"
	CreateNewGroupKey                                  = "CreateNewGroupKey"
	JoinToGroupKey                                     = "JoinToGroupKey"
	MessageWithChatId                                  = "MessageWithChatId"
	MainMenuTitle                                      = "MainMenuTitle"
	UserSettingsMenuTitle                              = "UserSettingsMenuTitle"
	ChangeLanguageKey                                  = "ChangeLanguageKey"
	ChangeUserName                                     = "ChangeUserName"
	MessageBeforeNameChanging                          = "MessageBeforeNameChanging"
	MessageEnterNewGroupName                           = "MessageEnterNewGroupName"
	MessageCreatedNewGroup                             = "MessageCreatedNewGroup"
	MessageEmployeeCodeBroken                          = "MessageEmployeeCodeBroken"
	MessageNoEmployeeWithThisCode                      = "MessageNoEmployeeWithThisCode"
	MessageEmployeeCodeEqualsChiefCode                 = "MessageEmployeeCodeEqualsChiefCode"
	MessageEmployeeAddingSuccess                       = "MessageEmployeeAddingSuccess"
	GroupMenuKey                                       = "GroupMenuKey"
	GroupMenuTitle                                     = "GroupMenuTitle"
	ShowAllChiefsGroups                                = "ShowAllChiefsGroups"
	RenameGroupKey                                     = "RenameGroupKey"
	RenameGroupTitle                                   = "RenameGroupTitle"
	MessageNewGroupNameAccepted                        = "MessageNewGroupNameAccepted"
	ShowAllEmployeesFromGroupWithIdKey                 = "ShowAllEmployeesFromGroupWithIdKey"
	AddNewEmployeeToExistingGroup                      = "AddNewEmployeeToExistingGroup"
	NoEmployeesInTheGroup                              = "NoEmployeesInTheGroup"
	DeleteEmployeeFromGroupKeyText                     = "DeleteEmployeeFromGroupKeyText"
	CopyEmployeeToAnotherGroupKeyText                  = "CopyEmployeeToAnotherGroupKeyText"
	MoveEmployeeToAnotherGroupKeyText                  = "MoveEmployeeToAnotherGroupKeyText"
	NoGroups                                           = "NoGroups"
	ShownMembersOfTheGroupWithName                     = "ShownMembersOfTheGroupWithName"
	YouAreDeletedFromGroup                             = "YouAreDeletedFromGroup"
	EmployeeHaveBeenDeletedFromGroup                   = "EmployeeHaveBeenDeletedFromGroup"
	AddNewGroupFromGroupListMenu                       = "AddNewGroupFromGroupListMenu"
	ShowAllEmployeeGroups                              = "ShowAllEmployeeGroups"
	MessageWhenCopyEmployeeToAnotherGroup              = "MessageWhenCopyEmployeeToAnotherGroup"
	CopeEmployeeKey                                    = "CopeEmployeeKey"
	DeleteGroupKey                                     = "DeleteGroupKey"
	WarningMessageBeforeDeletingEmployeeFromGroup      = "WarningMessageBeforeDeletingEmployeeFromGroup"
	ConfirmationButtonTextForDeletingEmployeeFromGroup = "ConfirmationButtonTextForDeletingEmployeeFromGroup"
	ConfirmationButtonTextForDeletingGroup             = "ConfirmationButtonTextForDeletingGroup"
	WarningBeforeGroupDeleting                         = "WarningBeforeGroupDeleting"
	GroupIsDeleted                                     = "GroupIsDeleted"
	MessageToChiefUserCopied                           = "MessageToChiefUserCopied"
	MessageToEmployeeCopied                            = "MessageToEmployeeCopied"
	GoToMainMenu                                       = "GoToMainMenu"
	ExitFromTheGroup                                   = "ExitFromTheGroup"
	LeaveGroupWarning                                  = "LeaveGroupWarning"
	ConfirmLeavingGroup                                = "ConfirmLeavingGroup"
	MessageForEmployeeWhoLeavedGroup                   = "MessageForEmployeeWhoLeavedGroup"
	MessageForChiefAboutEmployeeLeftGroup              = "MessageForChiefAboutEmployeeLeftGroup"
	CreateNewTaskKey                                   = "CreateNewTask"
	ExpectingNewTaskTitleMessage                       = "ExpectingNewTaskTitleMessage"
	NewTaskTitleAcceptedExpectingNewTaskDesc           = "NewTaskTitleAcceptedExpectingNewTaskDesc"
	SkipKeyTitle                                       = "SkipKeyTitle"
	TaskDescriptionAccepted                            = "TaskDescriptionAccepted"
	AddWholeGroupToTaskKeyTitle                        = "AddWholeGroupToTaskKeyTitle"
	AddSomeEmployeesToTaskKeyTitle                     = "AddSomeEmployeesToTaskKeyTitle"
	AllEmployeesFromGroupSuccessfullyAddedToTask       = "AllEmployeesFromGroupSuccessfullyAddedToTask"
	NoEmployeesToAssignToTheTask                       = "NoEmployeesToAssignToTheTask"
	ToAssignKeyTitle                                   = "ToAssignKeyTitle"
	YouHaveBeenAssignedTheTask                         = "YouHaveBeenAssignedTheTask"
	ToSeeTaskDetailsKeyTitle                           = "ToSeeTaskDetailsKeyTitle"
	TaskDescriptionSkipped                             = "TaskDescriptionSkipped"
	TaskDetailsForEmployees                            = "TaskDetailsForEmployees"
	NoTaskDescription = "NoTaskDescription"
)

type MessageTitle string

var Messages = map[MessageTitle]map[Language]string{
	MessageForUnregisteredUsers: {
		EN: "You are unregistered user",
		UA: "Ви не зареєстрований користувач",
	},
	MessageIfUserAlreadyExists: {
		EN: "You are already registered user",
		UA: "Ви вже зареєстрований користувач",
	},
	MessageAfterFirstLanguageSelection: {
		EN: "Language selected. Enter your name:",
		UA: "Мова обрана. Введіть ваше ім'я:",
	},
	MessageAfterLanguageUpdate: {
		EN: "Language is chosen. To return to main menu click the button.",
		UA: "Мова обрана. Для повернення до головного меню натисніть кнопку.",
	},
	MessageAfterFirstNameEntering: {
		EN: "Congratulations, you now have a name:\n%s\nwhich will be displayed to other members of this bot. Next, you can create a new group and become its leader, or join another group and become a subordinate.",
		UA: "Вітаю, тепер ваше ім'я:\n%s\nвоно буде відображатись у інших учасників цього боту. Далі ви можете створити нову групу та стати її керівником, або приєднатися до іншої групи і стати у ній підлеглим.",
	},
	MessageAfterUserNameUpdate: {
		EN: "Congratulations, you now have a name:\n%s\nwhich will be displayed to other members of this bot. To return to the main menu, press the button.",
		UA: "Вітаю, тепер ваше ім'я:\n%s\nвоно буде відображатись у інших учасників цього боту. Для повернення до головного меню нажміть кнопку.",
	},
	ToMainMenuKey: {
		EN: "Main Menu",
		UA: "Головне меню",
	},
	CreateNewGroupKey: {
		EN: "Create New Group",
		UA: "Створити нову групу",
	},
	JoinToGroupKey: {
		EN: "Join the group",
		UA: "Приєднатись до групи",
	},
	MessageWithChatId: {
		EN: "To join the group, pass this code to the group leader:",
		UA: "Для приєднання до групи передайте цей код керівнику групи:",
	},
	MainMenuTitle: {
		EN: "Main Menu",
		UA: "Головне меню",
	},
	UserSettingsMenuTitle: {
		EN: "Profile Settings",
		UA: "Налаштування профілю",
	},
	ChangeLanguageKey: {
		EN: "Change the language",
		UA: "Змінити мову",
	},
	ChangeUserName: {
		EN: "Change the name",
		UA: "Змінити ім'я",
	},
	MessageBeforeNameChanging: {
		EN: "Enter a new name:",
		UA: "Введіть нове ім'я:",
	},
	MessageEnterNewGroupName: {
		EN: "Enter a name for the new group, or return to the main menu to cancel the operation:",
		UA: "Введіть назву для нової групи, або поверніться до головного меню для скасування цієї операції:",
	},
	MessageCreatedNewGroup: {
		EN: "A new group has been created with the name:\n%s\nTo add a new group member, enter the code they provided or return to the main menu.",
		UA: "Створено нову групу з назвою:\n%s\nЩоб додати нового учасника групи введіть наданий ним код або поверніться до головного меню.",
	},
	MessageEmployeeCodeBroken: {
		EN: "The code you entered does not match the specified format. Check the code for correctness and try again or return to the main menu.",
		UA: "Вказаний код не відповідає заданому формату. Перевірте правильність коду та спробуйте ще раз, або поверніться до головного меню.",
	},
	MessageNoEmployeeWithThisCode: {
		EN: "The specified code does not match any user. Check the code for correctness and try again or return to the main menu.",
		UA: "Вказаний код не відповідає жодному користувачу. Перевірте правильність коду та спробуйте ще раз, або поверніться до головного меню.",
	},
	MessageEmployeeCodeEqualsChiefCode: {
		EN: "The code you enter corresponds to you as a user. Check that the code is correct and try again or return to the main menu.",
		UA: "Введений код відповідає вам як користувачу. Перевірте правильність коду та спробуйте ще раз, або поверніться до головного меню.",
	},
	MessageEmployeeAddingSuccess: {
		EN: "User\n%s\nhas been successfully added to the group. Add another user or return to the main menu:",
		UA: "Користувача\n%s\nуспішно додано до групи. Додайте ще одного або поверніться до головного меню:",
	},
	GroupMenuKey: {
		EN: "Groups menu",
		UA: "Меню груп",
	},
	GroupMenuTitle: {
		EN: "Groups menu",
		UA: "Меню груп",
	},
	ShowAllChiefsGroups: {
		EN: "Show all groups (as a leader)",
		UA: "Показати всі групи (як керівник)",
	},
	RenameGroupKey: {
		EN: "Change the name",
		UA: "Змінити назву",
	},
	RenameGroupTitle: {
		EN: "Enter a new name for the group, or return to the main menu",
		UA: "Введіть нову назву для групи, або поверніться до головного меню",
	},
	MessageNewGroupNameAccepted: {
		EN: "Group name '%s' has been changed to '%s'",
		UA: "Назву групи '%s' змінено на '%s'",
	},
	ShowAllEmployeesFromGroupWithIdKey: {
		EN: "Display group members",
		UA: "Відобразити учасників групи",
	},
	AddNewEmployeeToExistingGroup: {
		EN: "To add a new member to group '%s', enter the code they provided or return to the main menu.",
		UA: "Щоб додати нового учасника до групи '%s' введіть наданий ним код або поверніться до головного меню.",
	},
	NoEmployeesInTheGroup: {
		EN: "No users in the group '%s'",
		UA: "В групі '%s' відсутні учасники",
	},
	DeleteEmployeeFromGroupKeyText: {
		EN: "Remove from the group",
		UA: "Видалити з групи",
	},
	CopyEmployeeToAnotherGroupKeyText: {
		EN: "Add a member to another group",
		UA: "Додати учасника до іншої групи",
	},
	MoveEmployeeToAnotherGroupKeyText: {
		EN: "Move a member to another group",
		UA: "Перемістити учасника до іншої групи",
	},
	NoGroups: {
		EN: "There are no groups.",
		UA: "Групи відсутні.",
	},
	ShownMembersOfTheGroupWithName: {
		EN: "Members of the group '%s':",
		UA: "Учасники групи '%s':",
	},
	YouAreDeletedFromGroup: {
		EN: "You have been removed from the group: %s",
		UA: "Вас видалено з групи: %s",
	},
	EmployeeHaveBeenDeletedFromGroup: {
		EN: "Member %s has been removed from group %s.",
		UA: "Учасника %s видалено з групи %s.",
	},
	AddNewGroupFromGroupListMenu: {
		EN: "To create a new group, enter its name or return to the main menu.",
		UA: "Щоб створити нову групу введіть її назву або поверніться до головного меню.",
	},
	ShowAllEmployeeGroups: {
		EN: "Show all groups (as a member)",
		UA: "Показати всі групи (як учасник)",
	},
	MessageWhenCopyEmployeeToAnotherGroup: {
		EN: "Add %s to the %s group %s?",
		UA: "Додати %s до групи %s?",
	},
	CopeEmployeeKey: {
		EN: "Add",
		UA: "Додати",
	},
	DeleteGroupKey: {
		EN: "Delete the group",
		UA: "Видалити групу",
	},
	WarningMessageBeforeDeletingEmployeeFromGroup: {
		EN: "Are you sure you want to remove employee %s from group %s?",
		UA: "Ви впевнені, що хочете видалити працівника %s з групи %s?",
	},
	ConfirmationButtonTextForDeletingEmployeeFromGroup: {
		EN: "Confirm the deletion",
		UA: "Підтвердити видалення",
	},
	ConfirmationButtonTextForDeletingGroup: {
		EN: "Confirm the deletion",
		UA: "Підтвердити видалення",
	},
	WarningBeforeGroupDeleting: {
		EN: "Are you sure you want to delete group %s?\nWarning: All assignments of this group will be lost!",
		UA: "Ви впевнені, що хочете видалити групу %s?\nУвага! Буде втрачено всі доручення цієї групи!",
	},
	GroupIsDeleted: {
		EN: "Group %s has been deleted",
		UA: "Групу %s видалено",
	},
	MessageToChiefUserCopied: {
		EN: "User %s has been added to group %s",
		UA: "Користувача %s додано до групи %s",
	},
	MessageToEmployeeCopied: {
		EN: "You have been added to group %s",
		UA: "Вас додано до групи %s",
	},
	GoToMainMenu: {
		EN: "Return to the main menu",
		UA: "Повернутись до головного меню",
	},
	ExitFromTheGroup: {
		EN: "Leave the group",
		UA: "Вийти з групи",
	},
	LeaveGroupWarning: {
		EN: "Are you sure you want to leave the '%s' group?",
		UA: "Ви впевнені, що хочете вийти з групи '%s'?",
	},
	ConfirmLeavingGroup: {
		EN: "Confirm exit",
		UA: "Підтвердити вихід",
	},
	MessageForEmployeeWhoLeavedGroup: {
		EN: "You have left the '%s' group",
		UA: "Ви вийшли з групи '%s'",
	},
	MessageForChiefAboutEmployeeLeftGroup: {
		EN: "User %s has left group %s",
		UA: "Користувач %s вийшов з групи %s",
	},
	CreateNewTaskKey: {
		EN: "Create new task",
		UA: "Створити нову задачу",
	},
	ExpectingNewTaskTitleMessage: {
		EN: "Enter a title for the task. To cancel, return to the Main Menu.",
		UA: "Введіть заголовок завдання. Для скасування поверніться до Головного Меню.",
	},
	NewTaskTitleAcceptedExpectingNewTaskDesc: {
		EN: "Task '%s' has been created. Enter a description of the task or click the button to skip.",
		UA: "Задачу '%s' створено. Введіть опис задачі або натисніть кнопку, щоб пропустити.",
	},
	SkipKeyTitle: {
		EN: "Without description",
		UA: "Без опису",
	},
	TaskDescriptionAccepted: {
		EN: "Description\n%s\nis accepted and added to the task. Assign it to the entire group or to individual employees?",
		UA: "Опис\n%s\nприйнято та додано до задачі. Доручити її виконання всій групі чи окремим працівникам?",
	},
	AddWholeGroupToTaskKeyTitle: {
		EN: "Assign the entire group",
		UA: "Доручити всій групі",
	},
	AddSomeEmployeesToTaskKeyTitle: {
		EN: "Assign individual employees",
		UA: "Доручити окремим працівникам",
	},
	AllEmployeesFromGroupSuccessfullyAddedToTask: {
		EN: "Task '%s' is assigned to all members of group '%s'.",
		UA: "Виконання задачі '%s' доручено всім членам групи '%s'.",
	},
	NoEmployeesToAssignToTheTask: {
		EN: "There are no employees who can be entrusted with this task",
		UA: "Відсутні працівники, яким можна доручити виконання цієї задачі",
	},
	ToAssignKeyTitle: {
		EN: "Assign",
		UA: "Доручити",
	},
	YouHaveBeenAssignedTheTask: {
		EN: "You have been assigned the following task:\n'%s'.",
		UA: "Вам доручено виконання задачі:\n'%s'.",
	},
	ToSeeTaskDetailsKeyTitle: {
		EN: "View details",
		UA: "Переглянути деталі",
	},
	TaskDescriptionSkipped: {
		EN: "Adding a description to the task has been skipped. You can do it later. Should I assign the task to the entire group or to individual employees?",
		UA: "Додавання опису до задачі пропущено. Ви можете зробити це пізніше. Доручити виконання задачі всій групі чи окремим працівникам?",
	},
	TaskDetailsForEmployees: {
		EN: "Title:\n%s\n\nDescription:\n%s\n\nExecutors:\n%s",
		UA: "Назва:\n%s\n\nОпис:\n%s\n\nВиконавці:\n%s",
	},
	NoTaskDescription: {
		EN: "No description available",
		UA: "Опис відсутній",
	},
}

func ReturnMessageByLanguage(msgTitle MessageTitle, lng Language) (string, error) {
	msg, ok := Messages[msgTitle]
	if !ok {
		return "", fmt.Errorf("unknown message title: %s", msgTitle)
	}
	str, ok := msg[lng]
	if !ok {
		return msg[EN], fmt.Errorf("unknown language: %s", lng)
	}
	return str, nil
}
