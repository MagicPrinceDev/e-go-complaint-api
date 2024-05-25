package complaint

import (
	"e-complaint-api/controllers/base"
	complaintResponse "e-complaint-api/controllers/complaint/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ComplaintController struct {
	complaintUseCase entities.ComplaintUseCaseInterface
}

func NewComplaintController(complaintUseCase entities.ComplaintUseCaseInterface) *ComplaintController {
	return &ComplaintController{
		complaintUseCase: complaintUseCase,
	}
}

func (cc *ComplaintController) GetPaginated(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	regency_filter := c.QueryParam("regency")
	category_filter, _ := strconv.Atoi(c.QueryParam("category"))
	status_filter := c.QueryParam("status")
	filter := map[string]interface{}{}
	if regency_filter == "" && category_filter == 0 && status_filter == "" {
		filter = nil
	} else {
		if regency_filter != "" {
			filter["regency"] = regency_filter
		}
		if category_filter != 0 {
			filter["category"] = category_filter
		}
		if status_filter != "" {
			filter["status"] = status_filter
		}
	}

	sort_by := c.QueryParam("sort_by")
	sort_type := c.QueryParam("sort_type")

	complaints, err := cc.complaintUseCase.GetPaginated(limit, page, search, filter, sort_by, sort_type)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	complaintResponses := []*complaintResponse.Get{}
	for _, complaint := range complaints {
		complaintResponses = append(complaintResponses, complaintResponse.GetFromEntitiesToResponse(&complaint))
	}

	metaData, err := cc.complaintUseCase.GetMetaData(limit, page, search, filter)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	metaDataResponse := base.NewMetadata(metaData.TotalData, metaData.Pagination.TotalDataPerPage, metaData.Pagination.FirstPage, metaData.Pagination.LastPage, metaData.Pagination.CurrentPage, metaData.Pagination.NextPage, metaData.Pagination.PrevPage)

	return c.JSON(200, base.NewSuccessResponseWithMetadata("Success Get Reports", complaintResponses, *metaDataResponse))
}

func (cc *ComplaintController) GetByID(c echo.Context) error {
	id := c.Param("id")

	complaint, err := cc.complaintUseCase.GetByID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	complaintResponse := complaintResponse.GetFromEntitiesToResponse(&complaint)

	return c.JSON(200, base.NewSuccessResponse("Success Get Report", complaintResponse))
}
