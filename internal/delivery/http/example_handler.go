package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"orc-system/internal/usecase/example"
	"orc-system/pkg/httpErrors"
	"orc-system/pkg/logger"
)

type ExampleHandler struct {
	Example example.IUseCase
}

func NewExampleHandler(e *echo.Group, us example.IUseCase) {
	handler := &ExampleHandler{
		Example: us,
	}
	e.GET("/getuser", handler.GetUsers)
	e.GET("/listuser", handler.GetAllUser)
}

func (h *ExampleHandler) GetUsers(c echo.Context) error {
	ctx, _, _ := GetContextInfo(c)
	//if !isLogin {
	//	return APIResponseError(c, http.StatusUnauthorized, httpErrors.ErrUnauthorized)
	//}

	var param example.GetByIDInput
	if err := c.Bind(&param); err != nil {
		logger.Info(err)
		return APIResponseError(c, http.StatusBadRequest, httpErrors.ErrBadRequest)
	}

	if err := param.Validate(); err != nil {
		logger.Info(err)
		return APIResponseError(c, http.StatusBadRequest, httpErrors.ErrBadRequest)
	}

	resp, err := h.Example.GetByID(ctx, param)
	if err != nil {
		logger.Info(err)
		return HandlerError(c, err)
	}
	return APIResponseOK(c, resp)
}

func (h *ExampleHandler) GetAllUser(c echo.Context) error {
	ctx, _, _ := GetContextInfo(c)
	//if !isLogin {
	//	logger.Error(httpErrors.ErrUnauthorized)
	//	return APIResponseError(c, http.StatusUnauthorized, httpErrors.ErrUnauthorized)
	//}

	resp, err := h.Example.GetAllUser(ctx)
	if err != nil {
		logger.Error(err)
		return HandlerError(c, err)
	}
	return APIResponseOK(c, resp)
}
