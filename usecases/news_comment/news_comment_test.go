package news_comment

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockComment struct {
	mock.Mock
}

func (nc *MockComment) CommentNews(newsComment *entities.NewsComment) error {
	args := nc.Called(newsComment)
	return args.Error(0)
}

func (nc *MockComment) GetById(id int) (*entities.NewsComment, error) {
	args := nc.Called(id)
	return args.Get(0).(*entities.NewsComment), args.Error(1)
}

func (nc *MockComment) GetByNewsId(newsId int) ([]entities.NewsComment, error) {
	args := nc.Called(newsId)
	return args.Get(0).([]entities.NewsComment), args.Error(1)
}

func (nc *MockComment) UpdateComment(newsComment *entities.NewsComment) error {
	args := nc.Called(newsComment)
	return args.Error(0)
}

func (nc *MockComment) DeleteComment(id int) error {
	args := nc.Called(id)
	return args.Error(0)
}

func TestNewsCommentUseCase_CommentNews(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		comment := &entities.NewsComment{
			Comment: "Test comment",
		}
		mockRepo.On("CommentNews", comment).Return(nil)
		err := ncu.CommentNews(comment)
		assert.Nil(t, err)
	})

	t.Run("comment cannot be empty", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		emptyComment := &entities.NewsComment{
			Comment: "",
		}
		mockRepo.On("CommentNews", emptyComment).Return(nil)
		err := ncu.CommentNews(emptyComment)
		assert.Equal(t, constants.ErrCommentCannotBeEmpty, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		comment := &entities.NewsComment{
			Comment: "Test comment",
		}
		mockRepo.On("CommentNews", comment).Return(constants.ErrInternalServerError)
		err := ncu.CommentNews(comment)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})
}

func TestNewsCommentUseCase_GetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		mockRepo.On("GetById", 1).Return(&entities.NewsComment{}, nil)
		_, err := ncu.GetById(1)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		mockRepo.On("GetById", 1).Return(&entities.NewsComment{}, constants.ErrInternalServerError)
		_, err := ncu.GetById(1)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})

}

func TestNewsCommentUseCase_GetByNewsId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		mockRepo.On("GetByNewsId", 1).Return([]entities.NewsComment{}, nil)
		_, err := ncu.GetByNewsId(1)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		mockRepo.On("GetByNewsId", 1).Return([]entities.NewsComment{}, constants.ErrInternalServerError)
		_, err := ncu.GetByNewsId(1)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})
}

func TestNewsCommentUseCase_UpdateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		comment := &entities.NewsComment{
			Comment: "Test comment",
		}
		mockRepo.On("UpdateComment", comment).Return(nil)
		err := ncu.UpdateComment(comment)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		comment := &entities.NewsComment{
			Comment: "Test comment",
		}
		mockRepo.On("UpdateComment", comment).Return(constants.ErrInternalServerError)
		err := ncu.UpdateComment(comment)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})
}

func TestNewsCommentUseCase_DeleteComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		mockRepo.On("DeleteComment", 1).Return(nil)
		err := ncu.DeleteComment(1)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockComment)
		ncu := NewNewsCommentUseCase(mockRepo)
		mockRepo.On("DeleteComment", 1).Return(constants.ErrInternalServerError)
		err := ncu.DeleteComment(1)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})
}
