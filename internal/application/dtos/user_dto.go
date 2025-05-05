package dtos

import (
	"time"

	"github.com/skrodrigo/products/internal/application/entities"
)

type UserDTO struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool   `json:"active"`
}

func MapUserDTO(user *entities.User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		ProfilePicture: user.ProfilePicture,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Active:    user.Active,
	}
}