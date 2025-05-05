package dtos

type ProductDTO struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Price       string `json:"price" validate:"required,min=0"`
	Active			bool   `json:"active"`
}