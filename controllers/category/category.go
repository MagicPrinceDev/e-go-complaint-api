package category

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	useCase entities.CategoryUseCaseInterface
}

func NewCategoryController(useCase entities.CategoryUseCaseInterface) *CategoryController {
	return &CategoryController{useCase: useCase}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	categories, err := cc.useCase.GetAll()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all categories", categories))
}
