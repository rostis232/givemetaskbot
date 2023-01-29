package telegram

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/keyboards"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/state_service"
	"log"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	//TODO: what with keyboard?
	msg.ReplyMarkup = keyboards.RemoveKeyboard
	user, err := b.service.GetUser(message.Chat.ID)

	switch {
	case err == sql.ErrNoRows:
		msg, err = b.service.NewUserRegistration(message.Chat.ID)
		if err != nil {
			return err
		}
	case user.Status == state_service.Expecting_new_user_name:
		msg, err = b.service.SetUserName(&user, message)
		if err != nil {
			return err
		}
	case user.Status == state_service.Changing_user_name:
		msg, err = b.service.SetUserName(&user, message)
		if err != nil {
			return err
		}
	}

	if _, err = b.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, messages.UnknownCommand)
	//TODO: what with keyboard?
	msg.ReplyMarkup = keyboards.RemoveKeyboard
	user, err := b.service.GetUser(message.Chat.ID)
	if err != nil {
		log.Println(err)
	}

	switch message.Command() {
	case keys.CommandStart:
		switch {
		case err == sql.ErrNoRows:
			msg, err = b.service.NewUserRegistration(message.Chat.ID)
			if err != nil {
				return err
			}
		default:
			msg.Text, err = messages.ReturnMessageByLanguage(messages.MessageIfUserAlreadyExists, user.Language)
			if err != nil {
				msg.Text = fmt.Sprintf(messages.UnknownError, err)
			}
		}
	}

	if _, err = b.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Вітаю в телегра-боті Доручення-бот. "+
		"Для продовження натисніть кнопку зареєструватись, "+
		"для отримання додаткової інформації - кнопку довідка.")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Зареєструватись", keys.Register),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Довідка", keys.Info),
		),
	)
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
