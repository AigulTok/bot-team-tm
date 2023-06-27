package service

import (
	"bot-team-tm/internal/repository"
)

type Message interface {
	GetAllChatID() ([]int, error)
}

type MessageService struct {
	repo repository.Message
}

func newMessageService(r repository.Message) *MessageService {
	return &MessageService{
		repo: r,
	}
}

func (m *MessageService) GetAllChatID() ([]int, error) {
	return m.repo.GetAllChatID()
}
