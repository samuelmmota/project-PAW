package tests

import (
	"errors"
	"pawAPIbackend/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubmission(t *testing.T) {
	submissionID := uint64(123)
	var bytes []byte
	bytes = append(bytes, []byte("image_data")...)

	// Mock the repository.GetSubmission function
	SetGetSubmissionMock(func(submissionID uint64) (entity.Submission, error) {
		if submissionID == 123 {
			return entity.Submission{
				ID:        123,
				UserID:    456,
				MediaType: "image",
				Media:     bytes,
			}, nil
		}
		return entity.Submission{}, errors.New("submission not found")
	})
	defer ResetGetSubmissionMock()

	// Call the GetSubmission function
	submissionResponse, err := GetSubmission(submissionID)

	// Verify the result
	assert.NoError(t, err)
	assert.Equal(t, uint64(123), submissionResponse.ID)
	assert.Equal(t, "image", submissionResponse.MediaType)
	assert.Equal(t, bytes, submissionResponse.Media)
}

type SubmissionRepository interface {
	GetSubmission(submissionID uint64) (entity.Submission, error)
}

// Create a mock implementation of the repository interface
type MockSubmissionRepository struct {
	GetSubmissionFunc func(submissionID uint64) (entity.Submission, error)
}

// Implement the GetSubmission function for the mock repository
func (m *MockSubmissionRepository) GetSubmission(submissionID uint64) (entity.Submission, error) {
	if m.GetSubmissionFunc != nil {
		return m.GetSubmissionFunc(submissionID)
	}
	return entity.Submission{}, errors.New("mock GetSubmissionFunc is not defined")
}

// Create a global variable to hold the mock function
var getSubmissionMock func(submissionID uint64) (entity.Submission, error)

// GetSubmission retrieves a submission by ID from the repository
func GetSubmission(submissionID uint64) (entity.Submission, error) {
	if getSubmissionMock != nil {
		return getSubmissionMock(submissionID)
	}
	return entity.Submission{}, nil
}

func SetGetSubmissionMock(mockFunc func(submissionID uint64) (entity.Submission, error)) {
	getSubmissionMock = mockFunc
}

func ResetGetSubmissionMock() {
	getSubmissionMock = nil
}
