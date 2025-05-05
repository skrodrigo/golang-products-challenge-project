package dtos

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	PlainPwd string `json:"password" validate:"required,min=6,max=20"`
}