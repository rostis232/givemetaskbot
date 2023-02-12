package telegram

import (
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/service"
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

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go b.handleUpdates(updates, wg)
	wg.Add(1)
	go b.sendMessages(wg)
	wg.Wait()

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel, wg *sync.WaitGroup) error {
	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					log.Printf("handling command error: %s", err)
					continue
				}
			} else {
				if err := b.handleMessage(update.Message); err != nil {
					log.Printf("handling message error: %s", err)
					continue
				}
			}
		} else if update.CallbackQuery != nil {
			//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "working")
			//if _, err := b.bot.Request(callback); err != nil {
			//	panic(err)
			//}
			if err := b.handleCallback(update.CallbackQuery); err != nil {
				log.Printf("handling callback error: %s", err)
				continue
			}
		}
	}
	wg.Done()
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) sendMessages(wg *sync.WaitGroup) {
	for i := range service.MsgChan {
		if _, err := b.bot.Send(i); err != nil {
			log.Println(err)
			continue
		}
	}
	wg.Done()
}
