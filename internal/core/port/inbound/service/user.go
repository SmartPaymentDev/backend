package service

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) error
	GetUser(ctx context.Context) (dto.UsersResponse, error)
	GetUserById(ctx context.Context, id int) (dto.UserResponse, error)
	UpdateUserById(ctx context.Context, params dto.UpdateUserRequest) error
	DeleteUserById(ctx context.Context, id int) error
}
