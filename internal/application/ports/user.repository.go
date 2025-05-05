package ports

import (
	"github.com/skrodrigo/products/internal/application/dtos"
	"github.com/skrodrigo/products/internal/application/entities"
	"github.com/skrodrigo/products/internal/application/shared"
)

type UserRepository interface {
	FindByID(id string) (*entities.User, error)
	FindAll(query *dtos.FindAllUsersQuery, PageRequest shared.PageRequest) (*shared.Page[entities.User], error)
	FindByUsername(username string) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(user *entities.User) error
}
