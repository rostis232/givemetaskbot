package telegram

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/keyboards"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/state_service"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s - %d", message.From.UserName, message.Text, message.Chat.ID)

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
	case user.Status == state_service.Expecting_new_group_name:
		msg, err = b.service.CreatingNewGroup(&user, message)
		if err != nil {
			return err
		}
	case user.Status == state_service.Adding_employee_to_group:
		//Adding employee to group
		msg, err = b.service.AddingEmployeeToGroup(&user, message)
		if err != nil {
			return err
		}
	case strings.HasPrefix(message.Text, keys.ChatIdPrefix):
		//Adding employee to group
		msg, err = b.service.AddingEmployeeToGroup(&user, message)
		if err != nil {
			return err
		}
	case user.Status == state_service.Changing_group_name && user.ActiveGroup != 0:
		//Changing name of existing group
		msg, err = b.service.UpdateGroupName(&user, message.Text)
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
	log.Printf("[%s] %s - %d", message.From.UserName, message.Text, message.Chat.ID)

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

// func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

// 	msg := tgbotapi.NewMessage(message.Chat.ID, "Вітаю в телегра-боті Доручення-бот. "+
// 		"Для продовження натисніть кнопку зареєструватись, "+
// 		"для отримання додаткової інформації - кнопку довідка.")
// 	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("Зареєструватись", keys.Register),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("Довідка", keys.Info),
// 		),
// 	)
// 	_, err := b.bot.Send(msg)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (b *Bot) handleCallback(callbackQuery *tgbotapi.CallbackQuery) error {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)

			user, err := b.service.GetUser(callbackQuery.Message.Chat.ID)
			if err != nil {
				log.Println(err)
			}

			switch {
			case strings.Contains(callbackQuery.Data, messages.LanguageKey):
				//Обробка вибору мови
				showMainMenuKeyboard := true
				if user.Status == state_service.Expecting_language {
					showMainMenuKeyboard = false
				}
				_, lng, ok := strings.Cut(callbackQuery.Data, ":")
				if !ok {
					log.Println("error while cutting string with language")
				}
				msg, err = b.service.SelectLanguage(&user, messages.Language(lng))
				if showMainMenuKeyboard {
					msg.ReplyMarkup = keyboards.NewToMainMenuKeyboard(&user)
				}
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case callbackQuery.Data == keys.JoinTheExistGroup:
				//Показуємо код для приєднання до групи
				msg, err = b.service.ShowChatId(&user)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case callbackQuery.Data == keys.GoToMainMenu:
				//Показуємо головне меню
				msg, err = b.service.MainMenu(&user)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}

			case callbackQuery.Data == keys.GoToUserMenuSettings:
				//Показуємо меню налаштувань користувача
				title, err := messages.ReturnMessageByLanguage(messages.UserSettingsMenuTitle, user.Language)
				if err != nil {
					log.Println(err)
				}
				msg = tgbotapi.NewMessage(user.ChatId, title)
				msg.ReplyMarkup = keyboards.NewUserSettingsMenuKeyboard(&user)

			case callbackQuery.Data == keys.GoToChangeLanguageMenu:
				//Показуємо меню зміни мови
				title, err := messages.ReturnMessageByLanguage(messages.ChangeLanguageKey, user.Language)
				if err != nil {
					log.Println(err)
				}
				msg = tgbotapi.NewMessage(user.ChatId, title)
				msg.ReplyMarkup = keyboards.LanguageKeyboard
			case callbackQuery.Data == keys.GoToChangeUserName:
				//Просимо ввести нове ім'я
				msg, err = b.service.UserNameChanging(&user)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case callbackQuery.Data == keys.CreateNewGroup:
				//Просимо ввести ім'я нової групи
				msg, err = b.service.AskingForNewGroupTitle(&user)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case callbackQuery.Data == keys.GoToGroupsMenu:
				//Меню груп
				msg, err = b.service.GroupsMenu(&user)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case callbackQuery.Data == keys.ShowAllChiefsGroups:
				msg, err = b.service.ShowAllChiefsGroups(&user)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case strings.Contains(callbackQuery.Data, keys.RenameGroupWithId):
				//Запит нового імені для групи, що вже існує
				_, groupId, ok := strings.Cut(callbackQuery.Data, keys.RenameGroupWithId)
				if !ok {
					log.Println("Помилка отримання ID групи")
				}
				groupIdInt, err := strconv.Atoi(groupId)
				if err != nil {
					log.Println("Помилка отримання ID групи")
				}
				msg, err = b.service.AskingForUpdatedGroupName(&user, groupIdInt)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case strings.Contains(callbackQuery.Data, keys.ShowAllEmployeesFromGroupWithId):
				//Отримано запит на відображення всіх учасників групи
				_, groupId, ok := strings.Cut(callbackQuery.Data, keys.ShowAllEmployeesFromGroupWithId)
				if !ok {
					log.Println("Помилка отримання ID групи")
				}
				groupIdInt, err := strconv.Atoi(groupId)
				if err != nil {
					log.Println("Помилка отримання ID групи")
				}
				msg, err = b.service.ShowAllEmploysFromGroup(&user, groupIdInt)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}
			case strings.Contains(callbackQuery.Data, keys.EmployeeIDtoDeleteFromGroup):
				//Отримано запит на видалення учасника групи
				msg, err = b.service.DeleteEmployeeFromGroup(&user, callbackQuery.Data)
				if err != nil {
					log.Println(err)
					msg = tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, messages.UnknownError)
				}

			}

			if msg.Text != "" {
				if _, err := b.bot.Send(msg); err != nil {
					log.Println(err)
					return err
				}
			}
	return nil
}