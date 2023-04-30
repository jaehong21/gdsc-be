package dto

type CreateStudentDto struct {
	ID       int    `json:"id" validate:"required,gte=20000000,lte=30000000"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Name     string `json:"name" validate:"required"`
	Major    string `json:"major" validate:"required"`
	Phone    string `json:"phone" validate:"required,phone"`
	Email    string `json:"email" validate:"required,email"`
	DeviceID string `json:"device_id" validate:"required"`
}

type LoginStudentDto struct {
	ID       int    `json:"id" validate:"required,gte=20000000,lte=30000000"`
	Password string `json:"password" validate:"required"`
}
