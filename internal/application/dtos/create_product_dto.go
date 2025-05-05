package dtos

type CreateProductDTO struct {
	Name 			string  `json:"name" validate:"required,min=3,max=100"`
	Description 	string  `json:"description" validate:"required,min=3,max=100"`
	Price 			float64 `json:"price" validate:"required,min=0"`
}