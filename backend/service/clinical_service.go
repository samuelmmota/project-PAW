package service

import (
	"errors"
	"log"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"

	"github.com/mashingan/smapping"
)

func GetClinicals(userID uint64) []dto.ClinicalResponseDTO {
	var clinicalResponse []dto.ClinicalResponseDTO
	clinicals := repository.GetAllUserClinicals(userID)

	for _, clinical := range clinicals {
		response := dto.ClinicalResponseDTO{}
		err := smapping.FillStruct(&response, smapping.MapFields(&clinical))
		if err != nil {
			log.Fatal("failed to map clinical to response ", err)
			return clinicalResponse
		}
		clinicalResponse = append(clinicalResponse, response)
	}

	return clinicalResponse
}

func AddUserClinical(patientID uint64, clinicalCreateDTO dto.ClinicalCreateDTO) (dto.ClinicalRequestDTO, error) {
	var clinicalResponseDTO dto.ClinicalRequestDTO

	// Check if the patient exists
	patientUser, err := repository.GetUser(patientID)
	if err != nil {
		return clinicalResponseDTO, errors.New("patient not found")
	}

	clinicalUser, err := repository.CheckEmail(clinicalCreateDTO.ClinicalEmail)

	if err != nil {
		return clinicalResponseDTO, errors.New("clinical not found")
	}

	if clinicalUser.IsClinical == false {
		return clinicalResponseDTO, errors.New("User is not a clinical")
	}

	clinical := entity.Clinical{
		ClinicalID: clinicalUser.ID,
		Clinical:   clinicalUser,
		PatientID:  patientUser.ID,
		Patient:    patientUser,
	}

	// Check if the clinical already exists for the patient
	if repository.DoesClinicalExist(clinical) {
		return clinicalResponseDTO, errors.New("clinical already exists for patient")
	}

	err = repository.CreateClinical(clinical)
	if err != nil {
		return clinicalResponseDTO, err
	}

	clinicalResponseDTO.ClinicalEmail = clinical.Clinical.Email

	return clinicalResponseDTO, err
}

func GetClinicalSubmissions(clinicalID uint64) ([]dto.ClinicalSubmissionResponseDTO, error) {
	var clinicalSubmissionsResponse []dto.ClinicalSubmissionResponseDTO

	clinicals := repository.GetAllUserClinicals(clinicalID)

	for _, clinical := range clinicals {
		response := dto.ClinicalSubmissionResponseDTO{}
		response.Submission = GetAllSubmissions(clinical.PatientID)
		response.Email = clinical.Patient.Email
		clinicalSubmissionsResponse = append(clinicalSubmissionsResponse, response)
	}

	/*if err != nil {
		return clinicalSubmissionsResponse, nil
	}*/

	return clinicalSubmissionsResponse, nil
}
