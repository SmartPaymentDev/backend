package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	errorCode "smart-payment-dev-be/shared/error"
	"smart-payment-dev-be/shared/util"
	"strings"
)

func (h *Handler) LoginUser() func(echo.Context) error {
	return func(c echo.Context) error {
		var user dto.LoginUserRequest
		if err := c.Bind(&user); err != nil {
			return err
		}

		service := h.serviceRegistry.GetUserService()
		resp, err := service.LoginUser(c.Request().Context(), user)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", resp)
	}
}

func (h *Handler) GetUserByMe() func(echo.Context) error {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")

		service := h.serviceRegistry.GetUserService()
		resp, err := service.GetUserByMe(c.Request().Context(), token)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", resp)
	}
}
