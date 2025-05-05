package dtos

type RemoveFavoriteProductDTO struct {
	CustomerId	string `json:"customer_id"`
	ProductId	string `json:"product_id"`
}