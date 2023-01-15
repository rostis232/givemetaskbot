package service

import "github.com/rostis232/givemetaskbot/internal/repository"

type TaskService struct {
	repository repository.Task
}

func NewTaskService(repository repository.Task) *TaskService {
	return &TaskService{repository: repository}
}
