package dto

type UserUpdateDTO struct {
	ID                  uint64 `json:"id" form:"id"`
	Email               string `json:"email" form:"email"`
	Password            string `json:"password" form:"password,omitempty"`
	ExportToReasearcher string `json:"exportToReasearcher" form:"exportToReasearcher"`
	IsClinical          string `json:"isClinical" form:"isClinical"`
}

type RegisterDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password,omitempty" validate:"min:6" binding:"required"`
}

type UserResponseDTO struct {
	ID                  uint64 `json:"id" form:"id"`
	Email               string `json:"email" form:"email" binding:"required" validate:"email"`
	IsClinical          bool   `json:"isClinical" form:"isClinical"`
	ExportToReasearcher bool   `json:"exportToReasearcher" form:"exportToReasearcher"`
}

type UserProfileResponseDTO struct {
	ID                  uint64 `json:"id" form:"id"`
	Email               string `json:"email" form:"email" binding:"required" validate:"email"`
	IsClinical          bool   `json:"isClinical" form:"isClinical"`
	ExportToReasearcher bool   `json:"exportToReasearcher" form:"exportToReasearcher"`
}

type EvaluateResponseDTO struct {
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
}
