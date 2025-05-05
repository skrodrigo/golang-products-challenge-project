package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string         `json:"id" gorm:"primaryKey;type:uuid;"`
	Name        string         `json:"name" gorm:"type:varchar(150);"`
	Description string         `json:"description" gorm:"type:varchar(255);"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2);"`
	PictureUrl     string         `json:"picture" gorm:"type:varchar(255);"`
	Active      bool           `json:"active" gorm:"type:boolean;default:true"`
	CreatedAt   time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func NewProduct(name, description, pictureUrl string, price float64) *Product {
	return &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Price:       price,
		PictureUrl:  pictureUrl,
		Active:      true,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		DeletedAt:   gorm.DeletedAt{},
	}
}