package service

import (
	"errors"
	"log"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"

	"github.com/mashingan/smapping"
)

func GetAllSubmissions() []dto.SubmissionResponseDTO {
	var submissionResponse []dto.SubmissionResponseDTO
	submissions := repository.GetAllSubmissions()

	for _, user := range submissions {
		response := dto.SubmissionResponseDTO{}
		err := smapping.FillStruct(&response, smapping.MapFields(&user))
		if err != nil {
			log.Fatal("failed to map submission to response ", err)
			return submissionResponse
		}
		submissionResponse = append(submissionResponse, response)
	}

	return submissionResponse
}

func InsertSubmission(submissionCreateDTO dto.SubmissionCreateDTO, userID uint64) dto.SubmissionResponseDTO {
	submission := entity.Submission{}
	submissionResponse := dto.SubmissionResponseDTO{}

	err := smapping.FillStruct(&submission, smapping.MapFields(&submissionCreateDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return submissionResponse
	}

	submission.UserID = userID
	submission = repository.InsertSubmission(submission)

	err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return submissionResponse
	}

	return submissionResponse
}

func GetSubmission(submissionID uint64) (dto.SubmissionResponseDTO, error) {
	submissionResponse := dto.SubmissionResponseDTO{}

	if submission, err := repository.GetSubmission(submissionID); err == nil {

		err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))

		if err != nil {
			log.Fatal("failed to map to response ", err)
			return submissionResponse, err
		}

		return submissionResponse, nil
	}
	return submissionResponse, errors.New("submission do not exist")
}

func UpdateSubmission(submissionDTO dto.SubmissionUpdateDTO) (dto.SubmissionResponseDTO, error) {
	submission := entity.Submission{}
	submissionResponse := dto.SubmissionResponseDTO{}

	err := smapping.FillStruct(&submission, smapping.MapFields(&submissionDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return submissionResponse, nil
	}

	if submission, err = repository.UpdateSubmission(submission); err == nil {

		err = smapping.FillStruct(&submissionResponse, smapping.MapFields(&submission))

		if err != nil {
			log.Fatal("failed to map to response ", err)
			return submissionResponse, err
		}

		return submissionResponse, nil
	}

	return submissionResponse, errors.New("submission do not exists")
}

func DeleteSubmission(submissionID uint64) error {
	if err := repository.DeleteSubmission(submissionID); err == nil {
		return nil
	}
	return errors.New("submission do not exists")
}

func IsAllowedToEdit(userID uint64, submissionID uint64) bool {
	submission := repository.GetTheSubmissionUsingID(submissionID)
	return userID == submission.UserID
}
