package service

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
)

type UserService interface {
	LoginUser(ctx context.Context, req dto.LoginUserRequest) (interface{}, error)
	GetUserByMe(ctx context.Context, token string) (interface{}, error)
	ChangePassword(ctx context.Context, req dto.ChangePasswordRequest) error
}
