package tests

import (
	"pawAPIbackend/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	// Prepare mock data for repository.GetAllUsers
	userID := uint64(1)
	mockUsers := []dto.UserResponseDTO{
		{
			ID:    1,
			Email: "jonhdoe@mail.com",
		},
		{
			ID:    2,
			Email: "janesmith@mail.com",
		},
	}

	// Mock the repository.GetAllUsers function
	SetGetAllUsersMock(func(userID uint64) []dto.UserResponseDTO {
		return []dto.UserResponseDTO{
			{
				ID:    1,
				Email: "jonhdoe@mail.com",
			},
			{
				ID:    2,
				Email: "janesmith@mail.com",
			},
		}
	})
	defer ResetGetAllUsersMock()

	// Call the GetAllUsers function
	users := GetAllUsers(userID)

	// Assert the expected number of users
	assert.Len(t, users, len(mockUsers), "Unexpected number of users")

	// Assert the user data
	for i, user := range users {
		assert.Equal(t, mockUsers[i].ID, user.ID, "Unexpected user ID")
		assert.Equal(t, mockUsers[i].Email, user.Email, "Unexpected user name")
	}
}

var getGetAllUsersMock func(userID uint64) []dto.UserResponseDTO

// Mock the repository.GetAllUsers function
func SetGetAllUsersMock(mockFunc func(userID uint64) []dto.UserResponseDTO) {
	getGetAllUsersMock = mockFunc
}

// Reset the mock for GetAllUsers function
func ResetGetAllUsersMock() {
	getGetAllUsersMock = nil
}

// GetAllUserSubmissions retrieves all submissions for a given user.
func GetAllUsers(userID uint64) []dto.UserResponseDTO {
	if getGetAllUsersMock != nil {
		return getGetAllUsersMock(userID)
	}

	return nil
}
