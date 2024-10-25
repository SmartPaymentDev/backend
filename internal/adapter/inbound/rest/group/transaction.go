package group

import (
	"github.com/labstack/echo/v4"
	hv "smart-payment-dev-be/internal/adapter/inbound/rest/handler"
)

func NewTransactionV1(e *echo.Echo, g *echo.Group, handler *hv.Handler) {
	g.GET("/spp", handler.GetTransactionsByCustId())
	g.GET("/spp/balance", handler.GetTransactionByCustId())
	g.GET("/saku", handler.GetTransactionDetailsByCustId())
	g.GET("/saku/balance", handler.GetTransactionDetailByCustId())
	g.POST("/saku", handler.CreateTransaction())
}
