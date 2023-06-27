package service

import (
	"bot-team-tm/internal/repository"
	"bot-team-tm/models"
)

type User interface {
	AddNewUser(user models.User) error
	Get(userID int64) (*models.User, error)
}

type UserService struct {
	repo repository.User
}

func newUserService(r repository.User) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Get(userID int64) (*models.User, error) {
	return s.repo.GetUserByChatId(userID)
}

func (s *UserService) AddNewUser(user models.User) error {
	if s.repo.CheckUserByChatId(user.ChatId) {
		if err := s.repo.AddUser(user); err != nil {
			return err
		}
	}

	return nil
}
