package chatbot

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type Chatbot struct {
	mock.Mock
}

func (m *Chatbot) Create(chatbot *entities.Chatbot) error {
	args := m.Called(chatbot)
	return args.Error(0)

}

func (m *Chatbot) ClearHistory(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *Chatbot) GetChatCompletion(chatbot *Chatbot) error {
	args := m.Called(chatbot)
	return args.Error(0)
}

func (m *Chatbot) GetHistory(userID int) ([]entities.Chatbot, error) {
	args := m.Called(userID)
	return args.Get(0).([]entities.Chatbot), args.Error(1)
}

func (m *Chatbot) GetChatbotByID(id int) (*entities.Chatbot, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Chatbot), args.Error(1)
}

type Faq struct {
	mock.Mock
}

func (m *Faq) GetAll() ([]entities.Faq, error) {
	args := m.Called()
	return args.Get(0).([]entities.Faq), args.Error(1)
}

type Complaint struct {
	mock.Mock
}

func (m *Complaint) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	args := m.Called(limit, page, search, filter, sortBy, sortType)
	return args.Get(0).([]entities.Complaint), args.Error(1)
}

func (m *Complaint) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	args := m.Called(limit, page, search, filter)
	return args.Get(0).(entities.Metadata), args.Error(1)
}

func (m *Complaint) GetByID(id string) (entities.Complaint, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Complaint), args.Error(1)
}

func (m *Complaint) GetByUserID(userId int) ([]entities.Complaint, error) {
	args := m.Called(userId)
	return args.Get(0).([]entities.Complaint), args.Error(1)
}

func (m *Complaint) Create(complaint *entities.Complaint) error {
	args := m.Called(complaint)
	return args.Error(0)
}

func (m *Complaint) Delete(id string, userId int) error {
	args := m.Called(id, userId)
	return args.Error(0)
}

func (m *Complaint) AdminDelete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *Complaint) Update(complaint entities.Complaint) (entities.Complaint, error) {
	args := m.Called(complaint)
	return args.Get(0).(entities.Complaint), args.Error(1)
}

func (m *Complaint) UpdateStatus(id string, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *Complaint) GetStatus(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *Complaint) Import(complaints []entities.Complaint) error {
	args := m.Called(complaints)
	return args.Error(0)
}

func (m *Complaint) IncreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *Complaint) DecreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *Complaint) GetComplaintIDsByUserID(userId int) ([]string, error) {
	args := m.Called(userId)
	return args.Get(0).([]string), args.Error(1)
}

type OpenAIAPI struct {
	mock.Mock
}

func (m *OpenAIAPI) GetChatCompletion(prompt []string, userMessage string) (string, error) {
	args := m.Called(prompt, userMessage)
	return args.String(0), args.Error(1)
}

func TestChatbotUseCase_ClearHistory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		chatbotRepo.On("ClearHistory", 1).Return(nil)

		uc := NewChatbotUseCase(chatbotRepo, nil, nil, nil)

		err := uc.ClearHistory(1)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		chatbotRepo.On("ClearHistory", 1).Return(constants.ErrInternalServerError)

		uc := NewChatbotUseCase(chatbotRepo, nil, nil, nil)

		err := uc.ClearHistory(1)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})
}

func TestChatbotUseCase_GetChatCompletion(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		faqRepo := new(Faq)
		complaintRepo := new(Complaint)
		openAIAPI := new(OpenAIAPI)

		chatbot := &entities.Chatbot{UserID: 1, UserMessage: "Hello"}

		faqRepo.On("GetAll").Return([]entities.Faq{}, nil)
		complaintRepo.On("GetByUserID", chatbot.UserID).Return([]entities.Complaint{}, nil)
		openAIAPI.On("GetChatCompletion", mock.Anything, chatbot.UserMessage).Return("Hello, how can I assist you?", nil)
		chatbotRepo.On("Create", chatbot).Return(nil)

		uc := NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)

		err := uc.GetChatCompletion(chatbot)
		assert.Nil(t, err)
		assert.Equal(t, "Hello, how can I assist you?", chatbot.BotResponse)
	})

	t.Run("success with user complaints", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		faqRepo := new(Faq)
		complaintRepo := new(Complaint)
		openAIAPI := new(OpenAIAPI)

		chatbot := &entities.Chatbot{UserID: 1, UserMessage: "Hello"}

		faqRepo.On("GetAll").Return([]entities.Faq{}, nil)
		complaintRepo.On("GetByUserID", chatbot.UserID).Return([]entities.Complaint{{ID: "1", Description: "Test", Status: "Open", CreatedAt: time.Now()}}, nil)
		openAIAPI.On("GetChatCompletion", mock.Anything, chatbot.UserMessage).Return("Hello, how can I assist you?", nil)
		chatbotRepo.On("Create", chatbot).Return(nil)

		uc := NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)

		err := uc.GetChatCompletion(chatbot)
		assert.Nil(t, err)
		assert.Equal(t, "Hello, how can I assist you?", chatbot.BotResponse)
	})

	t.Run("success with FAQs", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		faqRepo := new(Faq)
		complaintRepo := new(Complaint)
		openAIAPI := new(OpenAIAPI)

		chatbot := &entities.Chatbot{UserID: 1, UserMessage: "Hello"}

		faqRepo.On("GetAll").Return([]entities.Faq{{Question: "Test", Answer: "Test"}}, nil)
		complaintRepo.On("GetByUserID", chatbot.UserID).Return([]entities.Complaint{}, nil)
		openAIAPI.On("GetChatCompletion", mock.Anything, chatbot.UserMessage).Return("Hello, how can I assist you?", nil)
		chatbotRepo.On("Create", chatbot).Return(nil)

		uc := NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)

		err := uc.GetChatCompletion(chatbot)
		assert.Nil(t, err)
		assert.Equal(t, "Hello, how can I assist you?", chatbot.BotResponse)
	})

	t.Run("error", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		faqRepo := new(Faq)
		complaintRepo := new(Complaint)
		openAIAPI := new(OpenAIAPI)

		chatbot := &entities.Chatbot{UserID: 1, UserMessage: "Hello"}

		faqRepo.On("GetAll").Return([]entities.Faq{}, errors.New("error"))
		uc := NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)

		err := uc.GetChatCompletion(chatbot)
		assert.NotNil(t, err)
	})

	t.Run("error on GetChatCompletion", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		faqRepo := new(Faq)
		complaintRepo := new(Complaint)
		openAIAPI := new(OpenAIAPI)

		chatbot := &entities.Chatbot{UserID: 1, UserMessage: "Hello"}

		faqRepo.On("GetAll").Return([]entities.Faq{}, nil)
		complaintRepo.On("GetByUserID", chatbot.UserID).Return([]entities.Complaint{}, nil)
		openAIAPI.On("GetChatCompletion", mock.Anything, chatbot.UserMessage).Return("", errors.New("error"))

		uc := NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)

		err := uc.GetChatCompletion(chatbot)
		assert.NotNil(t, err)
	})

	t.Run("error on Create", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		faqRepo := new(Faq)
		complaintRepo := new(Complaint)
		openAIAPI := new(OpenAIAPI)

		chatbot := &entities.Chatbot{UserID: 1, UserMessage: "Hello"}

		faqRepo.On("GetAll").Return([]entities.Faq{}, nil)
		complaintRepo.On("GetByUserID", chatbot.UserID).Return([]entities.Complaint{}, nil)
		openAIAPI.On("GetChatCompletion", mock.Anything, chatbot.UserMessage).Return("Hello, how can I assist you?", nil)
		chatbotRepo.On("Create", chatbot).Return(errors.New("error"))

		uc := NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)

		err := uc.GetChatCompletion(chatbot)
		assert.NotNil(t, err)
	})

	t.Run("error_on_GetByUserID", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		faqRepo := new(Faq)
		complaintRepo := new(Complaint)
		openAIAPI := new(OpenAIAPI)

		chatbot := &entities.Chatbot{UserID: 1, UserMessage: "Hello"}

		faqRepo.On("GetAll").Return([]entities.Faq{}, nil)
		complaintRepo.On("GetByUserID", chatbot.UserID).Return([]entities.Complaint{}, errors.New("error"))

		uc := NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)

		err := uc.GetChatCompletion(chatbot)
		assert.NotNil(t, err)
	})

}

func TestChatbotUseCase_GetHistory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		chatbotRepo.On("GetHistory", 1).Return([]entities.Chatbot{}, nil)

		uc := NewChatbotUseCase(chatbotRepo, nil, nil, nil)

		chatbots, err := uc.GetHistory(1)
		assert.Nil(t, err)
		assert.Equal(t, []entities.Chatbot{}, chatbots)
	})

	t.Run("error", func(t *testing.T) {
		chatbotRepo := new(Chatbot)
		chatbotRepo.On("GetHistory", 1).Return([]entities.Chatbot{}, constants.ErrInternalServerError)

		uc := NewChatbotUseCase(chatbotRepo, nil, nil, nil)

		chatbots, err := uc.GetHistory(1)
		assert.NotNil(t, err)
		assert.Nil(t, chatbots)
	})

}
