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

func (h *Handler) GetBillsByCustId() func(echo.Context) error {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		token := strings.ReplaceAll(auth, "Bearer ", "")

		p := pagination.NewFromRequest(c)
		filter := dto.FilterBill{
			Token:     token,
			YearMonth: c.QueryParam("yearmonth"),
			PaidSt:    c.QueryParam("paid_st"),
			Page:      p.Page,
			PerPage:   p.PerPage,
		}

		service := h.serviceRegistry.GetBillService()
		resp, err := service.GetBillByCustId(c.Request().Context(), filter)
		if err != nil {
			return util.SetResponse(c, http.StatusInternalServerError, errorCode.INTERNAL_SERVER_ERROR, err.Error())
		}

		return util.SetResponse(c, http.StatusOK, "success", resp)
	}
}
