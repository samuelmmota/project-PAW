package repository

import (
	"errors"
	"pawAPIbackend/config"
	"pawAPIbackend/entity"
)

func InsertUser(user entity.User) (entity.User, error) {
	err := config.Db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func CheckEmail(email string) (entity.User, error) {
	var user entity.User
	err := config.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, errors.New("User Not Found")
	}

	return user, nil
}

func GetAllUsers() []entity.User {
	var users []entity.User
	config.Db.Preload("User").Find(&users)

	return users
}

func GetUser(userID uint64) (entity.User, error) {
	var user entity.User
	config.Db.Preload("User").First(&user, userID)
	if user.ID != 0 {
		return user, nil
	}

	return user, errors.New("User do not exists")
}

func UpdateUser(user entity.User) (entity.User, error) {
	if _, err := GetUser(user.ID); err == nil {
		config.Db.Save(&user)
		config.Db.Preload("User").Find(&user)
		return user, nil
	}
	return user, errors.New("User do not exists")
}

func DeleteUser(userID uint64) error {
	var user entity.User
	config.Db.First(&user, userID)
	if user.ID != 0 {
		config.Db.Delete(&user)
		return nil
	}
	return errors.New("User do not exists")
}
