package repository

import (
	"pawAPIbackend/config"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
)

// devovle o id do utilizador
func Login(loginDTO dto.LoginDTO) (uint64, error) {
	var user entity.User
	err := config.Db.Where("email = ? AND password = ?", loginDTO.Email, loginDTO.Password).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
