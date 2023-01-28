package messages

import (
	"errors"
	"fmt"
)

const (
	EN                                 = Language("EN")
	UA                                 = Language("UA")
	UnknownError                       = "Unknown error: %s. Contact with administrator @rostis232"
	UnknownCommand                     = "Unknown command"
	MessageForUnregisteredUsers        = MessageTitle("MessageForUnregisteredUsers")
	MessageIfUserAlreadyExists         = MessageTitle("MessageIfUserAlreadyExists")
	MessageAfterFirstLanguageSelection = "MessageAfterFirstLanguageSelection"
	MessageAfterLanguageUpdate         = "MessageAfterLanguageUpdate"
)

type Language string

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
		EN: "Language selected. To return to main menu click the button.",
		UA: "Мова обрана. Для повернення до головного меню нажміть кнопку.",
	},
}

func ReturnMessageByLanguage(msgTitle MessageTitle, lng Language) (string, error) {
	msg, ok := Messages[msgTitle]
	if !ok {
		return "", errors.New(fmt.Sprintf("unknown message title: %s", msgTitle))
	}
	str, ok := msg[lng]
	if !ok {
		return msg[EN], errors.New(fmt.Sprintf("unknown lanhuage: %s", lng))
	}
	return str, nil
}
