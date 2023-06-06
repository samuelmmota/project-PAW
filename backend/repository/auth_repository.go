package repository

import (
	"pawAPIbackend/config"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"

	"golang.org/x/crypto/bcrypt"
)

func Login(loginDTO dto.LoginDTO) (uint64, error) {
	var user entity.User
	err := config.Db.Where("email = ?", loginDTO.Email).First(&user).Error
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
