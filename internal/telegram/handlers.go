package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const commandStart = "start"

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.MessageConfig{}

	switch b.service.CheckIfUserIsRegistered(message.Chat.ID) {
	case true:
		msg = tgbotapi.NewMessage(message.Chat.ID, "You`re registered user")
	case false:
		msg = tgbotapi.NewMessage(message.Chat.ID, "You aren`t registered user")
	}

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	switch message.Command() {
	case commandStart:
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

	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome in Bot. You need to choice your role")
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "This bot will help you to organize your work")
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
