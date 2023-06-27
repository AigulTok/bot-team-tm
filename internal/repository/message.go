package repository

import (
	"bot-team-tm/models"
	"log"

	"gorm.io/gorm"
)

type Message interface {
	GetAllChatID() ([]int, error)
}

type MessageRepos struct {
	db *gorm.DB
}

func newMessageRepos(db *gorm.DB) *MessageRepos {
	return &MessageRepos{
		db: db,
	}
}

func (m *MessageRepos) GetAllChatID() ([]int, error) {
	log.Println("Получаем все идентификатор чатов...")

	var chatIDs []models.User

	res := m.db.Select("chat_id").Find(&chatIDs)

	var IDs []int
	for i := 0; i < len(chatIDs); i++ {
		IDs = append(IDs, int(chatIDs[i].ChatId))
	}
	
	return IDs, res.Error
}
