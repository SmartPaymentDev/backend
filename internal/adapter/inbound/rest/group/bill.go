package group

import (
	"github.com/labstack/echo/v4"
	hv "smart-payment-dev-be/internal/adapter/inbound/rest/handler"
)

func NewBillV1(e *echo.Echo, g *echo.Group, handler *hv.Handler) {
	g.GET("/bills", handler.GetBillsByCustId())
}
