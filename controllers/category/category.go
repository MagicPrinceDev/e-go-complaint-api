package category

import (
	"e-complaint-api/constants"
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/category/request"
	"e-complaint-api/controllers/category/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"
	"net/http"
	"strconv"

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

	responseCategories := make([]*response.Get, len(categories))
	for i, category := range categories {
		responseCategories[i] = response.GetFromEntitiesToResponse(&category)
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all categories", responseCategories))
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("ID must be an integer"))
		}
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	category, err := cc.useCase.GetByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(constants.ErrCategoryNotFound.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	responseCategory := response.GetFromEntitiesToResponse(&category)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get category by ID", responseCategory))
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
	var category request.CreateCategories
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	createdCategory, err := cc.useCase.CreateCategory(category.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	responseCategory := response.FromUseCaseToResponse(createdCategory)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success create category", responseCategory))
}

func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("ID must be an integer"))
		}
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	if idStr == "" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("ID must be filled"))
	}

	_, err = cc.useCase.GetByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(constants.ErrCategoryNotFound.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var category request.CreateCategories
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	updatedCategory, err := cc.useCase.UpdateCategory(id, category.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	updatedCategory.ID = id

	responseCategory := response.FromUseCaseToResponse(updatedCategory)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success update category", responseCategory))
}

func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("ID must be an integer"))
		}
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}
	_, err = cc.useCase.GetByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(constants.ErrCategoryNotFound.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	err = cc.useCase.DeleteCategory(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success delete category", nil))
}
