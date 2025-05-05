package dtos

type AuthenticateDTO struct {
	Username string `json:"username" validate:"required"`
	PlainPwd string `json:"password" validate:"required"`
}