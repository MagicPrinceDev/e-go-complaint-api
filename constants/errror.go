package constants

import "errors"

var (
	ErrInternalServerError       = errors.New("internal server error")
	ErrAllFieldsMustBeFilled     = errors.New("all fields must be filled")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrEmailAlreadyExists        = errors.New("email already exists")
	ErrUsernameAlreadyExists     = errors.New("username already exists")
	ErrUnauthorized              = errors.New("unauthorized")
	ErrInvalidJWT                = errors.New("invalid jwt")
	ErrLimitAndPageMustBeFilled  = errors.New("limit and page must be filled")
	ErrReportNotFound            = errors.New("report not found")
)
