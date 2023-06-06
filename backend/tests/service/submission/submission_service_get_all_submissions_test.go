package tests

import (
	"pawAPIbackend/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllSubmissionsMock(t *testing.T) {

	var bytes []byte
	bytes = append(bytes, []byte("image_data")...)
	userID := uint64(1)
	expectedResponse := []dto.SubmissionResponseDTO{

		{
			ID:          1,
			MediaType:   "image",
			Date:        "2021-01-01",
			Description: "description",
			BodyPart:    "Feet",
			Media:       bytes,
		},
		{
			ID:          2,
			MediaType:   "video",
			Date:        "2021-01-01",
			Description: "description",
			BodyPart:    "Head",
			Media:       bytes,
		},
	}

	// Mock the repository.GetAllUserSubmissions function
	SetGetAllUserSubmissionsMock(func(userID uint64) []dto.SubmissionResponseDTO {
		return []dto.SubmissionResponseDTO{
			{
				ID:          1,
				MediaType:   "image",
				Date:        "2021-01-01",
				Description: "description",
				BodyPart:    "Feet",
				Media:       bytes,
			},
			{
				ID:          2,
				MediaType:   "video",
				Date:        "2021-01-01",
				Description: "description",
				BodyPart:    "Head",
				Media:       bytes,
			},
		}
	})
	defer ResetGetAllUserSubmissionsMock()

	response := GetAllUserSubmissions(userID)

	// Verify the result
	assert.Equal(t, expectedResponse, response, "Unexpected response from GetAllSubmissions")

	// Additional assertions or checks can be added if needed
}

var getAllUserSubmissionsMock func(userID uint64) []dto.SubmissionResponseDTO

// SetGetAllUserSubmissionsMock sets the mock function for GetAllUserSubmissions.
func SetGetAllUserSubmissionsMock(mockFunc func(userID uint64) []dto.SubmissionResponseDTO) {
	getAllUserSubmissionsMock = mockFunc
}

// ResetGetAllUserSubmissionsMock resets the mock function for GetAllUserSubmissions.
func ResetGetAllUserSubmissionsMock() {
	getAllUserSubmissionsMock = nil
}

// GetAllUserSubmissions retrieves all submissions for a given user.
func GetAllUserSubmissions(userID uint64) []dto.SubmissionResponseDTO {
	if getAllUserSubmissionsMock != nil {
		return getAllUserSubmissionsMock(userID)
	}

	return nil
}
