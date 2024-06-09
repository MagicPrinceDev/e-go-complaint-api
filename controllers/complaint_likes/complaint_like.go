package complaint_like

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ComplaintLikeController struct {
	useCase          entities.ComplaintLikeUseCaseInterface
	useCaseComplaint entities.ComplaintUseCaseInterface
}

func NewComplaintLikeController(useCase entities.ComplaintLikeUseCaseInterface, useCaseComplaint entities.ComplaintUseCaseInterface) *ComplaintLikeController {
	return &ComplaintLikeController{
		useCase:          useCase,
		useCaseComplaint: useCaseComplaint,
	}
}

func (c *ComplaintLikeController) ToggleLike(ctx echo.Context) error {
	complaintID := ctx.Param("complaint-id")
	if complaintID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Complaint ID is required"})
	}

	_, err := c.useCaseComplaint.GetByID(complaintID)
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

	existingComplaintLike, err := c.useCase.FindByUserAndComplaint(complaintLike.UserID, complaintLike.ComplaintID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if existingComplaintLike != nil {
		err = c.useCase.Unlike(complaintLike)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
		return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Unlike", nil))
	} else {
		err = c.useCase.Likes(complaintLike)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
		return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Success Like", nil))
	}
}
