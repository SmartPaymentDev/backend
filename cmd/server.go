package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	cfg "smart-payment-dev-be/config"
	iregistry "smart-payment-dev-be/internal/adapter/inbound/registry"
	"smart-payment-dev-be/internal/adapter/inbound/rest"
	mregistry "smart-payment-dev-be/internal/adapter/outbound/repository/registry"
	"syscall"

	"smart-payment-dev-be/shared/util"
)

func RunServer() {
	cfg := cfg.GetConfig()

	e := echo.New()

	// Using JWT
	jwtGroup := e.Group("")
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS512" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return []byte(cfg.SECREET_KEY), nil
			}

			// decryptToken := util.Decrypt(auth)
			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, fmt.Errorf("invalid token")
			}
			return token, nil
		},
	}))

	e.Use(middleware.BodyLimit("5M"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = util.CustomHTTPErrorHandler

	mHost := cfg.Host
	mPort := cfg.PostgresPort
	mUser := cfg.Username
	mPass := cfg.Password
	mDatabase := cfg.Dbname

	mdb = InitSqlModule(mHost, mPort, mUser, mPass, mDatabase)
	defer mdb.Close()

	// Init repository registry
	repositoryRegistry := mregistry.NewRepositoryRegistry(mdb) // db resgits

	// Init service registry
	serviceRegistry := iregistry.NewServiceRegistry(repositoryRegistry)

	rest.Apply(e, jwtGroup, *serviceRegistry)

	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}
