package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;"`
	Fullname  string    `json:"full_name" gorm:"type:varchar(150);"`
	Email     string    `json:"email" gorm:"type:varchar(150);unique;"`
	UserId    string    `json:"user_id" gorm:"type:uuid;"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}

func NewCustomer(fullname, email, userId string) *Customer {	
	return &Customer{
		ID:        uuid.NewString(),
		Fullname:  fullname,
		Email:     email,
		UserId:    userId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: gorm.DeletedAt{},
	}
}
