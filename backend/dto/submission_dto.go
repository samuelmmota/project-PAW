package dto

type SubmissionCreateDTO struct {
	Description string `json:"description" form:"description" binding:"required"`
	BodyPart    string `json:"body_part" form:"body_part" binding:"required"`
	Date        string `json:"date" form:"date" binding:"required"`
	MediaType   string `json:"media_type" form:"media_type" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type SubmissionResponseDTO struct {
	BodyPart    string `json:"body_part" binding:"required"`
	Media       []byte `json:"media" binding:"required"`
	MediaType   string `json:"media_type" binding:"required"`
	Date        string `json:"date" binding:"required"`
	ID          uint64 `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type SubmissionUpdateDTO struct {
	ID                 uint64 `json:"id" form:"id"`
	Description        string `json:"description" form:"description"`
	BodyPart           string `json:"body_part" form:"body_part"`
	Date               string `json:"date" form:"date"`
	UserID             uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	ShareWithClinicals string `json:"shareWithClinicals" form:"shareWithClinicals"`
}

type SubmissionExportResearcherDTO struct {
	BodyPart  string `json:"body_part" binding:"required"`
	Media     []byte `json:"media" binding:"required"`
	MediaType string `json:"media_type" binding:"required"`
	Date      string `json:"date" binding:"required"`
}
