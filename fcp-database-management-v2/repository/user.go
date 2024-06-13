package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	// TODO: replace this
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	// TODO: replace this
	result := u.db.First(&user, "username = ? OR password = ?", user.Username, user.Password)
	if result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return result.Error
	}
	return nil
}
