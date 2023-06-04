package service

import (
	"errors"
	"pawAPIbackend/dto"
	"pawAPIbackend/entity"
	"pawAPIbackend/repository"
)

func GetClinicals(userID uint64) ([]dto.ClinicalResponseDTO, error) {
	var clinicalEmailResponse []dto.ClinicalResponseDTO

	patientUser, _ := repository.GetUser(userID)
	if patientUser.IsClinical != false {
		return clinicalEmailResponse, errors.New("User is not a patient")
	}

	clinicals := repository.GetClinicals(userID)

	if clinicals == nil {
		return clinicalEmailResponse, errors.New("no clinicals found")
	}

	for _, clinical := range clinicals {
		response := dto.ClinicalResponseDTO{}
		response.ClinicalEmail = clinical.Clinical.Email
		clinicalEmailResponse = append(clinicalEmailResponse, response)
	}

	return clinicalEmailResponse, nil

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

func GetPatientsSubmissions(clinicalID uint64) ([]dto.PatientResponseDTO, error) {
	var clinicalSubmissionsResponse []dto.PatientResponseDTO

	clinicalUser, nil := repository.GetUser(clinicalID)
	if clinicalUser.IsClinical == false {
		return clinicalSubmissionsResponse, errors.New("User is not a clinical")
	}

	clinicals := repository.GetPatients(clinicalID)

	for _, clinical := range clinicals {
		response := dto.PatientResponseDTO{}
		response.Submission = GetAllSubmissions(clinical.PatientID)

		if len(response.Submission) == 0 {
			response.Submission = []dto.SubmissionResponseDTO{}
		}

		response.PatientEmail = clinical.Patient.Email
		clinicalSubmissionsResponse = append(clinicalSubmissionsResponse, response)
	}

	return clinicalSubmissionsResponse, nil
}
