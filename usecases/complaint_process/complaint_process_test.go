package complaint_process

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockComplaintProcess struct {
	mock.Mock
}

func (m *MockComplaintProcess) Create(complaintProcesses *entities.ComplaintProcess) error {
	args := m.Called(complaintProcesses)
	return args.Error(0)
}

func (m *MockComplaintProcess) GetByComplaintID(complaintID string) ([]entities.ComplaintProcess, error) {
	args := m.Called(complaintID)
	result := args.Get(0)
	return result.([]entities.ComplaintProcess), args.Error(1)
}

func (m *MockComplaintProcess) Update(complaintProcesses *entities.ComplaintProcess) error {
	args := m.Called(complaintProcesses)
	return args.Error(0)
}

func (m *MockComplaintProcess) Delete(complaintID string, complaintProcessID int) (string, error) {
	args := m.Called(complaintID, complaintProcessID)
	return args.String(0), args.Error(1)
}

type MockComplaint struct {
	mock.Mock
}

func (m *MockComplaint) GetByUserID(userId int) ([]entities.Complaint, error) {
	args := m.Called(userId)
	result := args.Get(0)
	return result.([]entities.Complaint), args.Error(1)

}

func (m *MockComplaint) Create(complaint *entities.Complaint) error {
	args := m.Called(complaint)
	return args.Error(0)

}

func (m *MockComplaint) Delete(id string, userId int) error {
	args := m.Called(id, userId)
	return args.Error(0)

}

func (m *MockComplaint) AdminDelete(id string) error {
	args := m.Called(id)
	return args.Error(0)

}

func (m *MockComplaint) Update(complaint entities.Complaint) (entities.Complaint, error) {
	args := m.Called(complaint)
	result := args.Get(0)
	return result.(entities.Complaint), args.Error(1)

}

func (m *MockComplaint) UpdateStatus(id string, status string) error {
	args := m.Called(id, status)
	return args.Error(0)

}

func (m *MockComplaint) GetStatus(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)

}

func (m *MockComplaint) Import(complaints []entities.Complaint) error {
	args := m.Called(complaints)
	return args.Error(0)

}

func (m *MockComplaint) IncreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)

}

func (m *MockComplaint) DecreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockComplaint) GetComplaintIDsByUserID(userID int) ([]string, error) {
	args := m.Called(userID)
	result := args.Get(0)
	return result.([]string), args.Error(1)

}

func (m *MockComplaint) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	args := m.Called(limit, page, search, filter, sortBy, sortType)
	result := args.Get(0)
	return result.([]entities.Complaint), args.Error(1)
}

func (m *MockComplaint) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	args := m.Called(limit, page, search, filter)
	result := args.Get(0)
	return result.(entities.Metadata), args.Error(1)
}

func (m *MockComplaint) GetByID(id string) (entities.Complaint, error) {
	args := m.Called(id)
	result := args.Get(0)
	return result.(entities.Complaint), args.Error(1)
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Pending",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Pending", nil)
		mockComplaintProcessRepo.On("Create", mock.Anything).Return(nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Pending",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("", constants.ErrInternalServerError)
		mockComplaintProcessRepo.On("Create", mock.Anything).Return(constants.ErrInternalServerError)
		result, err := usecase.Create(dummyComplaintProcess)
		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})

	t.Run("error when message is empty", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "",
			Status:      "Pending",
			ComplaintID: "123",
		}

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)
	})

	t.Run("error when status is invalid", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Invalid",
			ComplaintID: "123",
		}

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrInvalidStatus, err)
	})

	t.Run("error when status is Pending and complaint status is On Progress", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Pending",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("On Progress", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintNotVerified, err)
	})

	t.Run("error when status is Pending and complaint status is Selesai", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Pending",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Selesai", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintNotVerified, err)
	})

	t.Run("error when complaint not found in repository", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Pending",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Pending", nil)
		mockComplaintProcessRepo.On("Create", mock.Anything).Return(errors.New("REFERENCES `complaints` (`id`)"))

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintNotFound, err)
	})

	t.Run("error when internal server error occurs in repository", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Pending",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Pending", nil)
		mockComplaintProcessRepo.On("Create", mock.Anything).Return(errors.New("some other error"))

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})

	t.Run("error when status is Verifikasi and complaint status is Verifikasi", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Verifikasi",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Verifikasi", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyVerified, err)
	})

	t.Run("error when status is Verifikasi and complaint status is Ditolak", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Verifikasi",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Ditolak", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyRejected, err)
	})

	t.Run("error when status is Verifikasi and complaint status is Selesai", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Verifikasi",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Selesai", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyFinished, err)
	})

	t.Run("error when status is Verifikasi and complaint status is On Progress", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Verifikasi",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("On Progress", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyOnProgress, err)
	})

	t.Run("error when status is On Progress and complaint status is On Progress", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "On Progress",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("On Progress", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyOnProgress, err)
	})

	t.Run("error when status is On Progress and complaint status is Ditolak", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "On Progress",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Ditolak", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyRejected, err)

	})

	t.Run("error when status is On Progress and complaint status is Selesai", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "On Progress",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Selesai", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyFinished, err)
	})

	t.Run("error when status is On Progress and complaint status is Pending", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "On Progress",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Pending", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintNotVerified, err)
	})

	t.Run("error when status is Selesai and complaint status is Selesai", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Selesai",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Selesai", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyFinished, err)
	})

	t.Run("error when status is Selesai and complaint status is Ditolak", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Selesai",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Ditolak", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyRejected, err)
	})

	t.Run("error when status is Selesai and complaint status is Pending", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Selesai",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Pending", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintNotVerified, err)
	})

	t.Run("error when status is Selesai and complaint status is Verifikasi", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Selesai",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Verifikasi", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintNotOnProgress, err)
	})

	t.Run("error when status is Ditolak and complaint status is Ditolak", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Ditolak",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Ditolak", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyRejected, err)
	})

	t.Run("error when status is Ditolak and complaint status is Selesai", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Ditolak",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Selesai", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyFinished, err)
	})

	t.Run("error when status is Ditolak and complaint status is On Progress", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Ditolak",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("On Progress", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyOnProgress, err)
	})

	t.Run("error when status is Ditolak and complaint status is Verifikasi", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)
		dummyComplaintProcess := &entities.ComplaintProcess{
			Message:     "Test Message",
			Status:      "Ditolak",
			ComplaintID: "123",
		}

		mockComplaintRepo.On("GetStatus", mock.Anything).Return("Verifikasi", nil)

		result, err := usecase.Create(dummyComplaintProcess)

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrComplaintAlreadyVerified, err)
	})

}
func TestComplaintProcessUseCase_GetByComplaintID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)

		mockComplaintProcessRepo.On("GetByComplaintID", mock.Anything).Return([]entities.ComplaintProcess{}, nil)

		result, err := usecase.GetByComplaintID("123")

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)

		mockComplaintProcessRepo.On("GetByComplaintID", mock.Anything).Return([]entities.ComplaintProcess{}, constants.ErrInternalServerError)

		result, err := usecase.GetByComplaintID("123")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, constants.ErrInternalServerError, err)
	})

	t.Run("complaint process not found", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)

		mockComplaintProcessRepo.On("GetByComplaintID", mock.Anything).Return([]entities.ComplaintProcess{}, constants.ErrComplaintProcessNotFound)

		result, err := usecase.GetByComplaintID("123")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, constants.ErrComplaintProcessNotFound, err)
	})

}

func TestComplaintProcessUseCase_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)

		mockComplaintProcessRepo.On("Update", mock.Anything).Return(nil)

		result, err := usecase.Update(&entities.ComplaintProcess{
			ID:          1,
			ComplaintID: "123",
			Message:     "Test Message",
			Status:      "Pending",
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("error when message is empty", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)

		result, err := usecase.Update(&entities.ComplaintProcess{
			ID:          1,
			ComplaintID: "123",
			Message:     "",
			Status:      "Pending",
		})

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)
	})

	t.Run("error when repository update fails", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		mockComplaintRepo := new(MockComplaint)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, mockComplaintRepo)

		mockComplaintProcessRepo.On("Update", mock.Anything).Return(errors.New("update error"))

		result, err := usecase.Update(&entities.ComplaintProcess{
			ID:          1,
			ComplaintID: "123",
			Message:     "Test Message",
			Status:      "Pending",
		})

		assert.Error(t, err)
		assert.Equal(t, entities.ComplaintProcess{}, result)
		assert.Equal(t, "update error", err.Error())
	})

}

func TestComplaintProcessUseCase_Delete(t *testing.T) {
	t.Run("error when complaintID is empty", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, nil)

		status, err := usecase.Delete("", 1)

		assert.Error(t, err)
		assert.Equal(t, "", status)
		assert.Equal(t, constants.ErrInvalidIDFormat, err)
	})

	t.Run("error when complaintProcessID is zero", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, nil)

		status, err := usecase.Delete("123", 0)

		assert.Error(t, err)
		assert.Equal(t, "", status)
		assert.Equal(t, constants.ErrInvalidIDFormat, err)
	})

	t.Run("error when repository delete fails", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, nil)

		mockComplaintProcessRepo.On("Delete", mock.Anything, mock.Anything).Return("", errors.New("delete error"))

		status, err := usecase.Delete("123", 1)

		assert.Error(t, err)
		assert.Equal(t, "", status)
		assert.Equal(t, "delete error", err.Error())
	})

	t.Run("success with status Verifikasi", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, nil)

		mockComplaintProcessRepo.On("Delete", mock.Anything, mock.Anything).Return("Verifikasi", nil)

		status, err := usecase.Delete("123", 1)

		assert.NoError(t, err)
		assert.Equal(t, "Pending", status)
	})

	t.Run("success with status On Progress", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, nil)

		mockComplaintProcessRepo.On("Delete", mock.Anything, mock.Anything).Return("On Progress", nil)

		status, err := usecase.Delete("123", 1)

		assert.NoError(t, err)
		assert.Equal(t, "Verifikasi", status)
	})

	t.Run("success with status Selesai", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, nil)

		mockComplaintProcessRepo.On("Delete", mock.Anything, mock.Anything).Return("Selesai", nil)

		status, err := usecase.Delete("123", 1)

		assert.NoError(t, err)
		assert.Equal(t, "On Progress", status)
	})

	t.Run("success with status Ditolak", func(t *testing.T) {
		mockComplaintProcessRepo := new(MockComplaintProcess)
		usecase := NewComplaintProcessUseCase(mockComplaintProcessRepo, nil)

		mockComplaintProcessRepo.On("Delete", mock.Anything, mock.Anything).Return("Ditolak", nil)

		status, err := usecase.Delete("123", 1)

		assert.NoError(t, err)
		assert.Equal(t, "Pending", status)
	})
}
