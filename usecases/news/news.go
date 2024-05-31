package news

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
)

type NewsUseCase struct {
	repository entities.NewsRepositoryInterface
}

func NewNewsUseCase(repository entities.NewsRepositoryInterface) *NewsUseCase {
	return &NewsUseCase{
		repository: repository,
	}
}

func (u *NewsUseCase) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.News, error) {
	if limit == 0 || page == 0 {
		return nil, constants.ErrLimitAndPageMustBeFilled
	}

	if sortBy == "" {
		sortBy = "created_at"
	}

	if sortType == "" {
		sortType = "DESC"
	}

	news, err := u.repository.GetPaginated(limit, page, search, filter, sortBy, sortType)
	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return news, nil
}

func (u *NewsUseCase) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var pagination entities.Pagination
	metaData, err := u.repository.GetMetaData(limit, page, search, filter)

	if err != nil {
		return entities.Metadata{}, constants.ErrInternalServerError
	}

	pagination.FirstPage = 1
	pagination.LastPage = (metaData.TotalData + limit - 1) / limit
	pagination.CurrentPage = page
	if pagination.CurrentPage == pagination.LastPage {
		pagination.TotalDataPerPage = metaData.TotalData - (pagination.LastPage-1)*limit
	} else {
		pagination.TotalDataPerPage = limit
	}

	if page > 1 {
		pagination.PrevPage = page - 1
	} else {
		pagination.PrevPage = 0
	}

	if page < pagination.LastPage {
		pagination.NextPage = page + 1
	} else {
		pagination.NextPage = 0
	}

	metaData.Pagination = pagination

	return metaData, nil
}
