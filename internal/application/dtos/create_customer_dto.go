package dtos

type CreateCustomerDTO struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	UserId   string `json:"user_id" validate:"required"`
}