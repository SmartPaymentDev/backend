package util

import (
	"fmt"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message string

	if he, ok := err.(*echo.HTTPError); ok {
		// If the error is an HTTPError
		code = he.Code
		message = he.Message.(string)
	} else {
		// Default to internal server error
		code = http.StatusInternalServerError
		message = "Internal Server Error"

	}
	// Respond with custom error response
	SetResponse(c, code, http.StatusText(code), message)
}

func SetResponse(c echo.Context, statusCode int, msg string, data interface{}) error {
	conditionNotfound := fmt.Sprintf("%v", data)
	if strings.Contains(strings.ToLower(conditionNotfound), "not found") {
		statusCode = http.StatusNotFound
		msg = http.StatusText(statusCode)
	}

	res := c.JSON(statusCode, JSONResponse{
		Code:    statusCode,
		Message: msg,
		Data:    data,
	})

	return res
}
