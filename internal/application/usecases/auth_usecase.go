package usecases

import (
	"github.com/skrodrigo/products/internal/application/dtos"
	"github.com/skrodrigo/products/internal/application/ports"
)

type (
	AuthUseCase interface {
		Authenticate(payload *dtos.AuthenticateDTO) (*ports.AuthToken, error)
	}
)