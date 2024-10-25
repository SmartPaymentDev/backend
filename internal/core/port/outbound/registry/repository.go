package registry

import "smart-payment-dev-be/internal/core/port/outbound/repository"

type RepositoryRegistry interface {
	GetUserRepository() repository.UserRepository
	GetBillRepository() repository.BillRepository
	GetVPaymentDetailRepository() repository.VPaymentDetailRepository
	GetTransactionRepository() repository.TransactionRepository
}
