package entities

import (
	"time"

	"github.com/google/uuid"
)

type FavoriteProduct struct {
	ID         string     `json:"id" gorm:"primaryKey;type:uuid;"`
	ProductID  string     `json:"product_id" gorm:"type:uuid;index;"`
	CustomerID string     `json:"customer_id" gorm:"type:uuid;index;"`
	Priority   int        `json:"priority" gorm:"type:int;"`
	FavoriteAt time.Time     `json:"favorite_at" gorm:"type:timestamp;"`
	RemovedAt  *time.Time `json:"removed_at" gorm:"type:timestamp;"`
}

func NewFavoriteProduct(productID, customerID string, priority int) *FavoriteProduct {	
	return &FavoriteProduct{
		ID:         uuid.NewString(),
		ProductID:  productID,
		CustomerID: customerID,
		Priority:   priority,
		FavoriteAt: time.Now().UTC(),
		RemovedAt: nil,
	}
}
