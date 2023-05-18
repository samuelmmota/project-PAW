package service

import (
	"errors"
	"log"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"

	"github.com/mashingan/smapping"
)

func GetAllUsers() []dto.UserResponseDTO {
	var responseList []dto.UserResponseDTO

	users := repository.GetAllUsers()

	for _, user := range users {
		response := dto.UserResponseDTO{}
		err := smapping.FillStruct(&response, smapping.MapFields(&user))
		if err != nil {
			log.Fatal("failed to map to response ", err)
			return responseList
		}
		responseList = append(responseList, response)
	}

	return responseList
}

func Register(user entity.User) entity.User {
	user = repository.InsertUser(user)
	return user
}

func Profile(id uint64) (dto.UserProfileResponseDTO, error) {
	userResponse := dto.UserProfileResponseDTO{}

	user, err := repository.GetUser(id)

	if err != nil {
		return userResponse, err
	}

	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return userResponse, err
	}

	return userResponse, err
	//return repository.GetUser(id)
}

func UpdateProfile(userDto dto.UserUpdateDTO) (dto.UserResponseDTO, error) {
	user := entity.User{}
	userResponse := dto.UserResponseDTO{}

	err := smapping.FillStruct(&user, smapping.MapFields(&userDto))
	if err != nil {
		log.Fatal("failed to map ", err)
		return userResponse, nil
	}

	if user, err = repository.UpdateUser(user); err == nil {

		err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))

		if err != nil {
			log.Fatal("failed to map ", err)
			return userResponse, nil
		}

		return userResponse, nil
	}

	return userResponse, errors.New("User do not exists")
}

func DeleteAccount(userID uint64) error {
	if err := repository.DeleteUser(userID); err == nil {
		return nil
	}
	return errors.New("User do not exists")
}

func IsAllowedUser(userID uint64, user_id uint64) bool {
	return userID == user_id
}
