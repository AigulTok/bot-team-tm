package repository

import "gorm.io/gorm"

type Repository struct {
	User
	Message
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    newUserRepos(db),
		Message: newMessageRepos(db),
	}
}
