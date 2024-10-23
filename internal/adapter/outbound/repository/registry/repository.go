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
	dbExecutor any // eksekutor database
}

func NewRepositoryRegistry() registry.RepositoryRegistry {
	var db *sqlx.DB
	repo := &RepositoryRegistry{
		db: db,
		// cache: cache,
	}

	return repo
}

func (r *RepositoryRegistry) GetUserRepository() repository.UserRepository {
	if r.dbExecutor != nil {
		//return repository2.NewUserRepository(r.dbExecutor)
		return repository2.NewUserRepository(r.db)
	}

	return repository2.NewUserRepository(r.db)
}
