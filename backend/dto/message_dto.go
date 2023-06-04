package dto

type MessageRequestDTO struct {
	SubmissionID uint64 `json:"submission_id" form:"submission_id" binding:"required"`
	UserID       uint64 `json:"user_id" form:"user_id" binding:"required"`
}

type MessageCreateDTO struct {
	Date           string `json:"date" form:"date" binding:"required"`
	MessageContent string `json:"message_content" form:"message_content" binding:"required"`
	SubmissionID   uint64 `json:"submission_id" form:"submission_id" binding:"required"`
	ClinicalID     uint64 `json:"clinical_id" form:"clinical_id" binding:"required"`
}

type MessageResponseDTO struct {
	Date           string `json:"date" form:"date"`
	MessageContent string `json:"message_content" form:"message_content"`
	SubmissionID   uint64 `json:"submission_id" form:"submission_id"`
	ClinicalEmail  string `json:"clinical_email" form:"clinical_email"`
}
