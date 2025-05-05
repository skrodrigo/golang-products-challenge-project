package ports

import (
	"github.com/skrodrigo/products/internal/application/entities"
	"github.com/skrodrigo/products/internal/application/shared"
)

type FavoriteProductRepository interface {
	FindByID(id string) (*entities.FavoriteProduct, error)
	FindAll(filter shared.Filter, pageRequest shared.PageRequest) (*shared.Page[entities.FavoriteProduct], error)
	FindByCustomerID(customerID string) (*shared.Page[entities.FavoriteProduct], error)
	Create(favoriteProduct *entities.FavoriteProduct) error
	Update(favoriteProduct *entities.FavoriteProduct) error
	Delete(id string) error
}