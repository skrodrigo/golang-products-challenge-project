package dtos

import "github.com/skrodrigo/products/internal/application/entities"

type FavoriteProductDTO struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	FavoriteAt string `json:"favorite_at"`
}

func MapFavoriteProductToDTO(favoriteProduct *entities.FavoriteProduct) *FavoriteProductDTO {
	return &FavoriteProductDTO{
		ID:         favoriteProduct.ID,
		CustomerID: favoriteProduct.CustomerID,
		ProductID:  favoriteProduct.ProductID,
		FavoriteAt: favoriteProduct.FavoriteAt,
	}
}