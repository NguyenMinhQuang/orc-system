package httpErrors

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
	Message     string
	CustomeCode string
}

type ServeError struct {
	Message     string
	CustomeCode string
}
