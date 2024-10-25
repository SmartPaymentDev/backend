package repository

import (
	"context"
	"smart-payment-dev-be/internal/core/domain"
)

type UserRepository interface {
	GetUserByNoCust(ctx context.Context, nocust string) (domain.User, error)
	GetUserByCustId(ctx context.Context, custId string) (domain.User, error)
	UpdatePassword(ctx context.Context, req domain.User) error
}
