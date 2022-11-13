package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	exampleHandl "orc-system/internal/delivery/http"
	apiMiddleware "orc-system/internal/middleware"
	exampleRepo "orc-system/internal/repository/example"
	exampleSv "orc-system/internal/service/example"
	exampleUcase "orc-system/internal/usecase/example"
	"orc-system/pkg/logger"
	"orc-system/pkg/utils"
)

func (s *Server) NewHTTPHandler(e *echo.Echo) error {
	var (
		loggerCfg = middleware.DefaultLoggerConfig
	)
	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/api/v1/healthcheck"
	}

	e.Use(middleware.LoggerWithConfig(loggerCfg))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())

	// cors
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// cache-control
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-cache, no-store")
			c.Response().Header().Set("Pragma", "no-cache")
			c.Response().Header().Set("Expires", "0")
			return next(c)
		}
	})

	// non check login
	skipPaths := []string{
		"/api/v1/healthcheck",
		"/api/v1/example/listuser",
		"/api/v1/example/getuser",
	}
	e.Use(apiMiddleware.NewAuthenticator(skipPaths).Middleware(s.tokenMaker))

	// init repo
	exampleRepo := exampleRepo.NewExampleRepository(s.db)
	expService := exampleSv.NewExampleService(s.cfg.EndPoint)

	//init usecase
	exampleUc := exampleUcase.NewExampleUseCase(exampleRepo, expService)

	//handler
	v1 := e.Group("/api/v1")
	health := v1.Group("/healthcheck")
	exp := v1.Group("/example")

	exampleHandl.NewExampleHandler(exp, exampleUc)

	health.GET("", func(c echo.Context) error {
		logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	return nil
}
