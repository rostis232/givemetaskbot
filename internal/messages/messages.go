package messages

import (
	"errors"
	"fmt"
)

const (
	UnknownError                       = "Unknown error: %s. Contact with administrator @rostis232"
	UnknownCommand                     = "Unknown command"
	MessageForUnregisteredUsers        = "MessageForUnregisteredUsers"
	MessageIfUserAlreadyExists         = "MessageIfUserAlreadyExists"
	MessageAfterFirstLanguageSelection = "MessageAfterFirstLanguageSelection"
	MessageAfterLanguageUpdate         = "MessageAfterLanguageUpdate"
	MessageAfterFirstNameEntering      = "MessageAfterFirstNameEntering"
	MessageAfterUserNameUpdate         = "MessageAfterUserNameUpdate"
	ToMainMenuKey                      = "ToMainMenuKey"
	CreateNewGroupKey                  = "CreateNewGroupKey"
	JoinToGroupKey                     = "JoinToGroupKey"
	MessageWithChatId                  = "MessageWithChatId"
	MainMenuTitle                      = "MainMenuTitle"
	UserSettingsMenuTitle              = "UserSettingsMenuTitle"
	ChangeLanguageKey                  = "ChangeLanguageKey"
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
		EN: "Congratulations, you now have a name: %s, which will be displayed to other members of this bot. Next, you can create a new group and become its leader, or join another group and become a subordinate.",
		UA: "Вітаю, тепер ваше ім'я: %s, воно буде відображатись у інших учасників цього боту. Далі ви можете створити нову групу та стати її керівником, або приєднатися до іншої групи і стати у ній підлеглим.",
	},
	MessageAfterUserNameUpdate: {
		EN: "Congratulations, you now have a name: %s, which will be displayed to other members of this bot. To return to the main menu, press the button.",
		UA: "Вітаю, тепер ваше ім'я: %s, воно буде відображатись у інших учасників цього боту. Для повернення до головного меню нажміть кнопку.",
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
		EN: "To join the group, pass this code to the group leader:\n %s",
		UA: "Для приєднання до групи передайте цей код керівнику групи:\n %s",
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
}

func ReturnMessageByLanguage(msgTitle MessageTitle, lng Language) (string, error) {
	msg, ok := Messages[msgTitle]
	if !ok {
		return "", errors.New(fmt.Sprintf("unknown message title: %s", msgTitle))
	}
	str, ok := msg[lng]
	if !ok {
		return msg[EN], errors.New(fmt.Sprintf("unknown language: %s", lng))
	}
	return str, nil
}
