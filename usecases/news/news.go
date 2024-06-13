package news

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"strings"
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
	if limit != 0 && page == 0 {
		return nil, constants.ErrPageMustBeFilled
	} else if limit == 0 && page != 0 {
		return nil, constants.ErrLimitMustBeFilled
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

	if limit != 0 && page != 0 {
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
	} else {
		pagination.FirstPage = 1
		pagination.LastPage = 1
		pagination.CurrentPage = 1
		pagination.TotalDataPerPage = metaData.TotalData
		pagination.PrevPage = 0
		pagination.NextPage = 0
	}

	metaData.Pagination = pagination

	return metaData, nil
}

func (u *NewsUseCase) GetByID(id int) (entities.News, error) {
	news, err := u.repository.GetByID(id)
	if err != nil {
		return entities.News{}, err
	}

	return news, nil
}

func (u *NewsUseCase) Create(news *entities.News) (entities.News, error) {
	if news.Title == "" || news.Content == "" || news.CategoryID == 0 {
		return entities.News{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Create(news)
	if err != nil {
		if strings.HasSuffix(err.Error(), "REFERENCES `categories` (`id`))") {
			return entities.News{}, constants.ErrCategoryNotFound
		} else {
			return entities.News{}, constants.ErrInternalServerError
		}
	}

	return *news, nil
}

func (u *NewsUseCase) Delete(id int) error {
	err := u.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *NewsUseCase) Update(news entities.News) (entities.News, error) {
	if news.Title == "" || news.Content == "" || news.CategoryID == 0 {
		return entities.News{}, constants.ErrAllFieldsMustBeFilled
	}

	news, err := u.repository.Update(news)
	if err != nil {
		if strings.HasSuffix(err.Error(), "REFERENCES `categories` (`id`))") {
			return entities.News{}, constants.ErrCategoryNotFound
		} else {
			return entities.News{}, err
		}
	}

	return news, nil
}
