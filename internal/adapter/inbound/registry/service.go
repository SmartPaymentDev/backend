package registry

import (
	iservice "smart-payment-dev-be/internal/core/port/inbound/service"
	"smart-payment-dev-be/internal/core/port/outbound/registry"
	"smart-payment-dev-be/internal/service/bill"
	"smart-payment-dev-be/internal/service/transaction"
	"smart-payment-dev-be/internal/service/user"
)

type ServiceRegistry struct {
	repositoryRegistry registry.RepositoryRegistry
	userService        iservice.UserService
	billService        iservice.BillService
	transactonService  iservice.TransactionService
}

func NewServiceRegistry(repositoryRegistry registry.RepositoryRegistry) *ServiceRegistry {
	return &ServiceRegistry{
		repositoryRegistry: repositoryRegistry,
		userService:        user.NewUserService(repositoryRegistry),
		billService:        bill.NewBillService(repositoryRegistry),
		transactonService:  transaction.NewTransactionService(repositoryRegistry),
	}
}

func (r *ServiceRegistry) GetUserService() iservice.UserService {
	return r.userService
}

func (r *ServiceRegistry) GetBillService() iservice.BillService {
	return r.billService
}

func (r *ServiceRegistry) GetTransactionService() iservice.TransactionService {
	return r.transactonService
}
