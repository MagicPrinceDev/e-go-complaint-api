package news

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockNews struct {
	mock.Mock
}

func (m *MockNews) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.News, error) {
	args := m.Called(limit, page, search, filter, sortBy, sortType)
	return args.Get(0).([]entities.News), args.Error(1)
}

func (m *MockNews) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	args := m.Called(limit, page, search, filter)
	return args.Get(0).(entities.Metadata), args.Error(1)
}

func (m *MockNews) GetByID(id int) (entities.News, error) {
	args := m.Called(id)
	return args.Get(0).(entities.News), args.Error(1)
}

func (m *MockNews) Create(news *entities.News) error {
	args := m.Called(news)
	return args.Error(0)
}

func (m *MockNews) Update(news entities.News) (entities.News, error) {
	args := m.Called(news)
	return args.Get(0).(entities.News), args.Error(1)
}

func (m *MockNews) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewsUseCase_GetPaginated(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := []entities.News{
			{
				ID:         1,
				Title:      "title",
				Content:    "content",
				CategoryID: 1,
			},
			{
				ID:      2,
				Title:   "title2",
				Content: "content2",
			},
			{
				ID:      3,
				Title:   "title3",
				Content: "content3",
			},
		}
		mockNews.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "DESC").Return(news, nil)
		_, err := useCase.GetPaginated(10, 1, "", map[string]interface{}{}, "created_at", "DESC")
		assert.NoError(t, err)
	})

	t.Run("success - sort by is empty", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "DESC").Return([]entities.News{}, nil)
		_, err := useCase.GetPaginated(10, 1, "", map[string]interface{}{}, "", "DESC")
		assert.NoError(t, err)
	})

	t.Run("success - sort type is empty", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "DESC").Return([]entities.News{}, nil)
		_, err := useCase.GetPaginated(10, 1, "", map[string]interface{}{}, "created_at", "")
		assert.NoError(t, err)
	})

	t.Run("error - page must be filled", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		_, err := useCase.GetPaginated(10, 0, "", map[string]interface{}{}, "created_at", "DESC")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrPageMustBeFilled, err)
	})

	t.Run("error - limit must be filled", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		_, err := useCase.GetPaginated(0, 1, "", map[string]interface{}{}, "created_at", "DESC")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrLimitMustBeFilled, err)
	})

	t.Run("error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "DESC").Return([]entities.News{}, constants.ErrInternalServerError)
		_, err := useCase.GetPaginated(10, 1, "", map[string]interface{}{}, "created_at", "DESC")
		assert.Error(t, err)

	})

	t.Run("invalid sort type", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "INVALID").Return([]entities.News{}, errors.New("invalid sort type"))
		_, err := useCase.GetPaginated(10, 1, "", map[string]interface{}{}, "created_at", "INVALID")
		assert.Error(t, err)
	})

	t.Run("error - page must be filled", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		_, err := useCase.GetPaginated(10, 0, "", map[string]interface{}{}, "created_at", "DESC")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrPageMustBeFilled, err)
	})

	t.Run("error - limit must be filled", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		_, err := useCase.GetPaginated(0, 1, "", map[string]interface{}{}, "created_at", "DESC")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrLimitMustBeFilled, err)
	})

}

func TestNewsUseCase_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			Title:      "title",
			Content:    "content",
			CategoryID: 1,
		}
		mockNews.On("Create", &news).Return(nil)
		_, err := useCase.Create(&news)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			Title:      "title",
			Content:    "content",
			CategoryID: 1,
		}
		mockNews.On("Create", &news).Return(constants.ErrInternalServerError)
		_, err := useCase.Create(&news)
		assert.Error(t, err)
	})

	t.Run("empty title", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			Title:      "",
			Content:    "content",
			CategoryID: 1,
		}
		mockNews.On("Create", &news).Return(constants.ErrAllFieldsMustBeFilled)
		_, err := useCase.Create(&news)
		assert.Error(t, err)
	})

	t.Run("empty content", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			Title:      "title",
			Content:    "",
			CategoryID: 1,
		}
		mockNews.On("Create", &news).Return(constants.ErrAllFieldsMustBeFilled)
		_, err := useCase.Create(&news)
		assert.Error(t, err)
	})

	t.Run("empty category id", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			Title:      "title",
			Content:    "content",
			CategoryID: 0,
		}
		mockNews.On("Create", &news).Return(constants.ErrAllFieldsMustBeFilled)
		_, err := useCase.Create(&news)
		assert.Error(t, err)
	})

	t.Run("error - category not found due to reference error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			Title:      "title",
			Content:    "content",
			CategoryID: 999,
		}
		mockNews.On("Create", &news).Return(errors.New("foreign key constraint fails (`e-complaint-api`.`news`, CONSTRAINT `news_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`))"))
		_, err := useCase.Create(&news)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrCategoryNotFound, err)
	})
}

func TestNewsUseCase_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			ID:         1,
			Title:      "title",
			Content:    "content",
			CategoryID: 1,
		}
		mockNews.On("GetByID", 1).Return(news, nil)
		_, err := useCase.GetByID(1)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetByID", 1).Return(entities.News{}, constants.ErrInternalServerError)
		_, err := useCase.GetByID(1)
		assert.Error(t, err)
	})
}

func TestNewsUseCase_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("Delete", 1).Return(nil)
		err := useCase.Delete(1)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("Delete", 1).Return(constants.ErrInternalServerError)
		err := useCase.Delete(1)
		assert.Error(t, err)
	})
}

func TestNewsUseCase_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			ID:         1,
			Title:      "title",
			Content:    "content",
			CategoryID: 1,
		}
		mockNews.On("Update", news).Return(news, nil)
		_, err := useCase.Update(news)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			ID:         1,
			Title:      "title",
			Content:    "content",
			CategoryID: 1,
		}
		mockNews.On("Update", news).Return(entities.News{}, constants.ErrInternalServerError)
		_, err := useCase.Update(news)
		assert.Error(t, err)
	})

	t.Run("empty title", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			ID:         1,
			Title:      "",
			Content:    "content",
			CategoryID: 1,
		}
		mockNews.On("Update", news).Return(entities.News{}, constants.ErrAllFieldsMustBeFilled)
		_, err := useCase.Update(news)
		assert.Error(t, err)
	})

	t.Run("empty content", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			ID:         1,
			Title:      "title",
			Content:    "",
			CategoryID: 1,
		}
		mockNews.On("Update", news).Return(entities.News{}, constants.ErrAllFieldsMustBeFilled)
		_, err := useCase.Update(news)
		assert.Error(t, err)
	})

	t.Run("empty category id", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			ID:         1,
			Title:      "title",
			Content:    "content",
			CategoryID: 0,
		}
		mockNews.On("Update", news).Return(entities.News{}, constants.ErrAllFieldsMustBeFilled)
		_, err := useCase.Update(news)
		assert.Error(t, err)
	})

	t.Run("error - category not found due to reference error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		news := entities.News{
			ID:         1,
			Title:      "title",
			Content:    "content",
			CategoryID: 999,
		}
		mockNews.On("Update", news).Return(entities.News{}, errors.New("foreign key constraint fails (`e-complaint-api`.`news`, CONSTRAINT `news_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`))"))
		_, err := useCase.Update(news)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrCategoryNotFound, err)
	})
}

func TestNewsUseCase_GetMetaData(t *testing.T) {
	t.Run("success - with limit and page", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetMetaData", 10, 1, "", map[string]interface{}{}).Return(entities.Metadata{TotalData: 100}, nil)
		metaData, err := useCase.GetMetaData(10, 1, "", map[string]interface{}{})
		assert.NoError(t, err)
		assert.Equal(t, 1, metaData.Pagination.FirstPage)
		assert.Equal(t, 10, metaData.Pagination.LastPage)
		assert.Equal(t, 1, metaData.Pagination.CurrentPage)
		assert.Equal(t, 10, metaData.Pagination.TotalDataPerPage)
		assert.Equal(t, 0, metaData.Pagination.PrevPage)
		assert.Equal(t, 2, metaData.Pagination.NextPage)
	})

	t.Run("success - without limit and page", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetMetaData", 0, 0, "", map[string]interface{}{}).Return(entities.Metadata{TotalData: 100}, nil)
		metaData, err := useCase.GetMetaData(0, 0, "", map[string]interface{}{})
		assert.NoError(t, err)
		assert.Equal(t, 1, metaData.Pagination.FirstPage)
		assert.Equal(t, 1, metaData.Pagination.LastPage)
		assert.Equal(t, 1, metaData.Pagination.CurrentPage)
		assert.Equal(t, 100, metaData.Pagination.TotalDataPerPage)
		assert.Equal(t, 0, metaData.Pagination.PrevPage)
		assert.Equal(t, 0, metaData.Pagination.NextPage)
	})

	t.Run("success - current page is last page", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetMetaData", 10, 10, "", map[string]interface{}{}).Return(entities.Metadata{TotalData: 100}, nil)
		metaData, err := useCase.GetMetaData(10, 10, "", map[string]interface{}{})
		assert.NoError(t, err)
		assert.Equal(t, 10, metaData.Pagination.TotalDataPerPage)
	})

	t.Run("success - page is greater than 1", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetMetaData", 10, 2, "", map[string]interface{}{}).Return(entities.Metadata{TotalData: 100}, nil)
		metaData, err := useCase.GetMetaData(10, 2, "", map[string]interface{}{})
		assert.NoError(t, err)
		assert.Equal(t, 1, metaData.Pagination.PrevPage)
	})

	t.Run("success - page is less than last page", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetMetaData", 10, 1, "", map[string]interface{}{}).Return(entities.Metadata{TotalData: 100}, nil)
		metaData, err := useCase.GetMetaData(10, 1, "", map[string]interface{}{})
		assert.NoError(t, err)
		assert.Equal(t, 2, metaData.Pagination.NextPage)
	})

	t.Run("success - page is equal to last page", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetMetaData", 10, 10, "", map[string]interface{}{}).Return(entities.Metadata{TotalData: 100}, nil)
		metaData, err := useCase.GetMetaData(10, 10, "", map[string]interface{}{})
		assert.NoError(t, err)
		assert.Equal(t, 0, metaData.Pagination.NextPage)
	})

	t.Run("error - internal server error", func(t *testing.T) {
		mockNews := new(MockNews)
		useCase := NewNewsUseCase(mockNews)
		mockNews.On("GetMetaData", 10, 1, "", map[string]interface{}{}).Return(entities.Metadata{}, errors.New("internal server error"))
		_, err := useCase.GetMetaData(10, 1, "", map[string]interface{}{})
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})

}
