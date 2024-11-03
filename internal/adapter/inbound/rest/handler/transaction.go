package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	errorCode "smart-payment-dev-be/shared/error"
	"smart-payment-dev-be/shared/pagination"
	"smart-payment-dev-be/shared/util"
	"strings"
)

func (h *Handler) GetTransactionsByCustId() func(echo.Context) error {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")

		p := pagination.NewFromRequest(c)
		filter := dto.FilterTransaction{
			Token:   token,
			From:    c.QueryParam("from"),
			To:      c.QueryParam("to"),
			Page:    p.Page,
			PerPage: p.PerPage,
		}

		service := h.serviceRegistry.GetTransactionService()
		resp, err := service.GetTransactionByCustId(c.Request().Context(), filter)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", resp)
	}
}

func (h *Handler) GetTransactionByCustId() func(echo.Context) error {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")

		service := h.serviceRegistry.GetTransactionService()
		resp, err := service.GetVSaldoVaByCustId(c.Request().Context(), token)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", resp)
	}
}

func (h *Handler) GetTransactionDetailsByCustId() func(echo.Context) error {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")

		p := pagination.NewFromRequest(c)
		filter := dto.FilterTransactionDetail{
			Token:   token,
			From:    c.QueryParam("from"),
			To:      c.QueryParam("to"),
			Page:    p.Page,
			PerPage: p.PerPage,
		}

		service := h.serviceRegistry.GetTransactionService()
		resp, err := service.GetTransactionDetailByCustId(c.Request().Context(), filter)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", resp)
	}
}

func (h *Handler) GetTransactionDetailByCustId() func(echo.Context) error {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")

		service := h.serviceRegistry.GetTransactionService()
		resp, err := service.GetVSaldoVaCashlessByCustId(c.Request().Context(), token)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", resp)
	}
}

func (h *Handler) CreateTransaction() func(echo.Context) error {
	return func(c echo.Context) error {
		var req dto.TransactionRequest
		auth := c.Request().Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")

		service := h.serviceRegistry.GetTransactionService()
		if err := c.Bind(&req); err != nil {
			return err
		}

		err := service.CreateTransactionByCustId(c.Request().Context(), token, req)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", nil)
	}
}
