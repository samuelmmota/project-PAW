package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"email"`
	Password string `json:"password" form:"password,omitempty"`
	//Clinicals []EvaluateResponseDTO `json:"clinicals" form:"clinicals"`
	//Patients  []EvaluateResponseDTO `json:"patients" form:"patients"`
}

type RegisterDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password,omitempty" validate:"min:6" binding:"required"`
}

type UserResponseDTO struct {
	ID    uint64 `json:"id" form:"id"`
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
	//Clinicals []EvaluateResponseDTO `json:"clinicals" form:"clinicals"`
	//Patients  []EvaluateResponseDTO `json:"patients" form:"patients"`
}

type UserProfileResponseDTO struct {
	ID    uint64 `json:"id" form:"id"`
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
	//Clinicals []EvaluateResponseDTO `json:"clinicals" form:"clinicals"`
	//Patients  []EvaluateResponseDTO `json:"patients" form:"patients"`
}

type EvaluateResponseDTO struct {
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
}
