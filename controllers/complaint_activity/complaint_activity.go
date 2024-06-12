package complaint_activity

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/complaint_activity/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComplaintActivityController struct {
	complaintActivityUseCase entities.ComplaintActivityUseCaseInterface
	complaintUseCase         entities.ComplaintUseCaseInterface
}

func NewComplaintActivityController(complaintActivityUseCase entities.ComplaintActivityUseCaseInterface, complaintUseCase entities.ComplaintUseCaseInterface) *ComplaintActivityController {
	return &ComplaintActivityController{
		complaintActivityUseCase: complaintActivityUseCase,
		complaintUseCase:         complaintUseCase,
	}
}

func (ca *ComplaintActivityController) GetByComplaintID(c echo.Context) error {
	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	activityType := c.QueryParam("type")
	fmt.Println(activityType)

	complaintIDs, err := ca.complaintUseCase.GetComplaintIDsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	complaintActivities, err := ca.complaintActivityUseCase.GetByComplaintIDs(complaintIDs, activityType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	complaintActivitiesResponse := []response.Get{}
	for _, complaintActivity := range complaintActivities {
		// Cek apakah complaintActivity.Discussion dan complaintActivity.Like nil
		discussionUserID := complaintActivity.Discussion.UserID
		likeUserID := complaintActivity.Like.UserID

		// Pastikan untuk hanya menambahkan aktivitas yang tidak sesuai dengan userID dari JWT
		if (discussionUserID != nil && *discussionUserID == userID) || likeUserID == userID {
			continue
		}

		complaintActivitiesResponse = append(complaintActivitiesResponse, *response.GetFromEntitiesToResponse(&complaintActivity))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Complaint activities retrieved", complaintActivitiesResponse))
}
