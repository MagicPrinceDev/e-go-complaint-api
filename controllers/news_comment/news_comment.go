package news_comment

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/news_comment/request"
	"e-complaint-api/controllers/news_comment/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type NewsCommentController struct {
	newsCommentRepo entities.NewsCommentUseCaseInterface
	newsRepo        entities.NewsUseCaseInterface
}

func NewNewsCommentController(newsCommentRepo entities.NewsCommentUseCaseInterface, newsRepo entities.NewsUseCaseInterface) *NewsCommentController {
	return &NewsCommentController{
		newsCommentRepo: newsCommentRepo,
		newsRepo:        newsRepo,
	}
}

func (n *NewsCommentController) CommentNews(ctx echo.Context) error {
	newsIDStr := ctx.Param("news-id")
	if newsIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID required"))
	}

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID must be an integer"))
	}

	_, err = n.newsRepo.GetByID(newsID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News not found"))
	}

	userID, err := utils.GetIDFromJWT(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	role, err := utils.GetRoleFromJWT(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))

	}

	var req request.CommentNews
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	if role == "admin" || role == "super_admin" {
		req.AdminID = &userID
		req.UserID = nil
	} else {
		req.UserID = &userID
		req.AdminID = nil
	}

	comment := req.ToEntities(userID, newsID, role)
	if err := n.newsCommentRepo.CommentNews(comment); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	createCommentNews, err := n.newsCommentRepo.GetById(comment.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))

	}

	newsResponse := response.FromEntitiesToResponse(createCommentNews)

	return ctx.JSON(http.StatusCreated, base.NewSuccessResponse("Commented news successfully", newsResponse))

}

func (n *NewsCommentController) GetCommentNews(ctx echo.Context) error {
	newsIDStr := ctx.Param("news-id")
	if newsIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID required"))
	}

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID must be an integer"))
	}

	_, err = n.newsRepo.GetByID(newsID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News not found"))
	}

	comments, err := n.newsCommentRepo.GetByNewsId(newsID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if len(comments) == 0 {
		return ctx.JSON(http.StatusNotFound, base.NewErrorResponse("Comment not found"))
	}

	var commentsResponse []*response.NewsGet
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, response.FromEntitiesGetToResponse(&comment))
	}

	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Get comments news successfully", commentsResponse))
}

func (n *NewsCommentController) UpdateComment(ctx echo.Context) error {

	newsIDStr := ctx.Param("news-id")
	if newsIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID required"))
	}

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID must be an integer"))
	}

	commentIDStr := ctx.Param("comment-id")
	if commentIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment ID required"))
	}

	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment ID must be an integer"))
	}

	comment, err := n.newsCommentRepo.GetById(commentID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment not found"))
	}

	if comment.NewsID != newsID {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment does not belong to the specified news"))
	}

	userID, err := utils.GetIDFromJWT(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if comment.UserID != nil && *comment.UserID != userID {
		return ctx.JSON(http.StatusUnauthorized, base.NewErrorResponse("You are not authorized to update this comment"))
	}

	var req request.CommentNews
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	comment.Comment = req.Comment
	if err := n.newsCommentRepo.UpdateComment(comment); err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Comment updated successfully", nil))
}

func (n *NewsCommentController) DeleteComment(ctx echo.Context) error {
	newsIDStr := ctx.Param("news-id")
	if newsIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID required"))
	}

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("News ID must be an integer"))
	}

	commentIDStr := ctx.Param("comment-id")
	if commentIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment ID required"))
	}

	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment ID must be an integer"))
	}

	comment, err := n.newsCommentRepo.GetById(commentID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment not found"))
	}

	if comment.NewsID != newsID {
		return ctx.JSON(http.StatusBadRequest, base.NewErrorResponse("Comment does not belong to the specified news"))
	}

	userID, err := utils.GetIDFromJWT(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	role, err := utils.GetRoleFromJWT(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	if role != "admin" && comment.UserID != nil && *comment.UserID != userID {
		return ctx.JSON(http.StatusUnauthorized, base.NewErrorResponse("You are not authorized to delete this comment"))
	}

	if err := n.newsCommentRepo.DeleteComment(comment.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, base.NewSuccessResponse("Comment deleted successfully", nil))
}
