package registry

import (
	"github.com/jmoiron/sqlx"
	repository2 "smart-payment-dev-be/internal/adapter/outbound/repository"
	"smart-payment-dev-be/internal/core/port/outbound/registry"
	"smart-payment-dev-be/internal/core/port/outbound/repository"
)

type DLockConfig struct {
	ProviderUri    string
	AcquireTimeout int
}

type RepositoryRegistry struct {
	db *sqlx.DB
	// cache      *cache.RedisClient
	dbExecutor repository2.DBExecutor // eksekutor database
}

func NewRepositoryRegistry(db *sqlx.DB) registry.RepositoryRegistry {
	repo := &RepositoryRegistry{
		db: db,
		// cache: cache,
	}

	return repo
}

func (r *RepositoryRegistry) GetUserRepository() repository.UserRepository {
	if r.dbExecutor != nil {
		return repository2.NewUserRepository(r.dbExecutor)
	}

	return repository2.NewUserRepository(r.db)
}

func (r *RepositoryRegistry) GetBillRepository() repository.BillRepository {
	if r.dbExecutor != nil {
		return repository2.NewBillRepository(r.dbExecutor)
	}

	return repository2.NewBillRepository(r.db)
}

func (r *RepositoryRegistry) GetVPaymentDetailRepository() repository.VPaymentDetailRepository {
	if r.dbExecutor != nil {
		return repository2.NewVPaymentDetailRepository(r.dbExecutor)
	}

	return repository2.NewVPaymentDetailRepository(r.db)
}

func (r *RepositoryRegistry) GetTransactionRepository() repository.TransactionRepository {
	if r.dbExecutor != nil {
		return repository2.NewTransactionRepository(r.dbExecutor)
	}

	return repository2.NewTransactionRepository(r.db)
}
