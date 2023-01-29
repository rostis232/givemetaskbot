package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/service"
	"log"
	"strings"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func NewBot(token string, service *service.Service) (*Bot, error) {
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

	//go b.handleUpdates(updates)

	b.handleUpdates(updates)

	//b.sendMessages()

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					log.Println(fmt.Sprintf("handling command error: %s", err))
					continue
				}
			} else {
				if err := b.handleMessage(update.Message); err != nil {
					log.Println(fmt.Sprintf("handling message error: %s", err))
					continue
				}
			}

		} else if update.CallbackQuery != nil {
			//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			//if _, err := b.bot.Request(callback); err != nil {
			//	return err
			//}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, messages.UnknownError)

			user, err := b.service.GetUser(update.CallbackQuery.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}

			switch {
			case strings.Contains(update.CallbackQuery.Data, messages.LanguageKey):
				_, lng, ok := strings.Cut(update.CallbackQuery.Data, ":")
				if !ok {
					log.Println("error while cutting string with language")
				}
				msg, err = b.service.SelectLanguage(&user, messages.Language(lng))
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case update.CallbackQuery.Data == keys.JoinTheExistGroup:
				msg, err = b.service.ShowChatId(&user)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			}

			if _, err := b.bot.Send(msg); err != nil {
				log.Println(err)
				continue
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
