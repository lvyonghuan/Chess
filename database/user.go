package database

import (
	"Chess/model"
	"errors"
	"gorm.io/gorm"
)

func FindUserByUsername(username string) (err error) {
	var user model.User
	err = DB.Where("name=?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return nil
}

func Register(user model.User) (err error) {
	err = DB.Create(&user).Error
	return err
}

func Login(username string) (user model.User, err error) {
	err = DB.Where("name=?", username).First(&user).Error
	return user, err
}
