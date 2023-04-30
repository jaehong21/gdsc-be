package dto

type CreateProfessorDto struct {
	ID       int    `json:"id" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Name     string `json:"name" validate:"required"`
	Major    string `json:"major" validate:"required"`
	Phone    string `json:"phone" validate:"required,phone"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginProfessorDto struct {
	ID       int    `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
