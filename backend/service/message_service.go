package service

import (
	"errors"
	"log"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"
)

func GetMessages(submissionID uint64, userID uint64) ([]dto.MessageResponseDTO, error) {
	var messagesResponse []dto.MessageResponseDTO

	var submission, err = repository.GetSubmission(submissionID)

	if err != nil {
		return messagesResponse, errors.New("Submission not found")
	}
	user, err := repository.GetUser(userID)

	if user.IsClinical == false {
		if submission.UserID != userID {
			return messagesResponse, errors.New("User is not the owner of the submission")
		}
	} else {
		patientID := submission.UserID

		_, err := repository.GetCLinical(userID, patientID)

		if err != nil {
			return messagesResponse, errors.New("Clinical is not allowed to see the messages")
		}

	}

	messages := repository.GetMessages(submissionID)

	if messages == nil {
		return messagesResponse, errors.New("No messages found")
	}

	for _, message := range messages {
		response := dto.MessageResponseDTO{}

		clinical, _ := repository.GetUser(message.UserID)
		response.ClinicalEmail = clinical.Email

		response.MessageContent = message.MessageContent
		response.Date = message.Date
		response.SubmissionID = message.SubmissionID
		messagesResponse = append(messagesResponse, response)
	}

	return messagesResponse, nil

}

func AddMessage(messageCreate dto.MessageCreateDTO) (dto.MessageResponseDTO, error) {
	var messageResponse dto.MessageResponseDTO
	log.Default().Println("AddMessage => submissionID= ", messageCreate.SubmissionID, "&&", "ClinicalID= ", messageCreate.ClinicalID)

	var submission, err = repository.GetSubmission(messageCreate.SubmissionID)

	if err != nil {
		return messageResponse, errors.New("Submission not found")
	}
	user, err := repository.GetUser(messageCreate.ClinicalID)

	if user.IsClinical == false {
		return messageResponse, errors.New("User is cannot add a message to this submission since its not a clinical")
	} else {
		patientID := submission.UserID
		_, err := repository.GetCLinical(messageCreate.ClinicalID, patientID)

		if err != nil {
			return messageResponse, errors.New("Clinical is not allowed to add the messages on this submission")
		}
	}

	message := entity.Message{
		MessageContent: messageCreate.MessageContent,
		SubmissionID:   messageCreate.SubmissionID,
		UserID:         messageCreate.ClinicalID,
		Date:           messageCreate.Date,
	}

	err = repository.AddMessage(message)

	if err != nil {
		return messageResponse, errors.New("Error adding the message")
	}

	messageResponse.ClinicalEmail = message.User.Email
	messageResponse.MessageContent = message.MessageContent
	messageResponse.Date = message.Date
	messageResponse.SubmissionID = message.SubmissionID

	return messageResponse, nil
}
