package repository

import (
	"bot-team-tm/models"
	"fmt"

	"gorm.io/gorm"
)

type User interface {
	AddUser(user models.User) error
	GetUserByChatId(chatId int64) (*models.User, error)
	CheckUserByChatId(chatId int64) bool
}

type UserRepos struct {
	db *gorm.DB
}

func newUserRepos(db *gorm.DB) *UserRepos {
	return &UserRepos{
		db: db,
	}
}

func (r *UserRepos) AddUser(user models.User) error {
	res := r.db.Create(&user)
	if res.Error != nil {
		return fmt.Errorf("repository: AddUser %s", res.Error)
	}
	return nil
}

func (r *UserRepos) GetUserByChatId(chatId int64) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, models.User{ChatId: chatId})
	if result.Error != nil {
		return nil, fmt.Errorf("repository: GetUserByChatId %s", result.Error)
	}
	return &user, nil
}

func (r *UserRepos) CheckUserByChatId(chatId int64) bool {
	var user models.User
	var exists bool
	_ = r.db.Model(user).Select("count(*) > 0").Where("chat_id = ?", chatId).Find(&exists).Error

	return !exists
}
