package service

import (
	"bot-team-tm/internal/repository"
)

type Service struct {
	User
	Message
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    newUserService(repo.User),
		Message: newMessageService(repo.Message),
	}
}
