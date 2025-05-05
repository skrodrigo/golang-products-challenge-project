package dtos

import (
	"strings"

	"github.com/skrodrigo/products/internal/application/entities"
)

type CustomerDTO struct {
	ID       string `json:"id" gorm:"primaryKey;type:uuid;"`
	FullName string `json:"full_name" gorm:"type:varchar(150);"`
	Email    string `json:"email" gorm:"type:varchar(150);"`
	FirsName string `json:"first_name" gorm:"type:varchar(150);"`
	LastName string `json:"last_name" gorm:"type:varchar(150);"`
}

func MapCustomerToDTO(customer *entities.Customer) *CustomerDTO {
	parts := strings.Split(customer.Fullname, " ")
	
	return &CustomerDTO{
		ID:       customer.ID,
		FullName: customer.Fullname, 
		Email:    customer.Email,
		FirsName: parts[0],
		LastName: parts[len(parts)-1],
	}
}