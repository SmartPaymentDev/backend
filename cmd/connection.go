package cmd

import (
	"github.com/jmoiron/sqlx"
	"smart-payment-dev-be/internal/adapter/outbound/datastore"
)

var (
	mdb *sqlx.DB
	// rdb *cache.RedisClient
)

func InitSqlModule(host string, port int, user, password, dbName string) *sqlx.DB {
	db, err := datastore.NewMySqlDBConn(host, port, user, password, dbName)
	if err != nil {
		panic(err)
	}

	return db
}
