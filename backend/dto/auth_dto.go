package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password,omitempty" validate:"min:6" binding:"required"`
}

type EvaluateTokenDTO struct {
	Token string `json:"token" form:"token" binding:"required"`
}
