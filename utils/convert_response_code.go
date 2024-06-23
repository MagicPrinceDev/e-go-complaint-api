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
		constants.ErrMaxFileSizeExceeded,
		constants.ErrMaxFileCountExceeded,
		constants.ErrInvalidIDFormat,
		constants.ErrComplaintAlreadyVerified,
		constants.ErrComplaintAlreadyRejected,
		constants.ErrComplaintAlreadyOnProgress,
		constants.ErrComplaintAlreadyFinished,
		constants.ErrComplaintNotVerified,
		constants.ErrComplaintNotOnProgress,
		constants.ErrInvalidStatus,
		constants.ErrIDMustBeFilled,
		constants.ErrComplaintProcessCannotBeDeleted,
		constants.ErrEmailOrUsernameAlreadyExists,
		constants.ErrNoChangesDetected,
		constants.ErrCommentCannotBeEmpty,
		constants.ErrInvalidOTP,
		constants.ErrExpiredOTP,
		constants.ErrEmailNotVerified,
		constants.ErrPageMustBeFilled,
		constants.ErrLimitMustBeFilled,
		constants.ErrInvalidFileFormat,
		constants.ErrForgotPasswordOTPNotVerified,
		constants.ErrConfirmPasswordDoesntMatch,
		constants.ErrColumnsDoesntMatch,
		constants.ErrInvalidCategoryIDFormat,
		constants.ErrEmailNotRegistered,
		constants.ErrPasswordMustBeAtLeast8Characters,
		constants.ErrCategoryHasBeenUsed,
	}

	var notFoundErrors = []error{
		constants.ErrComplaintNotFound,
		constants.ErrRegencyNotFound,
		constants.ErrCategoryNotFound,
		constants.ErrComplaintProcessNotFound,
		constants.ErrAdminNotFound,
		constants.ErrNewsNotFound,
		constants.ErrUserNotFound,
		constants.ErrNotFound,
	}

	if contains(badRequestErrors, err) {
		return http.StatusBadRequest
	} else if contains(notFoundErrors, err) {
		return http.StatusNotFound
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
