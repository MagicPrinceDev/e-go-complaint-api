package complaint

import (
	"e-complaint-api/constants"
	"e-complaint-api/controllers/base"
	complaint_request "e-complaint-api/controllers/complaint/request"
	complaint_response "e-complaint-api/controllers/complaint/response"
	complaint_file_response "e-complaint-api/controllers/complaint_file/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ComplaintController struct {
	complaintUseCase     entities.ComplaintUseCaseInterface
	complaintFileUseCase entities.ComplaintFileUseCaseInterface
}

func NewComplaintController(complaintUseCase entities.ComplaintUseCaseInterface, complaintFileUseCase entities.ComplaintFileUseCaseInterface) *ComplaintController {
	return &ComplaintController{
		complaintUseCase:     complaintUseCase,
		complaintFileUseCase: complaintFileUseCase,
	}
}

func (cc *ComplaintController) GetPaginated(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	regency_filter := c.QueryParam("regency_id")
	category_filter, _ := strconv.Atoi(c.QueryParam("category_id"))
	status_filter := c.QueryParam("status")
	filter := map[string]interface{}{}
	if regency_filter == "" && category_filter == 0 && status_filter == "" {
		filter = nil
	} else {
		if regency_filter != "" {
			filter["regency_id"] = regency_filter
		}
		if category_filter != 0 {
			filter["category_id"] = category_filter
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

	complaintResponses := []*complaint_response.Get{}
	for _, complaint := range complaints {
		complaintResponses = append(complaintResponses, complaint_response.GetFromEntitiesToResponse(&complaint))
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

	complaintResponse := complaint_response.GetFromEntitiesToResponse(&complaint)

	return c.JSON(200, base.NewSuccessResponse("Success Get Report", complaintResponse))
}

func (cc *ComplaintController) Create(c echo.Context) error {
	user_id, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var complaintRequest complaint_request.Create
	c.Bind(&complaintRequest)
	complaintRequest.UserID = user_id

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}
	files := form.File["files"]

	if len(files) > 3 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrMaxFileCountExceeded.Error()))
	}

	// Count total file size
	totalFileSize := 0
	for _, file := range files {
		totalFileSize += int(file.Size)
	}

	if totalFileSize > 10*1024*1024 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrMaxFileSizeExceeded.Error()))
	}

	complaint, err1 := cc.complaintUseCase.Create(complaintRequest.ToEntities())
	if err1 != nil {
		return c.JSON(utils.ConvertResponseCode(err1), base.NewErrorResponse(err1.Error()))
	}

	complaintResponse := complaint_response.CreateFromEntitiesToResponse(&complaint)

	complaintFile, err2 := cc.complaintFileUseCase.Create(files, complaint.ID)
	if err2 != nil {
		return c.JSON(utils.ConvertResponseCode(err2), base.NewErrorResponse(err2.Error()))
	}

	complaintFileResponse := []*complaint_file_response.ComplaintFile{}
	for _, file := range complaintFile {
		complaintFileResponse = append(complaintFileResponse, complaint_file_response.FromEntitiesToResponse(&file))
	}

	complaintResponse.Files = complaintFileResponse

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Report", complaintResponse))
}

func (cc *ComplaintController) Delete(c echo.Context) error {
	id := c.Param("id")

	user_id, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	role, err := utils.GetRoleFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	err = cc.complaintUseCase.Delete(id, user_id, role)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	err = cc.complaintFileUseCase.DeleteByComplaintID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Report", nil))
}
