package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/keyboards"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/user_state"
	"log"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	//TODO: what with keyboard?
	msg.ReplyMarkup = keyboards.RemoveKeyboard

	user_status, ok := user_state.UsersStates[message.Chat.ID]

	switch {

	case !ok:
		//no user in map: add user to map if it`s registered

	}

	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	switch message.Command() {
	case keys.CommandStart:
		if err := b.handleStartCommand(message); err != nil {
			return errors.New(fmt.Sprintf("error while handling start command: %s", err))
		}
	case "help":
		if err := b.handleHelpCommand(message); err != nil {
			return errors.New(fmt.Sprintf("error while handling help command: %s", err))
		}
	default:
		if err := b.handleUnknownCommand(message); err != nil {
			return errors.New(fmt.Sprintf("error while handling unknown command: %s", err))
		}
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

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
			tgbotapi.NewInlineKeyboardButtonData("Test", fmt.Sprintf("chat_id=%d", message.Chat.ID)),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		),
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, "This bot will help you to organize your work")
	msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command")
	msg.ReplyToMessageID = message.MessageID
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
