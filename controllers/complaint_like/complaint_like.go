package complaint_like

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComplaintLikeController struct {
	complaintLikeUseCase     entities.ComplaintLikeUseCaseInterface
	complaintUseCase         entities.ComplaintUseCaseInterface
	complaintActivityUseCase entities.ComplaintActivityUseCaseInterface
}

func NewComplaintLikeController(complaintLikeUseCase entities.ComplaintLikeUseCaseInterface, complaintUseCase entities.ComplaintUseCaseInterface, complaintActivityUseCase entities.ComplaintActivityUseCaseInterface) *ComplaintLikeController {
	return &ComplaintLikeController{
		complaintLikeUseCase:     complaintLikeUseCase,
		complaintUseCase:         complaintUseCase,
		complaintActivityUseCase: complaintActivityUseCase,
	}
}

func (c *ComplaintLikeController) ToggleLike(ctx echo.Context) error {
	complaintID := ctx.Param("complaint-id")
	if complaintID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Complaint ID is required"})
	}

	_, err := c.complaintUseCase.GetByID(complaintID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, base.NewErrorResponse("Complaint not found"))
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

		var complaintActivity entities.ComplaintActivity
		complaintActivity.ComplaintID = complaintID
		complaintActivity.LikeID = &complaintLike.ID
		_, err = c.complaintActivityUseCase.Create(&complaintActivity)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}

	} else {
		err := c.complaintUseCase.DecreaseTotalLikes(complaintID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
		var complaintActivity entities.ComplaintActivity
		complaintActivity.ComplaintID = complaintID
		complaintActivity.LikeID = &complaintLike.ID
		err = c.complaintActivityUseCase.Delete(complaintActivity)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
	}

	message := "Complaint " + likeStatus

	successResponse := base.NewSuccessResponse(message, nil)
	return ctx.JSON(http.StatusOK, successResponse)
}
