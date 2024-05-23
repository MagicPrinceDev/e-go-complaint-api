package constants

import "errors"

var (
	ErrInternalServerError       = errors.New("internal server error")
	ErrAllFieldsMustBeFilled     = errors.New("all fields must be filled")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrEmailAlreadyExist         = errors.New("email already exist")
	ErrUsernameAlreadyExist      = errors.New("username already exist")
	ErrUnauthorized              = errors.New("unauthorized")
)
