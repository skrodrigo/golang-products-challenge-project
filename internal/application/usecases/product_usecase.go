package usecases

import (
	"github.com/skrodrigo/products/internal/application/dtos"
	"github.com/skrodrigo/products/internal/application/shared"
)

type (
	ProductUseCase interface {
		Create(payload *dtos.CreateProductDTO) (*dtos.ProductDTO, error)
		Update(id string, payload *dtos.UpdateProductDTO) (*dtos.ProductDTO, error)
		Delete(id string) error
		FindByID(id string) (*dtos.ProductDTO, error)
		FindAll(query dtos.FindAllProductsQuery, pageRequest shared.PageRequest) (*shared.Page[dtos.UserDTO], error)
		Active(id string) error
		Desactive(id string) error
	}
)