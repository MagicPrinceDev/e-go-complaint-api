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
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	createCommentNews, err := n.newsCommentRepo.GetById(comment.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))

	}

	newsResponse := response.FromEntitiesToResponse(createCommentNews)

	return ctx.JSON(http.StatusCreated, base.NewSuccessResponse("Commented news successfully", newsResponse))

}
