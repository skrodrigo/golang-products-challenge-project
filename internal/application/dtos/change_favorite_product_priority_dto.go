package dtos

type ChangeFavoriteProductPriorityDTO struct {
	CustomerID string `json:"customer_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Priority   int    `json:"priority" validate:"required"`
}