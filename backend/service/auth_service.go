package service

import (
	"pawAPIbackend/dto"
	"pawAPIbackend/repository"
)

// devovle string (token) ou erro
func Login(loginDTO dto.LoginDTO) (string, error) {
	token := ""
	UserID, err := repository.Login(loginDTO)
	if err != nil {
		return token, err
	}

	token, _ = CreateToken(UserID)
	return token, nil
}
