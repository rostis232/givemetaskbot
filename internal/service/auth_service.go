package service

import (
	"github.com/rostis232/givemetaskbot/internal/entities"
	"github.com/rostis232/givemetaskbot/internal/repository"
)

type AuthService struct {
	repository repository.Authorisation
}

func NewAuthService(repository repository.Authorisation) *AuthService {
	return &AuthService{repository: repository}
}

func (u *AuthService) CheckIfUserIsRegistered(chatId int64) bool {
	for _, v := range entities.UserStates {
		if v.ChatId == chatId {
			return true
		}
	}
	//TODO: Check BD
	return false
}

func (u *AuthService) InsertUserToUserStates(chatId int64) {
	entities.UserStates = append(entities.UserStates, entities.UserState{ChatId: chatId})
}
