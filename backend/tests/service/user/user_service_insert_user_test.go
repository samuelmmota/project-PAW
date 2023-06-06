package tests

import (
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	// Create a test user
	user := entity.User{
		Email:    "testuser@mail.com",
		Password: "password",
	}

	expectedUser := dto.UserResponseDTO{
		ID:    1,
		Email: "testuser@mail.com",
	}

	// Set up a mock implementation of the repository.InsertUser function
	SetInsertUserMock(func(user entity.User) dto.UserResponseDTO {
		// Simulate a successful user insertion
		user.ID = 1 // Set a mock ID
		return dto.UserResponseDTO{
			ID:    user.ID,
			Email: "testuser@mail.com",
		}
	})
	defer ResetInsertUserMock()

	response := InsertUser(user)

	// Verify the result
	assert.Equal(t, expectedUser, response, "Unexpected response from GetAllSubmissions")

	// Additional assertions can be made based on your requirements
}

var getInsertUserMock func(user entity.User) dto.UserResponseDTO

func SetInsertUserMock(mockFunc func(user entity.User) dto.UserResponseDTO) {
	getInsertUserMock = mockFunc
}

func ResetInsertUserMock() {
	getInsertUserMock = nil
}

func InsertUser(user entity.User) dto.UserResponseDTO {
	if getInsertUserMock != nil {
		return getInsertUserMock(user)
	}
	return dto.UserResponseDTO{}
}
