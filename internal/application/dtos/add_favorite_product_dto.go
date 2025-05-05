package dtos

type AddFavoriteProductDTO struct {
	CustomerID string `json:"customer_id" validate:"required,uuid"`
	ProductID  string `json:"product_id" validate:"required,uuid"`
}  