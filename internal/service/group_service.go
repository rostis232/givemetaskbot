package service

import "github.com/rostis232/givemetaskbot/internal/repository"

type GroupService struct {
	repository repository.Group
}

func NewGroupService(repository repository.Group) *GroupService {
	return &GroupService{repository: repository}
}
