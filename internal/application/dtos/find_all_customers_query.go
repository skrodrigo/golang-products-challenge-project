package dtos

type FindAllCustomerQuery struct {
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
}