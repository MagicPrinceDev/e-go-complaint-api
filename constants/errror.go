package constants

import "errors"

var (
	ErrInternalServerError              = errors.New("internal server error")
	ErrAllFieldsMustBeFilled            = errors.New("all fields must be filled")
	ErrInvalidUsernameOrPassword        = errors.New("invalid username or password")
	ErrEmailAlreadyExists               = errors.New("email already exists")
	ErrUsernameAlreadyExists            = errors.New("username already exists")
	ErrUnauthorized                     = errors.New("unauthorized")
	ErrInvalidJWT                       = errors.New("invalid jwt")
	ErrOldPasswordDoesntMatch           = errors.New("old password doesn't match")
	ErrLimitAndPageMustBeFilled         = errors.New("limit and page must be filled")
	ErrComplaintNotFound                = errors.New("complaint not found")
	ErrFailedToCreateClientGCS          = errors.New("failed to create client gcs")
	ErrFailedToUploadObject             = errors.New("failed to upload object")
	ErrFailedToDeleteObject             = errors.New("failed to delete object")
	ErrRegencyNotFound                  = errors.New("regency not found")
	ErrCategoryNotFound                 = errors.New("category not found")
	ErrMaxFileSizeExceeded              = errors.New("max file size exceeded")
	ErrMaxFileCountExceeded             = errors.New("max file count exceeded")
	ErrInvalidIDFormat                  = errors.New("invalid id format")
	ErrEmailOrUsernameAlreadyExists     = errors.New("email or username already exists")
	ErrNoChangesDetected                = errors.New("no changes detected")
	ErrNotFound                         = errors.New("not found")
	ErrAdminNotFound                    = errors.New("admin account not found")
	ErrSuperAdminCannotDeleteThemselves = errors.New("super admin cannot delete themselves")
	ErrComplaintAlreadyVerified         = errors.New("complaint already verified")
	ErrComplaintAlreadyRejected         = errors.New("complaint already rejected")
	ErrComplaintAlreadyOnProgress       = errors.New("complaint already on progress")
	ErrComplaintAlreadyFinished         = errors.New("complaint already finished")
	ErrComplaintNotVerified             = errors.New("complaint not verified yet")
	ErrComplaintNotOnProgress           = errors.New("complaint not on progress yet")
	ErrInvalidStatus                    = errors.New("invalid status")
	ErrIDMustBeFilled                   = errors.New("id must be filled")
	ErrComplaintProcessNotFound         = errors.New("complaint process not found")
	ErrComplaintProcessCannotBeDeleted  = errors.New("complaint process cannot be deleted")
	ErrNewsNotFound                     = errors.New("news not found")
)
