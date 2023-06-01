package dto

type ClinicalRequestDTO struct {
	ClinicalEmail string `json:"clinical_email" form:"clinical_email" binding:"required" validate:"email"`
}

type ClinicalResponseDTO struct {
	Clinical UserResponseDTO `json:"clinical" form:"clinical"`
	Patient  UserResponseDTO `json:"patient" form:"patient"`
}

type ClinicalCreateDTO struct {
	ClinicalEmail string `json:"clinical_email" form:"clinical_email" binding:"required" validate:"email"`
}
