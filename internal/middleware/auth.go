package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"orc-system/internal/token"
	"strings"
)

type Authenticator struct {
	skipPaths    []string
	nologinPaths []string
}

func NewAuthenticator(skipPaths []string, nologinPaths []string) *Authenticator {
	return &Authenticator{
		skipPaths:    skipPaths,
		nologinPaths: nologinPaths,
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
	for _, v := range a.nologinPaths {
		if v == c.Path() {
			return a.parseBearerToken(c) == ""
		}
	}
	return false
}

func (a Authenticator) parseBearerToken(c echo.Context) string {
	auth := c.Request().Header.Get(echo.HeaderAuthorization)
	authScheme := "Bearer"
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:]
	}
	return ""
}
