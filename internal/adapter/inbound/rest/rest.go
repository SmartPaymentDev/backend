package rest

import (
	"github.com/labstack/echo/v4"
	"smart-payment-dev-be/internal/adapter/inbound/registry"
	"smart-payment-dev-be/internal/adapter/inbound/rest/group"
	hv "smart-payment-dev-be/internal/adapter/inbound/rest/handler"
)

func Apply(e *echo.Echo, g *echo.Group, serviceRegistry registry.ServiceRegistry) {
	handlerV1 := hv.New(serviceRegistry)
	group.NewUserV1(e, g, handlerV1)
	group.NewBillV1(e, g, handlerV1)
	group.NewTransactionV1(e, g, handlerV1)
}
