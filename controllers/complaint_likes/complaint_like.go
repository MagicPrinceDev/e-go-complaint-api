package complaint_like

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComplaintLikeController struct {
	useCase entities.ComplaintLikeUseCaseInterface
}

func NewComplaintLikeController(useCase entities.ComplaintLikeUseCaseInterface) *ComplaintLikeController {
	return &ComplaintLikeController{
		useCase: useCase,
	}
}

func (c *ComplaintLikeController) ToggleLike(ctx echo.Context) error {
	complaintID := ctx.Param("complaint-id")
	if complaintID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Complaint ID is required"})
	}

	userID, err := utils.GetIDFromJWT(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	complaintLike := &entities.ComplaintLike{
		UserID:      userID,
		ComplaintID: complaintID,
	}

	likeStatus, err := c.useCase.ToggleLike(complaintLike)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	message := "Complaint " + likeStatus

	successResponse := base.NewSuccessResponse(message, nil)
	return ctx.JSON(http.StatusOK, successResponse)
}
