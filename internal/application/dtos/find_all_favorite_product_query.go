package dtos

import "time"

type FindAllFavoriteProductQuery struct {
	CustomerId       string    `json:"customer_id" validate:"uuid"`
	ProductId        string    `json:"product_id" validate:"uuid"`
	FavoritedAtStart *time.Time `json:"favorited_at_start" validate:"datetime"`
	FavoriteAtEnd    *time.Time    `json:"favorited_at_end" validate:"datetime"`
}