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
	repo     entities.NewsLikeUseCaseInterface
	newsRepo entities.NewsUseCaseInterface
}

func NewNewsLikeController(repo entities.NewsLikeUseCaseInterface, newsRepo entities.NewsUseCaseInterface) *NewsLikeController {
	return &NewsLikeController{
		repo:     repo,
		newsRepo: newsRepo,
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

	_, err = n.newsRepo.GetByID(newsID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, base.NewErrorResponse("News not found"))
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

	if likeStatus == "liked" {
		err := n.repo.IncreaseTotalLikes(strconv.Itoa(newsID))
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
	} else {
		err := n.repo.DecreaseTotalLikes(strconv.Itoa(newsID))
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
		}
	}

	message := "News " + likeStatus

	successResponse := base.NewSuccessResponse(message, nil)
	return ctx.JSON(http.StatusOK, successResponse)
}
