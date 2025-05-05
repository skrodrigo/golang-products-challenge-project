package ports

import (
	"github.com/skrodrigo/products/internal/application/entities"
	"github.com/skrodrigo/products/internal/application/shared"
)

type CustomerRepository interface {
	FindByID(id string) (*entities.Customer, error)
	FindByEmail(email string) (*entities.Customer, error)
	FindAll(Filter shared.Filter, PageRequest shared.PageRequest) (*shared.Page[entities.Customer], error)
	Create(customer *entities.Customer) error
	Update(customer *entities.Customer) error
	Delete(id string) error
}