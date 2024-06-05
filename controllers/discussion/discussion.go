package discussion

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/discussion/request"
	"e-complaint-api/controllers/discussion/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DiscussionController struct {
	discussionUseCase entities.DiscussionUseCaseInterface
	complaintUsecase  entities.ComplaintUseCaseInterface
}

func NewDiscussionController(discussionUseCase entities.DiscussionUseCaseInterface, complaintUsecase entities.ComplaintUseCaseInterface) *DiscussionController {
	return &DiscussionController{
		discussionUseCase: discussionUseCase,
		complaintUsecase:  complaintUsecase,
	}
}

func (dc *DiscussionController) CreateDiscussion(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Complaint ID is required",
		})
	}

	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	role, err := utils.GetRoleFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))

	}

	var req request.CreateDiscussion
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))

	}

	if role == "admin" {
		req.AdminID = &userID
		req.UserID = nil
	} else {
		req.UserID = &userID
		req.AdminID = nil
	}

	discussionEntity := req.ToEntities(userID, complaintID, role)
	err = dc.discussionUseCase.Create(discussionEntity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))

	}

	createdDiscussion, err := dc.discussionUseCase.GetById(discussionEntity.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))

	}

	discussionResponse := response.FromEntitiesToResponse(createdDiscussion)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Discussion created successfully", discussionResponse))
}

func (dc *DiscussionController) GetDiscussionByComplaintID(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Complaint ID is required",
		})
	}

	discussions, err := dc.discussionUseCase.GetByComplaintID(complaintID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Discussion not found",
		})
	}

	var discussionsResponse []*response.GetDiscussion
	for _, discussion := range *discussions {
		discussionResponse := response.FromEntitiesGetToResponse(&discussion)
		discussionsResponse = append(discussionsResponse, discussionResponse)
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Discussion found", discussionsResponse))
}

func (dc *DiscussionController) UpdateDiscussion(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Complaint ID is required",
		})
	}

	discussionIDStr := c.Param("discussion-id")
	if discussionIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Discussion ID is required",
		})
	}

	discussionID, err := strconv.Atoi(discussionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))

	}

	discussion, err := dc.discussionUseCase.GetById(discussionID)
	if err != nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
	}

	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(err.Error()))
	}

	if discussion.UserID != nil && *discussion.UserID != userID {
		return c.JSON(http.StatusForbidden, base.NewErrorResponse(err.Error()))

	}

	var req request.CreateDiscussion
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))

	}

	discussion.Comment = req.Comment
	err = dc.discussionUseCase.Update(discussion)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	updatedDiscussion, err := dc.discussionUseCase.GetById(discussion.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	discussionResponse := response.FromEntitiesToResponse(updatedDiscussion)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Discussion updated successfully", discussionResponse))
}

func (dc *DiscussionController) DeleteDiscussion(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Complaint ID is required",
		})
	}

	discussionIDStr := c.Param("discussion-id")
	if discussionIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Discussion ID is required",
		})
	}

	discussionID, err := strconv.Atoi(discussionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	discussion, err := dc.discussionUseCase.GetById(discussionID)
	if err != nil || discussion == nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
	}

	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(err.Error()))
	}

	if discussion.UserID != nil && *discussion.UserID != userID {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse("You are not authorized to delete this discussion"))
	}

	err = dc.discussionUseCase.Delete(discussionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Discussion deleted successfully", nil))
}
