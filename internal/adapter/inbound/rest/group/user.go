package group

import (
	"github.com/labstack/echo/v4"
	hv "smart-payment-dev-be/internal/adapter/inbound/rest/handler"
)

func NewUserV1(e *echo.Echo, g *echo.Group, handler *hv.Handler) {
	e.POST("/login", handler.LoginUser())
	g.GET("/user/me", handler.GetUserByMe())
}
