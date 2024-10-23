package cmd

import (
	"fmt"
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
	os.Setenv("HTTP_PORT", "5001")
	cfg := cfg.GetConfig()

	e := echo.New()
	// for token
	jwtGroup := e.Group("")

	e.Use(middleware.BodyLimit("5M"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = util.CustomHTTPErrorHandler

	// Init repository registry
	repositoryRegistry := mregistry.NewRepositoryRegistry() // db resgits

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
