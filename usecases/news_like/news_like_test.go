package news_like

import (
	"e-complaint-api/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockNewsLike struct {
	mock.Mock
}

func (m *MockNewsLike) FindByUserAndNews(userID int, newsID int) (*entities.NewsLike, error) {
	args := m.Called(userID, newsID)
	return args.Get(0).(*entities.NewsLike), args.Error(1)
}

func (m *MockNewsLike) Likes(newsLike *entities.NewsLike) error {
	args := m.Called(newsLike)
	return args.Error(0)
}

func (m *MockNewsLike) Unlike(newsLike *entities.NewsLike) error {
	args := m.Called(newsLike)
	return args.Error(0)
}

func (m *MockNewsLike) ToggleLike(newsLike *entities.NewsLike) (string, error) {
	args := m.Called(newsLike)
	return args.String(0), args.Error(1)
}

func (m *MockNewsLike) IncreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockNewsLike) DecreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewsLikeUseCase_ToggleLike(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNewsLike := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockNewsLike)
		newsLike := &entities.NewsLike{
			UserID: 1,
			NewsID: 1,
		}
		mockNewsLike.On("FindByUserAndNews", newsLike.UserID, newsLike.NewsID).Return(newsLike, nil)
		mockNewsLike.On("Unlike", newsLike).Return(nil)
		result, err := useCase.ToggleLike(newsLike)

		assert.NoError(t, err)
		assert.Equal(t, "unliked", result)
		mockNewsLike.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockNewsLike := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockNewsLike)
		newsLike := &entities.NewsLike{
			UserID: 1,
			NewsID: 1,
		}
		mockNewsLike.On("FindByUserAndNews", newsLike.UserID, newsLike.NewsID).Return((*entities.NewsLike)(nil), nil)
		mockNewsLike.On("Likes", newsLike).Return(nil)
		result, err := useCase.ToggleLike(newsLike)

		assert.NoError(t, err)
		assert.Equal(t, "liked", result)
		mockNewsLike.AssertExpectations(t)
	})

	t.Run("unlike", func(t *testing.T) {
		mockRepo := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockRepo)
		newsLike := &entities.NewsLike{
			UserID: 1,
			NewsID: 1,
		}

		mockRepo.On("FindByUserAndNews", newsLike.UserID, newsLike.NewsID).Return(newsLike, nil)
		mockRepo.On("Unlike", newsLike).Return(nil)

		result, err := useCase.ToggleLike(newsLike)

		assert.NoError(t, err)
		assert.Equal(t, "unliked", result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("like", func(t *testing.T) {
		mockRepo := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockRepo)
		newsLike := &entities.NewsLike{
			UserID: 1,
			NewsID: 1,
		}

		mockRepo.On("FindByUserAndNews", newsLike.UserID, newsLike.NewsID).Return((*entities.NewsLike)(nil), nil)
		mockRepo.On("Likes", newsLike).Return(nil)

		result, err := useCase.ToggleLike(newsLike)

		assert.NoError(t, err)
		assert.Equal(t, "liked", result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("like error", func(t *testing.T) {
		mockRepo := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockRepo)
		newsLike := &entities.NewsLike{
			UserID: 1,
			NewsID: 1,
		}

		mockRepo.On("FindByUserAndNews", newsLike.UserID, newsLike.NewsID).Return((*entities.NewsLike)(nil), nil)
		mockRepo.On("Likes", newsLike).Return(errors.New("error"))

		_, err := useCase.ToggleLike(newsLike)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("unlike error", func(t *testing.T) {
		mockRepo := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockRepo)
		newsLike := &entities.NewsLike{
			UserID: 1,
			NewsID: 1,
		}

		mockRepo.On("FindByUserAndNews", newsLike.UserID, newsLike.NewsID).Return(newsLike, nil)
		mockRepo.On("Unlike", newsLike).Return(errors.New("error"))

		_, err := useCase.ToggleLike(newsLike)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

}

func TestNewsLikeUseCase_IncreaseTotalLikes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNewsLike := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockNewsLike)
		id := "1"
		mockNewsLike.On("IncreaseTotalLikes", id).Return(nil)
		err := useCase.IncreaseTotalLikes(id)

		assert.NoError(t, err)
		mockNewsLike.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockNewsLike := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockNewsLike)
		id := "1"
		mockNewsLike.On("IncreaseTotalLikes", id).Return(errors.New("error"))
		err := useCase.IncreaseTotalLikes(id)

		assert.Error(t, err)
		mockNewsLike.AssertExpectations(t)
	})

}

func TestNewsLikeUseCase_DecreaseTotalLikes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockNewsLike := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockNewsLike)
		id := "1"
		mockNewsLike.On("DecreaseTotalLikes", id).Return(nil)
		err := useCase.DecreaseTotalLikes(id)

		assert.NoError(t, err)
		mockNewsLike.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockNewsLike := new(MockNewsLike)
		useCase := NewNewsLikeUseCase(mockNewsLike)
		id := "1"
		mockNewsLike.On("DecreaseTotalLikes", id).Return(errors.New("error"))
		err := useCase.DecreaseTotalLikes(id)

		assert.Error(t, err)
		mockNewsLike.AssertExpectations(t)
	})
}
