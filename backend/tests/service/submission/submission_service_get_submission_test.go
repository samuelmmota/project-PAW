package tests

import (
	"pawAPIbackend/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubmissionMock(t *testing.T){
	
	var bytes []byte
	bytes = append(bytes, []byte("image_data")...)
	submissionID := uint64(1)
	expectedResponse := dto.SubmissionResponseDTO{
			ID:          1,
			MediaType:   "image",
			Date:        "2021-01-01",
			Description: "description",
			BodyPart:    "Feet",
			Media:       bytes,
	}

	// Mock the repository.GetSubmission function
	SetGetSubmissionMock(func(submissionID uint64) dto.SubmissionResponseDTO {
		return dto.SubmissionResponseDTO{
				ID:          1,
				MediaType:   "image",
				Date:        "2021-01-01",
				Description: "description",
				BodyPart:    "Feet",
				Media:       bytes,
			
		}
	})
	defer ResetGetSubmissionMock()

	response := GetSubmission(submissionID)

	// Verify the result
	assert.Equal(t, expectedResponse, response, "Unexpected response from GetSubmission")

}
	
	var getSubmissionMock func(submissionID uint64) dto.SubmissionResponseDTO
	
	// SetGetSubmissionMock sets the mock function for GetSubmission.
	func SetGetSubmissionMock(mockFunc func(submissionID uint64) dto.SubmissionResponseDTO) {
		getSubmissionMock = mockFunc
	}
	
	// ResetGetSubmissionMock resets the mock function for GetSubmission.
	func ResetGetSubmissionMock() {
		getSubmissionMock = nil
	}
	
	// GetSubmission retrieves submissions for a given submission id.
	func GetSubmission(submissionID uint64) dto.SubmissionResponseDTO {
		submissionResponse := dto.SubmissionResponseDTO{}
		if getSubmissionMock != nil {
			return getSubmissionMock(submissionID)
		}
	
		return submissionResponse
	}
	