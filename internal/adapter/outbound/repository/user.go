package repository

import (
	"context"
	"smart-payment-dev-be/internal/core/domain"
)

type UserRepository struct {
	db DBExecutor
}

func NewUserRepository(db DBExecutor) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, req domain.User) error {
	return nil
}

func (u *UserRepository) GetUserById(ctx context.Context, ID int) (domain.User, error) {
	return domain.User{}, nil
}

func (u *UserRepository) GetUser(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (u *UserRepository) UpdateUserById(ctx context.Context, req domain.User) error {
	return nil
}

func (u *UserRepository) DeleteUserById(ctx context.Context, id int) error {
	return nil
}

func (u *UserRepository) CountGetUser(ctx context.Context) (int, error) {
	return 0, nil
}
