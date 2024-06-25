package regency

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/regency/response"
	"e-complaint-api/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegencyController struct {
	useCase entities.RegencyUseCaseInterface
}

func NewRegencyController(useCase entities.RegencyUseCaseInterface) *RegencyController {
	return &RegencyController{useCase: useCase}
}

func (rc *RegencyController) GetAll(c echo.Context) error {
	regencies, err := rc.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var responseRegencies []*response.Regency
	for _, regency := range regencies {
		responseRegencies = append(responseRegencies, response.FromEntitiesToResponse(&regency))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all regencies", responseRegencies))
}
