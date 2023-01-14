package service

import "github.com/rostis232/givemetaskbot/internal/entities"

type Repository interface {
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CheckIfUserIsRegistered(chatId int64) bool {
	for _, v := range entities.UserStates {
		if v.ChatId == chatId {
			return true
		}
	}
	//TODO: Check BD
	return false
}
