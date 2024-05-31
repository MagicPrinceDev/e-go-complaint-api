package news

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"
	"strconv"

	news_response "e-complaint-api/controllers/news/response"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	newsUseCase entities.NewsUseCaseInterface
	// newsFileUseCase entities.NewsFileUseCaseInterface
}

func NewNewsController(newsUseCase entities.NewsUseCaseInterface) *NewsController {
	return &NewsController{
		newsUseCase: newsUseCase,
		// newsFileUseCase: newsFileUseCase,
	}
}

func (nc *NewsController) GetPaginated(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	category_filter, _ := strconv.Atoi(c.QueryParam("category_id"))
	filter := map[string]interface{}{}
	if category_filter == 0 {
		filter = nil
	} else {
		filter["category_id"] = category_filter
	}

	sort_by := c.QueryParam("sort_by")
	sort_type := c.QueryParam("sort_type")

	news, err := nc.newsUseCase.GetPaginated(limit, page, search, filter, sort_by, sort_type)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	newsResponses := []*news_response.Get{}
	for _, news := range news {
		newsResponses = append(newsResponses, news_response.GetFromEntitiesToResponse(&news))
	}

	metaData, err := nc.newsUseCase.GetMetaData(limit, page, search, filter)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	metaDataResponse := base.NewMetadata(metaData.TotalData, metaData.Pagination.TotalDataPerPage, metaData.Pagination.FirstPage, metaData.Pagination.LastPage, metaData.Pagination.CurrentPage, metaData.Pagination.NextPage, metaData.Pagination.PrevPage)

	return c.JSON(http.StatusOK, base.NewSuccessResponseWithMetadata("Success Get News", newsResponses, *metaDataResponse))
}

func (nc *NewsController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	news, err := nc.newsUseCase.GetByID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	newsResponse := news_response.GetFromEntitiesToResponse(&news)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get News By ID", newsResponse))
}
