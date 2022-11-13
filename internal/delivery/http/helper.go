package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"orc-system/internal/token"
	"orc-system/pkg/httpErrors"
	"orc-system/pkg/logger"
	"orc-system/pkg/utils"
)

const (
	prettyIndent = "  "
)

type APIErrorResponse struct {
	Code      string `json:"code"`
	RequestID string `json:"requestId"`
	Message   string `json:"message,omitempty"`
}

type ContextInfo struct {
	UserID    string
	RequestID string
	logField  logger.Fields
}

func GetContextInfo(e echo.Context) (c context.Context, info ContextInfo, isLogin bool) {
	c = e.Request().Context()
	logFields := logger.Fields{}
	emptyStr := ""
	userID, ok := e.Get(token.UserIDKey).(string)
	if !ok || userID == emptyStr {
		logFields["user_id"] = emptyStr
		info.UserID = emptyStr
		isLogin = false
	} else {
		logFields["user_id"] = userID
		info.UserID = userID
		isLogin = true
	}
	if resp := e.Response(); resp != nil {
		reqID := utils.GetRequestID(e)
		info.RequestID = reqID
		logFields["request_id"] = reqID
	}
	logFields["login"] = isLogin
	info.logField = logFields
	return
}

func APIResponseOK(c echo.Context, data interface{}) error {
	return c.JSONPretty(http.StatusOK, data, prettyIndent)
}

func APIResponseError(c echo.Context, status int, message string) error {
	errResp := APIErrorResponse{Code: fmt.Sprintf("%d-000", status), Message: message}
	if resp := c.Response(); resp != nil {
		requestID := resp.Header().Get(echo.HeaderXRequestID)
		errResp.RequestID = requestID
	}
	return c.JSONPretty(status, errResp, prettyIndent)
}

func APIResponseCustomCode(c echo.Context, status int, customCode string, message string) error {
	errResp := APIErrorResponse{Code: fmt.Sprintf("%d-%d", status, customCode), Message: message}
	if resp := c.Response(); resp != nil {
		requestID := resp.Header().Get(echo.HeaderXRequestID)
		errResp.RequestID = requestID
	}
	return c.JSONPretty(status, errResp, prettyIndent)
}

func HandlerError(c echo.Context, err error) error {
	var clErr httpErrors.ClientError
	if errors.As(err, &clErr) {
		msg := httpErrors.ErrBadRequest
		if clErr.Message != "" {
			msg = clErr.Message
		}
		if clErr.CustomeCode != "" {
			return APIResponseCustomCode(c, http.StatusBadRequest, clErr.CustomeCode, msg)
		}
		return APIResponseError(c, http.StatusBadRequest, msg)
	}
	var svError httpErrors.ServeError
	if errors.As(err, &svError) {
		msg := httpErrors.ErrInternalServerError
		if svError.Message != "" {
			msg = svError.Message
		}
		if svError.CustomeCode != "" {
			return APIResponseCustomCode(c, http.StatusInternalServerError, svError.CustomeCode, msg)
		}
		return APIResponseError(c, http.StatusInternalServerError, msg)
	}
	return APIResponseError(c, http.StatusInternalServerError, httpErrors.ErrInternalServerError)
}
