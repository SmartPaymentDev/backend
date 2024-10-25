package user

import (
	"context"
	"errors"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/port/outbound/registry"
	"smart-payment-dev-be/shared/util"
)

type UserService struct {
	repositoryRegistry registry.RepositoryRegistry
}

func NewUserService(repositoryRegistry registry.RepositoryRegistry) *UserService {
	return &UserService{
		repositoryRegistry: repositoryRegistry,
	}
}

func (b *UserService) LoginUser(ctx context.Context, req dto.LoginUserRequest) (interface{}, error) {
	var userDTO dto.UserDTO
	var resp dto.UserResponse

	repo := b.repositoryRegistry.GetUserRepository()
	data := userDTO.LoginUserTransformIn(req)

	user, err := repo.GetUserByNoCust(ctx, data.NOCUST)
	if err != nil {
		return nil, errors.New("NO CUST NOT FOUND")
	}
	validPass := util.EncryptPassword(req.MOBILEPASSWORD, user.MOBILEPASSWORD)
	if !validPass {
		return nil, errors.New("PASSWORD NOT MATCH")
	}

	token, err := util.CreateJwtToken(user.CUSTID, user.NOCUST, user.NMCUST, user.NO_WA)
	if err != nil {
		return nil, err
	}

	resp.JwtToken = token
	return resp, nil
}

func (b *UserService) GetUserByMe(ctx context.Context, token string) (interface{}, error) {
	var userDTO dto.UserDTO

	repo := b.repositoryRegistry.GetUserRepository()

	custId, err := util.ParseJwtToken(token)

	if err != nil {
		return nil, err
	}

	user, err := repo.GetUserByCustId(ctx, custId)
	if err != nil {
		return nil, err
	}

	resp := userDTO.GetUserByMe(user)

	return resp, nil
}
