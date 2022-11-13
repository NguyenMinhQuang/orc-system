package utils

import (
	"github.com/labstack/echo/v4"
)

func IsHTTPSuccess(code int) bool {
	return (200 <= code && code < 300)
}

// Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
