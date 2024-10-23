package registry

import (
	iservice "smart-payment-dev-be/internal/core/port/inbound/service"
	"smart-payment-dev-be/internal/core/port/outbound/registry"
	"smart-payment-dev-be/internal/service/user"
)

type ServiceRegistry struct {
	repositoryRegistry registry.RepositoryRegistry
	userService        iservice.UserService
}

func NewServiceRegistry(repositoryRegistry registry.RepositoryRegistry) *ServiceRegistry {
	return &ServiceRegistry{
		repositoryRegistry: repositoryRegistry,
		userService:        user.NewUserService(repositoryRegistry),
	}
}

func (r *ServiceRegistry) GetUserService() iservice.UserService {
	return r.userService
}
