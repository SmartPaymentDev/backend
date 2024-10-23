package registry

import "smart-payment-dev-be/internal/core/port/inbound/service"

type ServiceRegistry interface {
	GetUserService() service.UserService
}
