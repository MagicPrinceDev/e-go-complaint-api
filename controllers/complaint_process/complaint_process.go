package complaint_process

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/complaint_process/request"
	"e-complaint-api/controllers/complaint_process/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ComplaintProcessController struct {
	complaintUseCase        entities.ComplaintUseCaseInterface
	complaintProcessUseCase entities.ComplaintProcessUseCaseInterface
}

func NewComplaintProcessController(complaintUseCase entities.ComplaintUseCaseInterface, complaintProcessUseCase entities.ComplaintProcessUseCaseInterface) *ComplaintProcessController {
	return &ComplaintProcessController{
		complaintUseCase:        complaintUseCase,
		complaintProcessUseCase: complaintProcessUseCase,
	}
}

func (cp *ComplaintProcessController) Create(c echo.Context) error {
	admin_id, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var complaintProcessRequest request.Create
	c.Bind(&complaintProcessRequest)

	complaintProcessRequest.AdminID = admin_id
	complaint_id := c.Param("complaint-id")
	complaintProcessRequest.ComplaintID = complaint_id

	complaintProcess, err := cp.complaintProcessUseCase.Create(complaintProcessRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	err = cp.complaintUseCase.UpdateStatus(complaint_id, complaintProcessRequest.Status)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	complaintProcessResponse := response.CreateFromEntitiesToResponse(&complaintProcess)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Complaint Process", complaintProcessResponse))
}

func (cp *ComplaintProcessController) GetByComplaintID(c echo.Context) error {
	complaint_id := c.Param("complaint-id")

	complaintProcesses, err := cp.complaintProcessUseCase.GetByComplaintID(complaint_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var complaintProcessesResponse []*response.Get
	for _, complaintProcess := range complaintProcesses {
		complaintProcessesResponse = append(complaintProcessesResponse, response.GetFromEntitiesToResponse(&complaintProcess))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Complaint Process", complaintProcessesResponse))
}

func (cp *ComplaintProcessController) Update(c echo.Context) error {
	admin_id, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	complaintID := c.Param("complaint-id")
	complaintProcessID, _ := strconv.Atoi(c.Param("process-id"))

	var complaintProcessRequest request.Update
	c.Bind(&complaintProcessRequest)
	complaintProcessRequest.ID = complaintProcessID
	complaintProcessRequest.AdminID = admin_id
	complaintProcessRequest.ComplaintID = complaintID

	complaintProcess, err := cp.complaintProcessUseCase.Update(complaintProcessRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	complaintProcessResponse := response.UpdateFromEntitiesToResponse(&complaintProcess)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Complaint Process", complaintProcessResponse))
}

func (cp *ComplaintProcessController) Delete(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	complaintProcessID, _ := strconv.Atoi(c.Param("process-id"))

	status, err := cp.complaintProcessUseCase.Delete(complaintID, complaintProcessID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	err = cp.complaintUseCase.UpdateStatus(complaintID, status)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Complaint Process", nil))
}
