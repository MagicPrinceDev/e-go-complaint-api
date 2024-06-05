package admin

import (
	"e-complaint-api/constants"
	"e-complaint-api/controllers/admin/request"
	"e-complaint-api/controllers/admin/response"
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUseCase entities.AdminUseCaseInterface
}

func NewAdminController(adminUseCase entities.AdminUseCaseInterface) *AdminController {
	return &AdminController{
		adminUseCase: adminUseCase,
	}
}

func (ac *AdminController) CreateAccount(c echo.Context) error {
	var adminRequest request.CreateAccount
	c.Bind(&adminRequest)

	admin, err := ac.adminUseCase.CreateAccount(adminRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.CreateAccountFromEntitiesToResponse(&admin)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Account", adminResponse))
}

func (ac *AdminController) Login(c echo.Context) error {
	var adminRequest request.Login
	c.Bind(&adminRequest)

	admin, err := ac.adminUseCase.Login(adminRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	adminResponse := response.LoginFromEntitiesToResponse(&admin)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", adminResponse))
}

func (ac *AdminController) GetAllAdmins(c echo.Context) error {
	admins, err := ac.adminUseCase.GetAllAdmins()

	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var adminsResponse []*response.Get
	for _, admin := range admins {
		adminsResponse = append(adminsResponse, response.GetFromEntitiesToResponse(&admin))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Admins", adminsResponse))
}

func (ac *AdminController) GetAdminByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrInvalidIDFormat.Error()))
	}

	userRole, err := utils.GetRoleFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if userRole != "super_admin" && userID != id {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(constants.ErrUnauthorized.Error()))
	}

	admin, err := ac.adminUseCase.GetAdminByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrAdminNotFound) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	adminResponse := response.GetFromEntitiesToResponse(admin)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Admin By ID", adminResponse))
}

func (ac *AdminController) DeleteAdmin(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID format"))
	}

	jwtID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	userRole, err := utils.GetRoleFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if userRole == "super_admin" && id == jwtID {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(constants.ErrSuperAdminCannotDeleteThemselves.Error()))
	}

	err = ac.adminUseCase.DeleteAdmin(id)
	if err != nil {
		if errors.Is(err, constants.ErrAdminNotFound) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Admin", nil))
}

func (ac *AdminController) UpdateAdmin(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrInvalidIDFormat.Error()))
	}

	var adminRequest request.UpdateAccount
	if err := c.Bind(&adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	admin, err := ac.adminUseCase.UpdateAdmin(id, adminRequest.ToEntities())
	if err != nil {
		if errors.Is(err, constants.ErrAdminNotFound) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse("Admin account does not exist"))
		}
		if errors.Is(err, constants.ErrEmailAlreadyExists) || errors.Is(err, constants.ErrUsernameAlreadyExists) {
			return c.JSON(http.StatusConflict, base.NewErrorResponse(err.Error()))
		}
		if errors.Is(err, constants.ErrNoChangesDetected) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	adminResponse := response.UpdateUserFromEntitiesToResponse(&admin)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Admin", adminResponse))
}
