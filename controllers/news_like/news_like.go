package news_like

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type NewsLikeController struct {
	repo entities.NewsLikeUseCaseInterface
}

func NewNewsLikeController(repo entities.NewsLikeUseCaseInterface) *NewsLikeController {
	return &NewsLikeController{
		repo: repo,
	}
}

func (n *NewsLikeController) ToggleLike(ctx echo.Context) error {
	newsIDStr := ctx.Param("news-id")
	if newsIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "News ID is required"})
	}

	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "News ID must be an integer"})
	}

	userID, err := utils.GetIDFromJWT(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	newsLike := &entities.NewsLike{
		UserID: userID,
		NewsID: newsID,
	}

	likeStatus, err := n.repo.ToggleLike(newsLike)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	message := "News " + likeStatus

	successResponse := base.NewSuccessResponse(message, nil)
	return ctx.JSON(http.StatusOK, successResponse)
}
