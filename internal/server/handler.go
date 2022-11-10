package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (s *Server) NewHTTPHandler(e *echo.Echo) error {
	var (
		loggerCfg = middleware.DefaultLoggerConfig
	)
	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/healthcheck"
	}
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.LoggerWithConfig(loggerCfg))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())

	// cors
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	//if s.cfg.AllowOrigins != "" {
	//	aos := strings.Split(s.cfg.AllowOrigins, ",")
	//	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//		AllowOrigins: aos,
	//	}))
	//}

	// cache-control
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-cache, no-store")
			c.Response().Header().Set("Pragma", "no-cache")
			c.Response().Header().Set("Expires", "0")
			return next(c)
		}
	})

	// healthcheck
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	//skipPaths := []string{
	//	"/healthcheck",
	//}
	//nologinPaths := []string{
	//	"/api/login",
	//}
	// middlerware
	// TODO: continue

	return nil
}
