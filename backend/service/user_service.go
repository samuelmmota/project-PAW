package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"
	"strconv"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
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

func Register(user entity.User) (entity.User, error) {
	err := EncryptPassword(&user)
	key, err := GenerateRandomString(20)
	if err != nil {
		log.Fatal("Failed to generate key:", err)
		return entity.User{}, err
	}
	user.Key = key
	user, err = repository.InsertUser(user)

	return user, err
}

func CheckEmail(email string) (entity.User, error) {
	return repository.CheckEmail(email)
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
	user, err := repository.GetUser(userDto.ID)
	if err != nil {
		return dto.UserResponseDTO{}, errors.New("User does not exist")
	}
	if userDto.Email != "" {
		user.Email = userDto.Email
	}

	if userDto.Password != "" {
		user.Password = userDto.Password
		//encript there
	}

	if userDto.IsClinical != "" {
		isClinical, err := strconv.ParseBool(userDto.IsClinical)
		if err != nil {
			return dto.UserResponseDTO{}, errors.New("Invalid value for IsClinical")
		}
		user.IsClinical = isClinical
	}

	if userDto.ExportToReasearcher != "" {
		exportToReasearcher, err := strconv.ParseBool(userDto.ExportToReasearcher)
		if err != nil {
			return dto.UserResponseDTO{}, errors.New("Invalid value for ExportToReasearcher")
		}
		user.ExportToReasearcher = exportToReasearcher
	}

	userResponse := dto.UserResponseDTO{}

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

func EncryptPassword(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func ComparePassword(user *entity.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func GenerateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)

	randomString = randomString[:length]

	return randomString, nil
}
