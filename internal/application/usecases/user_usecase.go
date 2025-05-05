package usecases

import (
	"errors"

	"github.com/skrodrigo/products/internal/application/dtos"
	"github.com/skrodrigo/products/internal/application/entities"
	"github.com/skrodrigo/products/internal/application/ports"
	"github.com/skrodrigo/products/internal/application/shared"
)

type (
	UserUseCase interface {
		Create(payload *dtos.CreateUserDTO) (*dtos.UserDTO, error)
		Delete(id string) error
		FindByID(id string) (*dtos.UserDTO, error)
		FindAll(query *dtos.FindAllUsersQuery, pageRequest shared.PageRequest) (*shared.Page[dtos.UserDTO], error)
		Activate(id string) (*dtos.UserDTO, error)
		Deactivate(id string) (*dtos.UserDTO, error)
		ChangeProfilePicture(id string, picture string) (*dtos.UserDTO, error)
	}

	userUseCase struct {
		userRepository ports.UserRepository
		hashPort ports.HashPort
	}
)

func (u *userUseCase) Activate(id string) (*dtos.UserDTO, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user.Active {
		return nil, nil
	}

	user.SetActive()

	if err := u.userRepository.Update(user); err != nil {
		return nil, err
	}

	return dtos.MapUserDTO(user), nil
}

func (u *userUseCase) Create(payload *dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	existingUser, err := u.userRepository.FindByUsername(payload.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	} 

	hash, err := u.hashPort.Hash(payload.PlainPwd)
	if err != nil {
		return nil,err 
	}
	user := entities.NewUser(payload.Username, hash)

	if err := u.userRepository.Create(user); err != nil {
		return nil, err
	}

	return dtos.MapUserDTO(user), nil
}

func (u *userUseCase) Deactivate(id string) (*dtos.UserDTO, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user.Active {
		return nil, nil
	}

	user.SetActive()

	if err := u.userRepository.Update(user); err != nil {
		return nil, err
	}

	return dtos.MapUserDTO(user), nil
}

func (u *userUseCase) Delete(id string) error {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return nil
	}

	if err := u.userRepository.Delete(user); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) FindAll(query *dtos.FindAllUsersQuery, pageRequest shared.PageRequest) (*shared.Page[dtos.UserDTO], error) {
	users, err := u.userRepository.FindAll(query, pageRequest)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New("no users found")
	}

	return shared.MapPage[entities.User, dtos.UserDTO](users, shared.MapFunc[entities.User, dtos.UserDTO](func(t *entities.User) dtos.UserDTO {
		return *dtos.MapUserDTO(t)
	})), nil
}

func (u *userUseCase) FindByID(id string) (*dtos.UserDTO, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return dtos.MapUserDTO(user), nil
}

func NewUserUseCase(userRepository ports.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}