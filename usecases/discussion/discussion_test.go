package discussion

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockDiscussion struct {
	mock.Mock
}

func (m *MockDiscussion) Create(discussion *entities.Discussion) error {
	args := m.Called(discussion)
	return args.Error(0)
}

func (m *MockDiscussion) GetById(id int) (*entities.Discussion, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Discussion), args.Error(1)
}

func (m *MockDiscussion) GetByComplaintID(complaintID string) (*[]entities.Discussion, error) {
	args := m.Called(complaintID)
	return args.Get(0).(*[]entities.Discussion), args.Error(1)
}

func (m *MockDiscussion) Update(discussion *entities.Discussion) error {
	args := m.Called(discussion)
	return args.Error(0)
}

func (m *MockDiscussion) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockFaq struct {
	mock.Mock
}

func (m *MockFaq) GetAll() ([]entities.Faq, error) {
	args := m.Called()
	return args.Get(0).([]entities.Faq), args.Error(1)

}

type OpenAIAPI struct {
	mock.Mock
}

func (m *OpenAIAPI) GetResponse(text string) (string, error) {
	args := m.Called(text)
	return args.String(0), args.Error(1)
}

func (m *OpenAIAPI) GetChatCompletion(prompt []string, userPrompt string) (string, error) {
	args := m.Called(prompt, userPrompt)
	return args.String(0), args.Error(1)
}

func TestDiscussionUseCase_GetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussion := entities.Discussion{
			ID:      1,
			Comment: "Hello",
		}

		mockDiscussion.On("GetById", 1).Return(&discussion, nil)
		result, err := useCase.GetById(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)

	})

	t.Run("discussion not found", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		mockDiscussion.On("GetById", 1).Return((*entities.Discussion)(nil), constants.ErrDiscussionNotFound)
		result, err := useCase.GetById(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		mockDiscussion.On("GetById", 1).Return((*entities.Discussion)(nil), constants.ErrInternalServerError)
		result, err := useCase.GetById(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func TestDiscussionUseCase_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussion := entities.Discussion{
			Comment: "Hello",
		}

		mockDiscussion.On("Create", &discussion).Return(nil)
		err := useCase.Create(&discussion)
		assert.Nil(t, err)
	})

	t.Run("comment cannot be empty", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussion := entities.Discussion{
			Comment: "",
		}

		err := useCase.Create(&discussion)
		assert.NotNil(t, err)
		assert.Equal(t, constants.ErrCommentCannotBeEmpty, err)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussion := entities.Discussion{
			Comment: "Hello",
		}

		mockDiscussion.On("Create", &discussion).Return(constants.ErrInternalServerError)
		err := useCase.Create(&discussion)
		assert.NotNil(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})
}

func TestDiscussionUseCase_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussion := entities.Discussion{
			Comment: "Hello",
		}

		mockDiscussion.On("Update", &discussion).Return(nil)
		err := useCase.Update(&discussion)
		assert.Nil(t, err)
	})

	t.Run("comment cannot be empty", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussion := entities.Discussion{
			Comment: "",
		}

		err := useCase.Update(&discussion)
		assert.NotNil(t, err)
		assert.Equal(t, constants.ErrCommentCannotBeEmpty, err)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussion := entities.Discussion{
			Comment: "Hello",
		}

		mockDiscussion.On("Update", &discussion).Return(constants.ErrInternalServerError)
		err := useCase.Update(&discussion)
		assert.NotNil(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})
}

func TestDiscussionUseCase_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		mockDiscussion.On("Delete", 1).Return(nil)
		err := useCase.Delete(1)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		mockDiscussion.On("Delete", 1).Return(constants.ErrInternalServerError)
		err := useCase.Delete(1)
		assert.NotNil(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})

}

func TestDiscussionUseCase_GetByComplaintID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		discussions := []entities.Discussion{
			{
				ID:          1,
				ComplaintID: "1",
				Comment:     "Hello",
			},
		}

		mockDiscussion.On("GetByComplaintID", "1").Return(&discussions, nil)
		result, err := useCase.GetByComplaintID("1")
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("discussion not found", func(t *testing.T) {
		mockDiscussion := new(MockDiscussion)
		mockFaq := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussion, mockFaq, mockOpenAIAPI)

		mockDiscussion.On("GetByComplaintID", "1").Return((*[]entities.Discussion)(nil), constants.ErrDiscussionNotFound)
		result, err := useCase.GetByComplaintID("1")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestDiscussionUseCase_GetAnswerRecommendation(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDiscussionRepo := new(MockDiscussion)
		mockFaqRepo := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussionRepo, mockFaqRepo, mockOpenAIAPI)

		discussions := []entities.Discussion{
			{
				ID:          1,
				ComplaintID: "1",
				Comment:     "Hello",
			},
		}

		faqs := []entities.Faq{
			{
				Question: "What is the meaning of life?",
				Answer:   "42",
			},
		}

		mockDiscussionRepo.On("GetByComplaintID", "1").Return(&discussions, nil)
		mockFaqRepo.On("GetAll").Return(faqs, nil)
		mockOpenAIAPI.On("GetChatCompletion", mock.Anything, "").Return("Test response", nil)

		result, err := useCase.GetAnswerRecommendation("1")
		assert.Nil(t, err)
		assert.Equal(t, "Test response", result)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscussionRepo := new(MockDiscussion)
		mockFaqRepo := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussionRepo, mockFaqRepo, mockOpenAIAPI)

		mockDiscussionRepo.On("GetByComplaintID", "1").Return((*[]entities.Discussion)(nil), constants.ErrDiscussionNotFound)

		result, err := useCase.GetAnswerRecommendation("1")
		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("GetAll returns error", func(t *testing.T) {
		mockDiscussionRepo := new(MockDiscussion)
		mockFaqRepo := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussionRepo, mockFaqRepo, mockOpenAIAPI)
		mockDiscussionRepo.On("GetByComplaintID", "1").Return((*[]entities.Discussion)(nil), nil)
		mockFaqRepo.On("GetAll").Return([]entities.Faq(nil), errors.New("some error"))
		_, err := useCase.GetAnswerRecommendation("1")
		assert.Error(t, err)
	})

	t.Run("GetChatCompletion returns error", func(t *testing.T) {
		mockDiscussionRepo := new(MockDiscussion)
		mockFaqRepo := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussionRepo, mockFaqRepo, mockOpenAIAPI)
		mockDiscussionRepo.On("GetByComplaintID", "1").Return((*[]entities.Discussion)(nil), nil)
		mockFaqRepo.On("GetAll").Return([]entities.Faq{{Question: "What is the meaning of life?", Answer: "42"}}, nil)
		mockOpenAIAPI.On("GetChatCompletion", mock.Anything, "").Return("", errors.New("some error"))
		_, err := useCase.GetAnswerRecommendation("1")
		assert.Error(t, err)
	})

	t.Run("UserID is not nil", func(t *testing.T) {
		mockDiscussionRepo := new(MockDiscussion)
		mockFaqRepo := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussionRepo, mockFaqRepo, mockOpenAIAPI)

		userID := 1
		discussions := []entities.Discussion{
			{
				ID:      1,
				UserID:  &userID,
				Comment: "Hello",
			},
		}

		mockDiscussionRepo.On("GetByComplaintID", "1").Return(&discussions, nil)
		mockFaqRepo.On("GetAll").Return([]entities.Faq{}, nil)
		mockOpenAIAPI.On("GetChatCompletion", mock.Anything, "").Return("1.)User: Hello\nTest response", nil)

		result, err := useCase.GetAnswerRecommendation("1")
		assert.Nil(t, err)
		assert.Contains(t, result, "1.)User: Hello")
	})

	t.Run("UserID is nil", func(t *testing.T) {
		mockDiscussionRepo := new(MockDiscussion)
		mockFaqRepo := new(MockFaq)
		mockOpenAIAPI := new(OpenAIAPI)
		useCase := NewDiscussionUseCase(mockDiscussionRepo, mockFaqRepo, mockOpenAIAPI)

		discussions := []entities.Discussion{
			{
				ID:      1,
				UserID:  nil,
				Comment: "Hello",
			},
		}

		mockDiscussionRepo.On("GetByComplaintID", "1").Return(&discussions, nil)
		mockFaqRepo.On("GetAll").Return([]entities.Faq{}, nil)
		mockOpenAIAPI.On("GetChatCompletion", mock.Anything, "").Return("Test response", nil)

		result, err := useCase.GetAnswerRecommendation("1")
		assert.Nil(t, err)
		assert.NotContains(t, result, "1.)User: Hello")
	})

}
