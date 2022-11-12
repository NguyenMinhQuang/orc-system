package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"orc-system/internal/usecase/example"
	"orc-system/pkg/httpErrors"
	"orc-system/pkg/logger"
	"orc-system/pkg/utils"
)

type ExampleHandler struct {
	Example example.IUseCase
	logger  logger.Logger
}

func NewExampleHandler(e *echo.Group, us example.IUseCase, log logger.Logger) {
	handler := &ExampleHandler{
		Example: us,
		logger:  log,
	}
	e.GET("/example", handler.GetUsers)
	e.GET("/listuser", handler.GetAllUser)
}

func (h *ExampleHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	var param example.GetByIDInput

	if err := c.Bind(&param); err != nil {
		h.logger.Info(err)
		return utils.APIResponseError(c, http.StatusBadRequest, httpErrors.ErrBadRequest)
	}

	if err := param.Validate(); err != nil {
		h.logger.Info(err)
		return utils.APIResponseError(c, http.StatusBadRequest, httpErrors.ErrBadRequest)
	}

	resp, err := h.Example.GetByID(ctx, param)
	if err != nil {
		h.logger.Info(err)
		return utils.HandlerError(c, err)
	}
	return utils.APIResponseOK(c, resp)
}

func (h *ExampleHandler) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.Example.GetAllUser(ctx)
	if err != nil {
		h.logger.Info(err)
		return utils.HandlerError(c, err)
	}
	return utils.APIResponseOK(c, resp)
}
