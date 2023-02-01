package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"log"
)

var (
	// LanguageKeyboard hold keys for language choosing
	LanguageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Українська", messages.LanguageKey+string(messages.UA)),
			tgbotapi.NewInlineKeyboardButtonData("English", messages.LanguageKey+string(messages.EN)),
		),
	)
	// RemoveKeyboard removes any keyboards
	RemoveKeyboard = tgbotapi.NewRemoveKeyboard(true)
)

// NewToMainMenuKeyboard returns new inline keyboard with only one button to MainMenu
func NewToMainMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	keyText, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	toMainMenuKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyText, keys.GoToMainMenu),
		),
	)
	return toMainMenuKeyboard
}

// NewKeyboardChooseCreateOrJoinGroup returns new inline keyboard with two keys: create new
// group or join to exist group
func NewKeyboardChooseCreateOrJoinGroup(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	keyCreate, err := messages.ReturnMessageByLanguage(messages.CreateNewGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyJoin, err := messages.ReturnMessageByLanguage(messages.JoinToGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyCreate, keys.CreateNewGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyJoin, keys.JoinTheExistGroup),
		),
	)
	return keyboard
}

// NewMainMenuKeyboard returns new inline keyboard for Main Menu
func NewMainMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	userMenuTitle, err := messages.ReturnMessageByLanguage(messages.UserSettingsMenuTitle, user.Language)
	if err != nil {
		log.Println(err)
	}
	groupMenuTitle, err := messages.ReturnMessageByLanguage(messages.GroupMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(userMenuTitle, keys.GoToUserMenuSettings),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(groupMenuTitle, keys.GoToGroupsMenu),
		),
	)
	return keyboard

}

// NewUserSettingsMenuKeyboard returns new inline keyboard for User Settings Menu
func NewUserSettingsMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	changeLanguageKeyTitle, err := messages.ReturnMessageByLanguage(messages.ChangeLanguageKey, user.Language)
	if err != nil {
		log.Println(err)
	}
	changeUserNameKeyTitle, err := messages.ReturnMessageByLanguage(messages.ChangeUserName, user.Language)
	if err != nil {
		log.Println(err)
	}
	keyText, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(changeLanguageKeyTitle, keys.GoToChangeLanguageMenu),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(changeUserNameKeyTitle, keys.GoToChangeUserName),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyText, keys.GoToMainMenu),
		),
	)
	return keyboard
}

func NewGroupMenuKeyboard(user *entities.User) tgbotapi.InlineKeyboardMarkup {
	createNewGroupKey, err := messages.ReturnMessageByLanguage(messages.CreateNewGroupKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	mainMenuKey, err := messages.ReturnMessageByLanguage(messages.ToMainMenuKey, user.Language)
	if err != nil {
		log.Println(err)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(createNewGroupKey, keys.CreateNewGroup),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(mainMenuKey, keys.GoToMainMenu),
		),
	)

	return keyboard
}
