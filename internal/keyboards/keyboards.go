package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/keys"
)

var (
	StartKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(keys.Register),
			tgbotapi.NewKeyboardButton(keys.Info),
		),
	)

	RemoveKeyboard = tgbotapi.NewRemoveKeyboard(true)
)
