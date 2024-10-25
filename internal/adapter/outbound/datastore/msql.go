package datastore

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySqlDBConn(host string, port int, user, password, dbName string) (*sqlx.DB, error) {
	datasource := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, password, host, port, dbName)
	dbConn, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// Check if the connection is established
	err = dbConn.Ping()
	if err != nil {

		panic(err)
	}

	fmt.Println("Connected to MySQL successfully!")

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
