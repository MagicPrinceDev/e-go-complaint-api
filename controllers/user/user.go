package user

import (
	"e-complaint-api/constants"
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/user/request"
	"e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase entities.UserUseCaseInterface
}

func NewUserController(userUseCase entities.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var userRequest request.Register
	c.Bind(&userRequest)

	user, err := uc.userUseCase.Register(userRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	userResponse := response.RegisterFromEntitiesToResponse(&user)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Register", userResponse))
}

func (uc *UserController) Login(c echo.Context) error {
	var userRequest request.Login
	c.Bind(&userRequest)

	user, err := uc.userUseCase.Login(userRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.LoginFromEntitiesToResponse(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", userResponse))
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.userUseCase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	usersResponse := response.GetAllUsersFromEntitiesToResponse(users)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Users", usersResponse))
}

func (uc *UserController) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrInvalidIDFormat.Error()))
	}

	user, err := uc.userUseCase.GetUserByID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.GetUsersFromEntitiesToResponse(user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get User By ID", userResponse))
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	jwtID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	var userRequest request.UpdateUser
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	user, err := uc.userUseCase.UpdateUser(jwtID, userRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UpdateUserFromEntitiesToResponse(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update User", userResponse))
}

func (uc *UserController) UpdateProfilePhoto(c echo.Context) error {
	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	profilePhoto, err := c.FormFile("profile_photo")
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrAllFieldsMustBeFilled.Error()))
	}

	// Check file format
	if profilePhoto.Header.Get("Content-Type") != "image/jpeg" && profilePhoto.Header.Get("Content-Type") != "image/png" && profilePhoto.Header.Get("Content-Type") != "image/jpg" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrInvalidFileFormat.Error()))
	}

	if profilePhoto.Size > 5*1024*1024 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrMaxFileSizeExceeded.Error()))
	}

	err = uc.userUseCase.UpdateProfilePhoto(userID, profilePhoto)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Profile Photo", nil))
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {

		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrInvalidIDFormat.Error()))
	}

	userRole, err := utils.GetRoleFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if userRole == "user" {
		jwtID, err := utils.GetIDFromJWT(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}

		if id != jwtID {
			return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(constants.ErrUnauthorized.Error()))
		}
	}

	err = uc.userUseCase.Delete(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete User", nil))
}

func (uc *UserController) UpdatePassword(c echo.Context) error {
	jwtID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	var passwordRequest request.UpdatePassword
	if err := c.Bind(&passwordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	err = uc.userUseCase.UpdatePassword(jwtID, passwordRequest.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Password", nil))
}

func (uc *UserController) SendOTPRegister(c echo.Context) error {
	var emailRequest request.SendOTP
	if err := c.Bind(&emailRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	err := uc.userUseCase.SendOTP(emailRequest.Email, "register")
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Send OTP", nil))
}

func (uc *UserController) VerifyOTPRegister(c echo.Context) error {
	var otpRequest request.VerifyOTP
	if err := c.Bind(&otpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	err := uc.userUseCase.VerifyOTP(otpRequest.Email, otpRequest.OTP, "register")
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Verify OTP", nil))
}

func (uc *UserController) SendOTPForgotPassword(c echo.Context) error {
	var emailRequest request.SendOTP
	if err := c.Bind(&emailRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	err := uc.userUseCase.SendOTP(emailRequest.Email, "forgot_password")
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Send OTP", nil))
}

func (uc *UserController) VerifyOTPForgotPassword(c echo.Context) error {
	var otpRequest request.VerifyOTP
	if err := c.Bind(&otpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	err := uc.userUseCase.VerifyOTP(otpRequest.Email, otpRequest.OTP, "forgot_password")
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Verify OTP", nil))
}

func (uc *UserController) UpdatePasswordForgot(c echo.Context) error {
	var passwordRequest request.UpdatePasswordForgot
	if err := c.Bind(&passwordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	err := uc.userUseCase.UpdatePasswordForgot(passwordRequest.Email, passwordRequest.NewPassword)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Password", nil))
}
