package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	errorCode "smart-payment-dev-be/shared/error"
	"smart-payment-dev-be/shared/util"
	"strconv"
)

func (h *Handler) CreateUser() func(echo.Context) error {
	return func(c echo.Context) error {
		var banner dto.CreateUserRequest
		if err := c.Bind(&banner); err != nil {
			return err
		}

		service := h.serviceRegistry.GetUserService()
		err := service.CreateUser(c.Request().Context(), banner)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}

func (h *Handler) GetUsers() func(echo.Context) error {
	return func(c echo.Context) error {
		service := h.serviceRegistry.GetUserService()
		res, err := service.GetUser(c.Request().Context())
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", res)
	}
}

func (h *Handler) GetUserById() func(echo.Context) error {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		service := h.serviceRegistry.GetUserService()
		res, err := service.GetUserById(c.Request().Context(), id)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}
		return util.SetResponse(c, http.StatusOK, "success", res)
	}
}

func (h *Handler) UpdateUserById() func(echo.Context) error {
	return func(c echo.Context) error {
		var updateUser dto.UpdateUserRequest
		updateUser.Id, _ = strconv.Atoi(c.Param("id"))
		if err := c.Bind(&updateUser); err != nil {
			return err
		}

		service := h.serviceRegistry.GetUserService()
		err := service.UpdateUserById(c.Request().Context(), updateUser)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "Success", nil)
	}
}

func (h *Handler) DeleteUserById() func(echo.Context) error {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		service := h.serviceRegistry.GetUserService()
		err := service.DeleteUserById(c.Request().Context(), id)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}
		return util.SetResponse(c, http.StatusOK, "Success", nil)
	}
}
