package admin

import (
	"e-complaint-api/controllers/admin/base"
	"e-complaint-api/controllers/admin/request"
	"e-complaint-api/controllers/admin/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
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

	var adminsResponse []*response.GetAllAdmins
	for _, admin := range admins {
		adminsResponse = append(adminsResponse, response.GetAdminsFromEntitiesToResponse(&admin))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Admins", adminsResponse))
}

func (ac *AdminController) GetAdminByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID format"))
	}

	admin, err := ac.adminUseCase.GetAdminByID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.GetAdminsFromEntitiesToResponse(admin)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Admin By ID", adminResponse))
}

func (ac *AdminController) DeleteAdmin(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID format"))
	}

	err = ac.adminUseCase.DeleteAdmin(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewDeletedResponse("Success Delete Admin"))
}

func (ac *AdminController) UpdateAdmin(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Invalid ID format"))
	}

	var adminRequest request.UpdateAccount
	c.Bind(&adminRequest)

	admin, err := ac.adminUseCase.UpdateAdmin(id, adminRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.GetAdminsFromEntitiesToResponse(&admin)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Admin", adminResponse))
}
