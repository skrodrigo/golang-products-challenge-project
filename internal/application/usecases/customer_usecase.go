package usecases

import (
	"github.com/skrodrigo/products/internal/application/dtos"
	"github.com/skrodrigo/products/internal/application/shared"
)

type (
	CustomerUseCase interface {
		Create(paylod *dtos.CreateCustomerDTO) (*dtos.CustomerDTO, error)
		Update(id string, payload *dtos.UpdateCustomerDTO) (*dtos.CustomerDTO, error)
		Delete(id string) error
		FindByID(id string) (*dtos.CustomerDTO, error)
		FindAll(query *dtos.FindAllCustomerQuery, pageRequest shared.PageRequest) (*shared.Page[dtos.CustomerDTO], error)
		AddFavoritesProducts(payload *dtos.AddFavoriteProductsDTO) error
		RemoveFavoriteProduct(payload *dtos.RemoveFavoriteProductDTO) error
		RemoveAllFavoriteProducts(customerID string) error
		FindFavoriteProducts(query dtos.FindAllFavoriteProductQuery, pageRequest shared.PageRequest) (*shared.Page[dtos.CustomerDTO], error)
		ChangeFavoriteProductPriority(payload *dtos.ChangeFavoriteProductPriorityDTO) error
	}
)