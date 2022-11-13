package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"orc-system/internal/token"
	"strings"
)

type Authenticator struct {
	skipPaths []string
}

func NewAuthenticator(skipPaths []string) *Authenticator {
	return &Authenticator{
		skipPaths: skipPaths,
	}
}

func (a *Authenticator) Middleware(tokenMaker token.Maker) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Skipper:    a.Skipper,
		Validator:  tokenMaker.VerifyToken,
	})
}

func (a *Authenticator) Skipper(c echo.Context) bool {
	for _, v := range a.skipPaths {
		if strings.HasPrefix(c.Path(), v) {
			return true
		}
	}
	return false
}
