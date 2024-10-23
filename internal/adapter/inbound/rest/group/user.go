package group

import (
	"github.com/labstack/echo/v4"
	hv "smart-payment-dev-be/internal/adapter/inbound/rest/handler"
)

func NewUserV1(e *echo.Echo, g *echo.Group, handler *hv.Handler) {
	g.POST("/v1/user", handler.CreateUser())
	g.GET("/v1/user", handler.GetUsers())
	g.GET("/v1/user/:id", handler.GetUserById())
	g.PUT("/v1/user/:id", handler.UpdateUserById())
	g.DELETE("/v1/user/:id", handler.DeleteUserById())
}
