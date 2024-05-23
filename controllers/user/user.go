package user

import (
	"e-complaint-api/controllers/admin/base"
	"e-complaint-api/controllers/user/request"
	"e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"

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
