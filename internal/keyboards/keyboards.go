package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
)

var (
	StartKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(keys.Register),
			tgbotapi.NewKeyboardButton(keys.Info),
		),
	)

	LanguageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Українська", "lang:"+string(messages.UA)),
			tgbotapi.NewInlineKeyboardButtonData("English", "lang:"+string(messages.EN)),
		),
	)

	RemoveKeyboard = tgbotapi.NewRemoveKeyboard(true)
)
