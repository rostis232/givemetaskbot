package telegram

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rostis232/givemetaskbot/internal/keys"
	"github.com/rostis232/givemetaskbot/internal/messages"
	"github.com/rostis232/givemetaskbot/internal/state_service"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s - %d", message.From.UserName, message.Text, message.Chat.ID)

	user, err := b.service.GetUser(message.Chat.ID)

	switch {
	case err == sql.ErrNoRows:
		if err = b.service.NewUserRegistration(message.Chat.ID); err != nil {
			return err
		}
	case user.Status == state_service.Expecting_new_user_name:
		if err = b.service.SetUserName(&user, message); err != nil {
			return err
		}
	case user.Status == state_service.Changing_user_name:
		if err = b.service.SetUserName(&user, message); err != nil {
			return err
		}
	case user.Status == state_service.Expecting_new_group_name:
		if err = b.service.CreatingNewGroup(&user, message); err != nil {
			return err
		}
	case user.Status == state_service.Adding_employee_to_group:
		//Adding employee to group
		if err = b.service.AddingEmployeeToGroup(&user, message); err != nil {
			return err
		}
	case strings.HasPrefix(message.Text, keys.ChatIdPrefix):
		//Adding employee to group
		if err = b.service.AddingEmployeeToGroup(&user, message); err != nil {
			return err
		}
	case user.Status == state_service.Changing_group_name && user.ActiveGroup != 0:
		//Changing name of existing group
		if err = b.service.UpdateGroupName(&user, message.Text); err != nil {
			return err
		}
	case user.Status == state_service.Expecting_new_task_title && user.ActiveGroup != 0:
		if err = b.service.CreateNewTaskGotTitle(&user, message.Text); err != nil {
			return err
		}
	case user.Status == state_service.Expecting_new_task_description && user.ActiveTask != 0:
		if err = b.service.UpdateTaskDescription(&user, message.Text); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	log.Printf("[%s] %s - %d", message.From.UserName, message.Text, message.Chat.ID)

	user, err := b.service.GetUser(message.Chat.ID)
	if err != nil {
		log.Println(err)
	}

	switch message.Command() {
	case keys.CommandStart:
		switch {
		case err == sql.ErrNoRows:
			if err = b.service.NewUserRegistration(message.Chat.ID); err != nil {
				return err
			}
		default:
			text, err := messages.ReturnMessageByLanguage(messages.MessageIfUserAlreadyExists, user.Language)
			if err != nil {
				text = fmt.Sprintf(messages.UnknownError, err)
			}
			msg := tgbotapi.NewMessage(message.Chat.ID, text)
			if _, err = b.bot.Send(msg); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Bot) handleCallback(callbackQuery *tgbotapi.CallbackQuery) error {
	log.Printf("[%s] %s - %d", callbackQuery.From.UserName, callbackQuery.Data, callbackQuery.Message.Chat.ID)
	user, err := b.service.GetUser(callbackQuery.Message.Chat.ID)
	switch {
	case err == sql.ErrNoRows:
		if err = b.service.NewUserRegistration(callbackQuery.Message.Chat.ID); err != nil {
			return err
		}
	case err != nil:
		log.Println(err)
		return err
	case strings.Contains(callbackQuery.Data, messages.LanguageKey):
		//?????????????? ???????????? ????????
		if err = b.service.SelectLanguage(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.JoinTheExistGroup:
		//?????????????????? ?????? ?????? ???????????????????? ???? ??????????
		if err = b.service.ShowChatId(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.GoToMainMenu:
		//?????????????????? ?????????????? ????????
		if err = b.service.MainMenu(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.GoToUserMenuSettings:
		//?????????????????? ???????? ?????????????????????? ??????????????????????
		if err = b.service.UserMenuSettings(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.GoToChangeLanguageMenu:
		//?????????????????? ???????? ?????????? ????????
		if err = b.service.ChangeLanguageMenu(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.GoToChangeUserName:
		//?????????????? ???????????? ???????? ????'??
		if err = b.service.UserNameChanging(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.CreateNewGroup:
		//?????????????? ???????????? ????'?? ?????????? ??????????
		if err = b.service.AskingForNewGroupTitle(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.GoToGroupsMenu:
		//???????? ????????
		if err = b.service.GroupsMenu(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.ShowAllChiefsGroups:
		if err = b.service.ShowAllChiefsGroups(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.ShowAllEmployeeGroups:
		err = b.service.ShowAllEmployeeGroups(&user)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.RenameGroupWithId):
		//?????????? ???????????? ?????????? ?????? ??????????, ???? ?????? ??????????

		if err = b.service.AskingForUpdatedGroupName(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.ShowAllEmployeesFromGroupWithId):
		//???????????????? ?????????? ???? ?????????? ???????? ?????????????????? ??????????
		if err = b.service.ShowAllEmploysFromGroup(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.EmployeeIDtoDeleteFromGroup):
		//???????????????? ?????????? ???? ?????????????????? ???????????????? ??????????, ?????????????????? ??????????????????????????
		if err = b.service.WarningBeforeDeletingEmployeeFromGroup(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.ConfirmDeletingEmployeeId):
		//???????????????? ?????????????????????????? ???? ?????????????????? ???????????????? ??????????
		if err = b.service.DeleteEmployeeFromGroup(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.EmployeeIDtoCopyToANotherGroup):
		//Request to copy employee to another group
		if err = b.service.CopyEmployeeToAnotherGroup(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.DeleteGroup):
		//Get request to delete group, pending confirmation
		if err = b.service.WarningBeforeGroupDeleting(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.ConfirmDeletingGroup):
		//Get confirmation to delete group
		if err = b.service.DeleteGroup(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.CopyEmployeeGroupID):
		//Got information with group id to copy in employee
		if err = b.service.ConfirmCopyEmployeeToAnotherGroup(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.LeaveGroupID):
		if err = b.service.LeaveGroupBeforeConfirmation(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.ConfirmLeavingGroup:
		if err := b.service.LeaveGroupWithConfirmation(&user); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.CreateNewTaskKeyData):
		if err := b.service.CreateNewTaskAskingTitle(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.AddWholeGroupToTaskKeyData:
		if err := b.service.AddAllEmployeesToTask(&user); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.AddSomeEmployeesToTaskKeyData:
		if err := b.service.AddSomeEmployeesToTaskShowList(&user, ""); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.AssignToEmployeeWithID):
		if err := b.service.AssignEmployee(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case callbackQuery.Data == keys.SkipDescritionEntering:
		if err := b.service.SkipDescritionEntering(&user); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.ToSeeTaskDetailsTaskIDForEmployee):
		if err := b.service.ShowTaskDetailsForEmployee(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.ToSeeTaskDetailsTaskIDForChief):
		if err := b.service.ShowTaskDetailsForChief(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}	
	case strings.Contains(callbackQuery.Data, keys.ShowAllTasksByGorupIDForChief):
		if err := b.service.ShowAllGroupTasksForChief(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.ShowAllTasksByGorupIDForEmployee):
		if err := b.service.ShowAllGroupTasksForEmployee(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.MarkTaskAsClosedForChiefKeyData):
		if err := b.service.MarkTaskAsComplete(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.DeleteTaskForChiefKeyData):
		if err := b.service.BeforeTaskDeleting(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.ConfirmTaskDeleting):
		if err := b.service.DeleteTask(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.AddExecutorForChiefKeyData):
		if err := b.service.AddSomeEmployeesToTaskShowList(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.RemoveExecutorForChiefKeyData):
		if err := b.service.DeleteExecutors(&user, callbackQuery.Data); err != nil {
			log.Println(err)
		}
	case strings.Contains(callbackQuery.Data, keys.GroupDetailsForChief):
		if err := b.service.ShowGroupDetailsForChief(&user, callbackQuery.Data); err != nil {
			return err
		}
	case strings.Contains(callbackQuery.Data, keys.GroupDetailsForEmployee):
		if err := b.service.ShowGroupDetailsForEmployee(&user, callbackQuery.Data); err != nil {
			return err
		}
	}
	return nil
}
