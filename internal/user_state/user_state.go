package user_state

import (
	"errors"
	"github.com/rostis232/givemetaskbot/internal/entities"
)

var UsersStates = make(map[int64]UserState)

type UserState struct {
	user        entities.User
	userStatus  int
	activeGroup int
	activeTask  int
}

type StateService struct {
}

func NewStateService() *StateService {
	return &StateService{}
}

func (ss *StateService) GetUserState(chatId int64) (UserState, error) {
	userState, ok := UsersStates[chatId]
	if !ok {
		return UserState{}, errors.New("no chatId in users stats")
	}
	return userState, nil
}

func (ss *StateService) AddUserState(chatId int64) (UserState, error) {
	userState, ok := UsersStates[chatId]
	if !ok {
		return userState, errors.New("chatId already exists")
	}
}

func (ss *StateService) UpdateUserStatus(chatId int64, userStatus int) error {
	userState, ok := UsersStates[chatId]
	if !ok {
		return errors.New("no chatId in users stats")
	}
	userState.userStatus = userStatus
	UsersStates[chatId] = userState

	return nil
}

func (ss *StateService) UpdateActiveGroup(chatId int64, activeGroup int) error {
	userState, ok := UsersStates[chatId]
	if !ok {
		return errors.New("no chatId in users stats")
	}
	userState.activeGroup = activeGroup
	UsersStates[chatId] = userState

	return nil
}

func (ss *StateService) UpdateActiveTask(chatId int64, activeTask int) error {
	userState, ok := UsersStates[chatId]
	if !ok {
		return errors.New("no chatId in users stats")
	}
	userState.activeTask = activeTask
	UsersStates[chatId] = userState

	return nil
}
