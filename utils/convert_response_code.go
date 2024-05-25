package utils

import (
	"e-complaint-api/constants"
	"net/http"
)

func ConvertResponseCode(err error) int {
	var badRequestErrors = []error{
		constants.ErrAllFieldsMustBeFilled,
		constants.ErrInvalidUsernameOrPassword,
		constants.ErrEmailAlreadyExists,
		constants.ErrUsernameAlreadyExists,
		constants.ErrInvalidJWT,
		constants.ErrOldPasswordDoesntMatch,
		constants.ErrLimitAndPageMustBeFilled,
		constants.ErrReportNotFound,
		constants.ErrRegencyNotFound,
		constants.ErrCategoryNotFound,
	}

	if contains(badRequestErrors, err) {
		return http.StatusBadRequest
	} else if err == constants.ErrUnauthorized {
		return http.StatusUnauthorized
	} else {
		return http.StatusInternalServerError
	}
}

func contains(slice []error, item error) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
