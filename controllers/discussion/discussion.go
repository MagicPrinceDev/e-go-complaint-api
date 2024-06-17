package discussion

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/discussion/request"
	"e-complaint-api/controllers/discussion/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DiscussionController struct {
	discussionUseCase        entities.DiscussionUseCaseInterface
	complaintUsecase         entities.ComplaintUseCaseInterface
	complaintActivityUseCase entities.ComplaintActivityUseCaseInterface
}

func NewDiscussionController(discussionUseCase entities.DiscussionUseCaseInterface, complaintUsecase entities.ComplaintUseCaseInterface, complaintActivityUseCase entities.ComplaintActivityUseCaseInterface) *DiscussionController {
	return &DiscussionController{
		discussionUseCase:        discussionUseCase,
		complaintUsecase:         complaintUsecase,
		complaintActivityUseCase: complaintActivityUseCase,
	}
}

func (dc *DiscussionController) CreateDiscussion(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Complaint ID is required",
		})
	}

	_, err := dc.complaintUsecase.GetByID(complaintID)
	if err != nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("Complaint not found"))
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

	if role == "admin" || role == "super_admin" {
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

	var complaintActivity entities.ComplaintActivity
	complaintActivity.ComplaintID = complaintID
	complaintActivity.DiscussionID = &discussionEntity.ID
	_, err = dc.complaintActivityUseCase.Create(&complaintActivity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	discussionResponse := response.FromEntitiesToResponse(createdDiscussion)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Discussion created successfully", discussionResponse))
}

func (dc *DiscussionController) GetDiscussionByComplaintID(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Complaint Is Required"))
	}

	_, err := dc.complaintUsecase.GetByID(complaintID)
	if err != nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("Complaint not found"))
	}

	discussions, err := dc.discussionUseCase.GetByComplaintID(complaintID)
	if err != nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("Error retrieving discussions"))
	}

	if len(*discussions) == 0 {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("No discussions found for this complaint"))
	}

	var discussionsResponse []*response.DiscussionGet
	for _, discussion := range *discussions {
		discussionResponse := response.FromEntitiesGetToResponse(&discussion)
		discussionsResponse = append(discussionsResponse, discussionResponse)
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Discussion found", discussionsResponse))
}

func (dc *DiscussionController) UpdateDiscussion(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Complaint Is Required"))

	}

	_, err := dc.complaintUsecase.GetByID(complaintID)
	if err != nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("Complaint not found"))
	}

	discussionIDStr := c.Param("discussion-id")
	if discussionIDStr == "" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Discussion ID is required"))

	}

	discussionID, err := strconv.Atoi(discussionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	discussion, err := dc.discussionUseCase.GetById(discussionID)
	if err != nil || discussion == nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("Discussion not found"))
	}

	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(err.Error()))
	}

	if discussion.UserID != nil && *discussion.UserID != userID {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse("You are not authorized to update this discussion"))
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

	discussionResponse := response.FromEntitiesUpdateToResponse(updatedDiscussion)

	var complaintActivity entities.ComplaintActivity
	complaintActivity.ComplaintID = complaintID
	complaintActivity.DiscussionID = &discussionID
	err = dc.complaintActivityUseCase.Update(complaintActivity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Discussion updated successfully", discussionResponse))
}

func (dc *DiscussionController) DeleteDiscussion(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Complaint ID is required"))
	}

	_, err := dc.complaintUsecase.GetByID(complaintID)
	if err != nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("Complaint not found"))
	}

	discussionIDStr := c.Param("discussion-id")
	if discussionIDStr == "" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Discussion ID is required"))
	}
	discussionID, err := strconv.Atoi(discussionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	discussion, err := dc.discussionUseCase.GetById(discussionID)
	if err != nil || discussion == nil {
		return c.JSON(http.StatusNotFound, base.NewErrorResponse("Discussion not found"))
	}

	userID, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(err.Error()))
	}

	role, err := utils.GetRoleFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(err.Error()))
	}

	if role != "admin" && discussion.UserID != nil && *discussion.UserID != userID {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse("You are not authorized to delete this discussion"))
	}

	var complaintActivity entities.ComplaintActivity
	complaintActivity.ComplaintID = complaintID
	complaintActivity.DiscussionID = &discussionID
	err = dc.complaintActivityUseCase.Delete(complaintActivity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	err = dc.discussionUseCase.Delete(discussionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Discussion deleted successfully", nil))
}

func (dc *DiscussionController) GetAnswerRecommendation(c echo.Context) error {
	complaintID := c.Param("complaint-id")
	if complaintID == "" {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Complaint ID is required"))
	}

	answer, err := dc.discussionUseCase.GetAnswerRecommendation(complaintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	var answerResponse response.AnswerRecommendation
	answerResponse.Answer = answer

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Answer recommendation found", answerResponse))
}
