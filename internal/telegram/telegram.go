package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Service interface {
	CheckIfUserIsRegistered(chatId int64) bool
}

type Bot struct {
	bot     *tgbotapi.BotAPI
	service Service
}

func NewBot(token string, service Service) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{bot: bot, service: service}, nil
}

func (b *Bot) Start() error {
	b.bot.Debug = false

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	//if err := b.handleUpdates(updates); err != nil {
	//	return errors.New(fmt.Sprintf("handling updates error: %s", err))
	//}

	go b.handleUpdates(updates)

	b.sendMessages()

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					return errors.New(fmt.Sprintf("handling command error: %s", err))
				}
			} else {
				if err := b.handleMessage(update.Message); err != nil {
					return errors.New(fmt.Sprintf("handling message error: %s", err))
				}
			}

		}
	}
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) sendMessages() {
	for {
		//for _, v := range repository.ChatIDs {
		//	msg := tgbotapi.NewMessage(v, "Still there?")
		//	b.bot.Send(msg)
		//}
		//time.Sleep(time.Minute)
	}
}
