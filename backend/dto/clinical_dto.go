package dto

// for Patients
type ClinicalRequestDTO struct {
	ClinicalEmail string `json:"clinical_email" form:"clinical_email" binding:"required" validate:"email"`
}

type ClinicalCreateDTO struct {
	ClinicalEmail string `json:"clinical_email" form:"clinical_email" binding:"required" validate:"email"`
}

type ClinicalResponseDTO struct {
	ClinicalEmail string `json:"clinical_email" form:"clinical_email"`
}

//for Clinicals

//type PatientRequestDTO struct {

type PatientSubmissionResponseDTO struct {
	//Submission []SubmissionResponseDTO `json:"submission" form:"submission"`
}

type PatientResponseDTO struct {
	Submission   []SubmissionResponseDTO `json:"submission" form:"submission"`
	PatientEmail string                  `json:"patient_email" form:"patient_email"`
}

type PatientSubmitFeedbackDTO struct {
	SubmissionID uint64 `json:"submission_id" form:"submission_id" binding:"required"`
	Date         string `json:"date" form:"date" binding:"required"`
	Feedback     string `json:"feedback" form:"feedback" binding:"required"`
}
