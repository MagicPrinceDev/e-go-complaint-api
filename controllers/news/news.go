package news

import (
	"e-complaint-api/constants"
	"e-complaint-api/controllers/base"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"
	"strconv"

	news_request "e-complaint-api/controllers/news/request"
	news_response "e-complaint-api/controllers/news/response"
	news_file_response "e-complaint-api/controllers/news_file/response"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	newsUseCase     entities.NewsUseCaseInterface
	newsFileUseCase entities.NewsFileUseCaseInterface
}

func NewNewsController(newsUseCase entities.NewsUseCaseInterface, newsFileUseCase entities.NewsFileUseCaseInterface) *NewsController {
	return &NewsController{
		newsUseCase:     newsUseCase,
		newsFileUseCase: newsFileUseCase,
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

func (nc *NewsController) Create(c echo.Context) error {
	admin_id, err := utils.GetIDFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, base.NewErrorResponse(err.Error()))
	}

	var newsRequest news_request.Create
	if err := c.Bind(&newsRequest); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	newsRequest.AdminID = admin_id

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}
	files := form.File["files"]

	if len(files) > 3 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Maximum 3 files allowed"))
	}

	totalFileSize := 0
	for _, file := range files {
		totalFileSize += int(file.Size)
	}

	if totalFileSize > 10*1024*1024 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrMaxFileSizeExceeded.Error()))
	}

	news, err := nc.newsUseCase.Create(newsRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	newsResponse := news_response.CreateFromEntitiesToResponse(&news)

	newsFile, err := nc.newsFileUseCase.Create(files, news.ID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	newsFileResponse := []*news_file_response.NewsFile{}
	for _, file := range newsFile {
		newsFileResponse = append(newsFileResponse, news_file_response.FromEntitiesToResponse(&file))
	}

	newsResponse.Files = newsFileResponse

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create News", newsResponse))
}

func (nc *NewsController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := nc.newsUseCase.Delete(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	err = nc.newsFileUseCase.DeleteByNewsID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete News", nil))
}
