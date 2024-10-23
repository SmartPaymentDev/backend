package user

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/port/outbound/registry"
)

type UserService struct {
	repositoryRegistry registry.RepositoryRegistry
}

func NewUserService(repositoryRegistry registry.RepositoryRegistry) *UserService {
	return &UserService{
		repositoryRegistry: repositoryRegistry,
	}
}

func (b *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	var userDTO dto.UserDTO

	repo := b.repositoryRegistry.GetUserRepository()

	//dto => model
	data := userDTO.CreateUserTransformIn(req)

	err := repo.CreateUser(ctx, data)

	if err != nil {
		return err
	}

	return nil
}

func (b *UserService) GetUser(ctx context.Context) (dto.UsersResponse, error) {
	var userDTO dto.UserDTO

	repo := b.repositoryRegistry.GetUserRepository()

	user, err := repo.GetUser(ctx)
	totalData, _ := repo.CountGetUser(ctx)

	res := userDTO.GetUsersTransformOut(user, totalData)

	return res, err
}

func (b *UserService) GetUserById(ctx context.Context, id int) (dto.UserResponse, error) {
	var userDTO dto.UserDTO
	repo := b.repositoryRegistry.GetUserRepository()
	user, err := repo.GetUserById(ctx, id)

	if err != nil {
		return dto.UserResponse{}, err
	}

	res := userDTO.GetUserTransformOut(user)
	return res, nil
}

func (b *UserService) UpdateUserById(ctx context.Context, params dto.UpdateUserRequest) error {
	var userDTO dto.UserDTO

	repo := b.repositoryRegistry.GetUserRepository()

	if _, err := repo.GetUserById(ctx, params.Id); err != nil {
		return err
	}

	req := userDTO.UpdateUserTransformIn(params)
	err := repo.UpdateUserById(ctx, req)

	if err != nil {
		return err
	}
	return nil
}

func (b *UserService) DeleteUserById(ctx context.Context, id int) error {
	repo := b.repositoryRegistry.GetUserRepository()
	if _, err := repo.GetUserById(ctx, id); err != nil {
		return err
	}
	err := repo.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
