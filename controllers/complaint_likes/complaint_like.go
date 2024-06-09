package complaint_like

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComplaintLikeController struct {
	complaintLikeUseCase entities.ComplaintLikeUseCaseInterface
	complaintUseCase     entities.ComplaintUseCaseInterface
}

func NewComplaintLikeController(complaintLikeUseCase entities.ComplaintLikeUseCaseInterface, complaintUseCase entities.ComplaintUseCaseInterface) *ComplaintLikeController {
	return &ComplaintLikeController{
		complaintLikeUseCase: complaintLikeUseCase,
		complaintUseCase:     complaintUseCase,
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

	likeStatus, err := c.complaintLikeUseCase.ToggleLike(complaintLike)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if likeStatus == "liked" {
		err := c.complaintUseCase.IncreaseTotalLikes(complaintID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
	} else {
		err := c.complaintUseCase.DecreaseTotalLikes(complaintID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
	}

	message := "Complaint " + likeStatus

	successResponse := base.NewSuccessResponse(message, nil)
	return ctx.JSON(http.StatusOK, successResponse)
}
