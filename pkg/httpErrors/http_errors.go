package httpErrors

import "fmt"

const (
	ErrBadRequest          = "Bad request"
	ErrInternalServerError = "Internal server error"
	ErrEmailAlreadyExists  = "User with given email already exists"
	ErrNoSuchUser          = "User not found"
	ErrWrongCredentials    = "Wrong Credentials"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrBadQueryParams      = "Invalid query params"
)

type ClientError struct {
	StatusCode string
	Message    string
}

func (e ClientError) Error() string {
	return fmt.Sprintf("code:%s, message:%s", e.StatusCode, e.Message)
}

type ServeError struct {
	StatusCode string
	Message    string
}

func (e ServeError) Error() string {
	return fmt.Sprintf("code:%s, message:%s", e.StatusCode, e.Message)
}
