package complaint_activity

import (
	"e-complaint-api/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ComplaintActivityRepository struct {
	mock.Mock
}

func (m *ComplaintActivityRepository) GetByComplaintIDs(complaintIDs []string, activityType string) ([]entities.ComplaintActivity, error) {
	args := m.Called(complaintIDs, activityType)
	return args.Get(0).([]entities.ComplaintActivity), args.Error(1)
}

func (m *ComplaintActivityRepository) Create(complaintActivity *entities.ComplaintActivity) error {
	args := m.Called(complaintActivity)
	return args.Error(0)
}

func (m *ComplaintActivityRepository) Delete(complaintActivity entities.ComplaintActivity) error {
	args := m.Called(complaintActivity)
	return args.Error(0)
}

func (m *ComplaintActivityRepository) Update(complaintActivity entities.ComplaintActivity) error {
	args := m.Called(complaintActivity)
	return args.Error(0)
}

func TestGetByComplaintIDs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := []entities.ComplaintActivity{
			{
				ID:           1,
				ComplaintID:  "1",
				LikeID:       nil,
				DiscussionID: nil,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
		}

		mockRepo.On("GetByComplaintIDs", []string{"1"}, "").Return(complaintActivity, nil)

		u := NewComplaintActivityUseCase(mockRepo)
		res, err := u.GetByComplaintIDs([]string{"1"}, "")
		assert.NoError(t, err)
		assert.Equal(t, complaintActivity, res)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		mockRepo.On("GetByComplaintIDs", []string{"1"}, "").Return(([]entities.ComplaintActivity)(nil), assert.AnError)

		u := NewComplaintActivityUseCase(mockRepo)
		res, err := u.GetByComplaintIDs([]string{"1"}, "")
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := &entities.ComplaintActivity{
			ID:           1,
			ComplaintID:  "1",
			LikeID:       nil,
			DiscussionID: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockRepo.On("Create", complaintActivity).Return(nil)

		u := NewComplaintActivityUseCase(mockRepo)
		res, err := u.Create(complaintActivity)
		assert.NoError(t, err)
		assert.Equal(t, *complaintActivity, res)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := &entities.ComplaintActivity{
			ID:           1,
			ComplaintID:  "1",
			LikeID:       nil,
			DiscussionID: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockRepo.On("Create", complaintActivity).Return(assert.AnError)

		u := NewComplaintActivityUseCase(mockRepo)
		res, err := u.Create(complaintActivity)
		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintActivity{}, res)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := &entities.ComplaintActivity{
			ID:           1,
			ComplaintID:  "1",
			LikeID:       nil,
			DiscussionID: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockRepo.On("Create", complaintActivity).Return(assert.AnError)

		u := NewComplaintActivityUseCase(mockRepo)
		res, err := u.Create(complaintActivity)
		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintActivity{}, res)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := entities.ComplaintActivity{
			ID:           1,
			ComplaintID:  "1",
			LikeID:       nil,
			DiscussionID: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockRepo.On("Delete", complaintActivity).Return(nil)

		u := NewComplaintActivityUseCase(mockRepo)
		err := u.Delete(complaintActivity)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := entities.ComplaintActivity{
			ID:           1,
			ComplaintID:  "1",
			LikeID:       nil,
			DiscussionID: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockRepo.On("Delete", complaintActivity).Return(assert.AnError)

		u := NewComplaintActivityUseCase(mockRepo)
		err := u.Delete(complaintActivity)
		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := entities.ComplaintActivity{
			ID:           1,
			ComplaintID:  "1",
			LikeID:       nil,
			DiscussionID: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockRepo.On("Update", complaintActivity).Return(nil)

		u := NewComplaintActivityUseCase(mockRepo)
		err := u.Update(complaintActivity)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(ComplaintActivityRepository)
		complaintActivity := entities.ComplaintActivity{
			ID:           1,
			ComplaintID:  "1",
			LikeID:       nil,
			DiscussionID: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockRepo.On("Update", complaintActivity).Return(assert.AnError)

		u := NewComplaintActivityUseCase(mockRepo)
		err := u.Update(complaintActivity)
		assert.Error(t, err)
	})
}
