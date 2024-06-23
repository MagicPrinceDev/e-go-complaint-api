package complaint_like

import (
	"e-complaint-api/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockComplaintLike struct {
	mock.Mock
}

func (m MockComplaintLike) Unlike(complaintLike *entities.ComplaintLike) error {
	args := m.Called(complaintLike)
	return args.Error(0)

}

func (m MockComplaintLike) Likes(complaintLike *entities.ComplaintLike) error {
	args := m.Called(complaintLike)
	return args.Error(0)

}

func (m MockComplaintLike) FindByUserAndComplaint(userID int, complaintID string) (*entities.ComplaintLike, error) {
	args := m.Called(userID, complaintID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ComplaintLike), args.Error(1)
}

func TestComplaintLikeUseCase_ToggleLike(t *testing.T) {
	t.Run("success - like", func(t *testing.T) {
		mockRepo := new(MockComplaintLike)
		useCase := NewComplaintLikeUseCase(mockRepo)

		complaintLike := &entities.ComplaintLike{
			UserID:      1,
			ComplaintID: "C-123j9ak280",
		}

		mockRepo.On("FindByUserAndComplaint", complaintLike.UserID, complaintLike.ComplaintID).Return(nil, nil)
		mockRepo.On("Likes", complaintLike).Return(nil)

		status, err := useCase.ToggleLike(complaintLike)
		assert.Nil(t, err)
		assert.Equal(t, "liked", status)
	})

	t.Run("success - unlike", func(t *testing.T) {
		mockRepo := new(MockComplaintLike)
		useCase := NewComplaintLikeUseCase(mockRepo)

		complaintLike := &entities.ComplaintLike{
			UserID:      1,
			ComplaintID: "1",
		}

		existingComplaintLike := &entities.ComplaintLike{
			UserID:      1,
			ComplaintID: "C-123j9ak280",
		}

		mockRepo.On("FindByUserAndComplaint", complaintLike.UserID, complaintLike.ComplaintID).Return(existingComplaintLike, nil)
		mockRepo.On("Unlike", existingComplaintLike).Return(nil)

		status, err := useCase.ToggleLike(complaintLike)
		assert.Nil(t, err)
		assert.Equal(t, "unliked", status)
	})

	t.Run("error - FindByUserAndComplaint", func(t *testing.T) {
		mockRepo := new(MockComplaintLike)
		useCase := NewComplaintLikeUseCase(mockRepo)

		complaintLike := &entities.ComplaintLike{
			UserID:      1,
			ComplaintID: "C-123j9ak280",
		}

		mockRepo.On("FindByUserAndComplaint", complaintLike.UserID, complaintLike.ComplaintID).Return(nil, errors.New("error"))

		_, err := useCase.ToggleLike(complaintLike)
		assert.NotNil(t, err)
	})

	t.Run("error - Likes", func(t *testing.T) {
		mockRepo := new(MockComplaintLike)
		useCase := NewComplaintLikeUseCase(mockRepo)

		complaintLike := &entities.ComplaintLike{
			UserID:      1,
			ComplaintID: "C-123j9ak280",
		}

		mockRepo.On("FindByUserAndComplaint", complaintLike.UserID, complaintLike.ComplaintID).Return(nil, nil)
		mockRepo.On("Likes", complaintLike).Return(errors.New("error"))

		_, err := useCase.ToggleLike(complaintLike)
		assert.NotNil(t, err)
	})

	t.Run("error - Unlike", func(t *testing.T) {
		mockRepo := new(MockComplaintLike)
		useCase := NewComplaintLikeUseCase(mockRepo)

		complaintLike := &entities.ComplaintLike{
			UserID:      1,
			ComplaintID: "C-123j9ak280",
		}

		existingComplaintLike := &entities.ComplaintLike{
			UserID:      1,
			ComplaintID: "C-123j9ak280",
		}

		mockRepo.On("FindByUserAndComplaint", complaintLike.UserID, complaintLike.ComplaintID).Return(existingComplaintLike, nil)
		mockRepo.On("Unlike", existingComplaintLike).Return(errors.New("error"))

		_, err := useCase.ToggleLike(complaintLike)
		assert.NotNil(t, err)
	})

}
