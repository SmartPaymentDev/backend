package repository

import (
	"context"
	"smart-payment-dev-be/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req domain.User) error
	GetUser(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, ID int) (domain.User, error)
	CountGetUser(ctx context.Context) (int, error)
	UpdateUserById(ctx context.Context, req domain.User) error
	DeleteUserById(ctx context.Context, id int) error
}
