package complaint

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"mime/multipart"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockComplaintRepo struct {
	mock.Mock
}

func (m *MockComplaintRepo) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	args := m.Called(limit, page, search, filter, sortBy, sortType)
	return args.Get(0).([]entities.Complaint), args.Error(1)
}

func (m *MockComplaintRepo) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	args := m.Called(limit, page, search, filter)
	return args.Get(0).(entities.Metadata), args.Error(1)
}

func (m *MockComplaintRepo) GetByID(id string) (entities.Complaint, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Complaint), args.Error(1)
}

func (m *MockComplaintRepo) GetByUserID(userId int) ([]entities.Complaint, error) {
	args := m.Called(userId)
	return args.Get(0).([]entities.Complaint), args.Error(1)
}

func (m *MockComplaintRepo) Create(complaint *entities.Complaint) error {
	args := m.Called(complaint)
	return args.Error(0)
}

func (m *MockComplaintRepo) Delete(id string, userId int) error {
	args := m.Called(id, userId)
	return args.Error(0)
}

func (m *MockComplaintRepo) AdminDelete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockComplaintRepo) Update(complaint entities.Complaint) (entities.Complaint, error) {
	args := m.Called(complaint)
	return args.Get(0).(entities.Complaint), args.Error(1)
}

func (m *MockComplaintRepo) UpdateStatus(id string, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockComplaintRepo) GetStatus(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *MockComplaintRepo) Import(complaints []entities.Complaint) error {
	args := m.Called(complaints)
	return args.Error(0)
}

func (m *MockComplaintRepo) IncreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockComplaintRepo) DecreaseTotalLikes(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockComplaintRepo) GetComplaintIDsByUserID(userId int) ([]string, error) {
	args := m.Called(userId)
	return args.Get(0).([]string), args.Error(1)
}

type MockComplaintFileRepo struct {
	mock.Mock
}

func (m *MockComplaintFileRepo) Create(complaintFiles []*entities.ComplaintFile) error {
	args := m.Called(complaintFiles)
	return args.Error(0)
}

func (m *MockComplaintFileRepo) DeleteByComplaintID(complaintID string) error {
	args := m.Called(complaintID)
	return args.Error(0)
}

type MockUtils struct {
	mock.Mock
}

func (m *MockUtils) GetRowsFromExcel(file *multipart.FileHeader) ([][]string, error) {
	args := m.Called(file)
	return args.Get(0).([][]string), args.Error(1)
}

func TestGetPaginated(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "desc").Return([]entities.Complaint{}, nil)

		result, err := mockUsecase.GetPaginated(10, 1, "", map[string]interface{}{}, "created_at", "desc")
		assert.NoError(t, err)
		assert.Equal(t, []entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("success with empty sortBy and sortType", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "DESC").Return([]entities.Complaint{}, nil)

		result, err := mockUsecase.GetPaginated(10, 1, "", map[string]interface{}{}, "", "")
		assert.NoError(t, err)
		assert.Equal(t, []entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("failed limit must filled when page is filled", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.GetPaginated(0, 1, "", map[string]interface{}{}, "created_at", "desc")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrLimitMustBeFilled, err)
		assert.Nil(t, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed page must filled when limit is filled", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.GetPaginated(10, 0, "", map[string]interface{}{}, "created_at", "desc")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrPageMustBeFilled, err)
		assert.Nil(t, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetPaginated", 10, 1, "", map[string]interface{}{}, "created_at", "desc").Return(([]entities.Complaint)(nil), constants.ErrInternalServerError)

		result, err := mockUsecase.GetPaginated(10, 1, "", map[string]interface{}{}, "created_at", "desc")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Nil(t, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

func TestGetMetaData(t *testing.T) {
	t.Run("success empty", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetMetaData", 10, 1, "", map[string]interface{}{}).Return(entities.Metadata{}, nil)

		result, err := mockUsecase.GetMetaData(10, 1, "", map[string]interface{}{})
		expectedMetadata := entities.Metadata{
			Pagination: entities.Pagination{
				FirstPage:   1,
				CurrentPage: 1,
				LastPage:    1,
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedMetadata, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("success not empty", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetMetaData", 10, 1, "", map[string]interface{}{}).Return(entities.Metadata{
			TotalData: 10,
		}, nil)

		result, err := mockUsecase.GetMetaData(10, 1, "", map[string]interface{}{})
		expectedMetadata := entities.Metadata{
			TotalData: 10,
			Pagination: entities.Pagination{
				FirstPage:        1,
				LastPage:         1,
				CurrentPage:      1,
				TotalDataPerPage: 10,
				PrevPage:         0,
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedMetadata, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("success not empty with page > 1", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetMetaData", 10, 2, "", map[string]interface{}{}).Return(entities.Metadata{
			TotalData: 10,
		}, nil)

		result, err := mockUsecase.GetMetaData(10, 2, "", map[string]interface{}{})
		expectedMetadata := entities.Metadata{
			TotalData: 10,
			Pagination: entities.Pagination{
				FirstPage:        1,
				LastPage:         1,
				CurrentPage:      2,
				TotalDataPerPage: 10,
				PrevPage:         1,
				NextPage:         0,
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedMetadata, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("success not empty with page > 1 and not in last page", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetMetaData", 10, 2, "", map[string]interface{}{}).Return(entities.Metadata{
			TotalData: 30,
		}, nil)

		result, err := mockUsecase.GetMetaData(10, 2, "", map[string]interface{}{})
		expectedMetadata := entities.Metadata{
			TotalData: 30,
			Pagination: entities.Pagination{
				FirstPage:        1,
				LastPage:         3,
				CurrentPage:      2,
				TotalDataPerPage: 10,
				PrevPage:         1,
				NextPage:         3,
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedMetadata, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("success without limit and page filled", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetMetaData", 0, 0, "", map[string]interface{}{}).Return(entities.Metadata{}, nil)

		result, err := mockUsecase.GetMetaData(0, 0, "", map[string]interface{}{})
		expectedMetadata := entities.Metadata{
			Pagination: entities.Pagination{
				FirstPage:   1,
				CurrentPage: 1,
				LastPage:    1,
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedMetadata, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetMetaData", 10, 1, "", map[string]interface{}{}).Return(entities.Metadata{}, constants.ErrInternalServerError)

		result, err := mockUsecase.GetMetaData(10, 1, "", map[string]interface{}{})
		expectedMetadata := entities.Metadata{}

		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, expectedMetadata, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetByID", "1").Return(entities.Complaint{}, nil)

		result, err := mockUsecase.GetByID("1")
		assert.NoError(t, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetByID", "1").Return(entities.Complaint{}, constants.ErrInternalServerError)

		result, err := mockUsecase.GetByID("1")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed complaint not found", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetByID", "1").Return(entities.Complaint{}, constants.ErrComplaintNotFound)

		result, err := mockUsecase.GetByID("1")
		assert.Error(t, err)
		assert.Equal(t, constants.ErrComplaintNotFound, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

func TestGetByUserID(t *testing.T) {
	t.Run("success empty", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetByUserID", 1).Return([]entities.Complaint{}, nil)

		result, err := mockUsecase.GetByUserID(1)
		assert.NoError(t, err)
		assert.Equal(t, []entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("success not empty", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetByUserID", 1).Return([]entities.Complaint{{ID: "1"}}, nil)

		result, err := mockUsecase.GetByUserID(1)
		assert.NoError(t, err)
		assert.Equal(t, []entities.Complaint{{ID: "1"}}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetByUserID", 1).Return([]entities.Complaint(nil), constants.ErrInternalServerError)

		result, err := mockUsecase.GetByUserID(1)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, []entities.Complaint(nil), result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)
		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("Create", &complaint).Return(nil)

		result, err := mockUsecase.Create(&complaint)
		assert.NoError(t, err)
		assert.Equal(t, complaint, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed all fields must be filled", func(t *testing.T) {
		complaint := entities.Complaint{
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        time.Time{},
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Create(&complaint)
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed regency not found", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Create", &complaint).Return(errors.New("REFERENCES `regencies` (`id`))"))

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Create(&complaint)
		assert.Error(t, constants.ErrRegencyNotFound, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed category not found", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Create", &complaint).Return(errors.New("REFERENCES `categories` (`id`))"))

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Create(&complaint)
		assert.Error(t, constants.ErrCategoryNotFound, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Create", &complaint).Return(constants.ErrInternalServerError)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Create(&complaint)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success admin delete", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("AdminDelete", "1").Return(nil)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.Delete("1", 1, "admin")
		assert.NoError(t, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("success user delete", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Delete", "1", 1).Return(nil)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.Delete("1", 1, "user")
		assert.NoError(t, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error admin delete", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("AdminDelete", "1").Return(constants.ErrInternalServerError)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.Delete("1", 1, "admin")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error user delete", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Delete", "1", 1).Return(constants.ErrInternalServerError)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.Delete("1", 1, "user")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			ID:          "1",
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Update", complaint).Return(complaint, nil)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Update(complaint)
		assert.NoError(t, err)
		assert.Equal(t, complaint, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed all fields must be filled", func(t *testing.T) {
		complaint := entities.Complaint{
			ID:          "1",
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        time.Time{},
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Update(complaint)
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed regency not found", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			ID:          "1",
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Update", complaint).Return(entities.Complaint{}, errors.New("REFERENCES `regencies` (`id`))"))

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Update(complaint)
		assert.Error(t, constants.ErrRegencyNotFound, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed category not found", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			ID:          "1",
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Update", complaint).Return(entities.Complaint{}, errors.New("REFERENCES `categories` (`id`))"))

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Update(complaint)
		assert.Error(t, constants.ErrCategoryNotFound, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		date, _ := time.Parse("2006-01-02", "2021-01-01")
		complaint := entities.Complaint{
			ID:          "1",
			UserID:      1,
			CategoryID:  1,
			RegencyID:   "1901",
			Description: "description",
			Address:     "address",
			Type:        "public",
			Date:        date,
		}

		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("Update", complaint).Return(entities.Complaint{}, constants.ErrInternalServerError)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		result, err := mockUsecase.Update(complaint)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.Complaint{}, result)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("UpdateStatus", "1", "Selesai").Return(nil)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.UpdateStatus("1", "Selesai")
		assert.NoError(t, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockComplaintRepo.On("UpdateStatus", "1", "Selesai").Return(constants.ErrInternalServerError)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.UpdateStatus("1", "Selesai")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed invalid status", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.UpdateStatus("1", "Invalid")
		assert.Error(t, constants.ErrInvalidStatus, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})

	t.Run("failed empty id", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		mockUsecase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		err := mockUsecase.UpdateStatus("", "Selesai")
		assert.Error(t, constants.ErrIDMustBeFilled, err)

		mockComplaintRepo.AssertExpectations(t)
		mockComplaintFileRepo.AssertExpectations(t)
	})
}

// Mock utility function
func GetRowsFromExcel(file *multipart.FileHeader) ([][]string, error) {
	// This is just a placeholder. You should mock this function as needed.
	return nil, nil
}

func TestImport(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		fileHeader := &multipart.FileHeader{}

		rows := [][]string{
			{"UserID", "CategoryID", "RegencyID", "Address", "Description", "Status", "Type", "Date", "Files"},
			{"1", "2", "Regency1", "Address1", "Description1", "Verifikasi", "Type1", "05-05-2024", "file1.jpg,file2.jpg"},
			{"1", "2", "Regency1", "Address1", "Description1", "On Progress", "Type1", "05-05-2024", "file1.jpg"},
			{"1", "2", "Regency1", "Address1", "Description1", "Selesai", "Type1", "05-05-2024", "file1.jpg,file2.jpg"},
			{"1", "2", "Regency1", "Address1", "Description1", "Ditolak", "Type1", "05-05-2024", "file1.jpg"},
			{"1", "2", "Regency1", "Address1", "Description1", "Pending", "Type1", "05-05-2024", "file1.jpg"},
		}

		complaintUseCase.getRowsFromExcel = func(file *multipart.FileHeader) ([][]string, error) {
			return rows, nil
		}

		mockComplaintRepo.On("Import", mock.Anything).Return(nil)

		err := complaintUseCase.Import(fileHeader)
		assert.NoError(t, err)

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("failed getting rows from excel", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		fileHeader := &multipart.FileHeader{}
		complaintUseCase.getRowsFromExcel = func(file *multipart.FileHeader) ([][]string, error) {
			return nil, errors.New("failed to read excel")
		}

		err := complaintUseCase.Import(fileHeader)
		assert.Error(t, err)
		assert.Equal(t, "failed to read excel", err.Error())
	})

	t.Run("failed importing complaints", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		fileHeader := &multipart.FileHeader{}

		rows := [][]string{
			{"UserID", "CategoryID", "RegencyID", "Address", "Description", "Status", "Type", "Date", "Files"},
			{"1", "2", "Regency1", "Address1", "Description1", "Verifikasi", "Type1", "05-05-2024", "file1.jpg"},
		}

		complaintUseCase.getRowsFromExcel = func(file *multipart.FileHeader) ([][]string, error) {
			return rows, nil
		}

		mockComplaintRepo.On("Import", mock.Anything).Return(errors.New("import failed"))

		err := complaintUseCase.Import(fileHeader)
		assert.Error(t, err)
		assert.Equal(t, "import failed", err.Error())

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("failed not enough columns", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		fileHeader := &multipart.FileHeader{}

		rows := [][]string{
			{"UserID", "CategoryID", "RegencyID", "Address", "Description"},
			{"1", "2", "Regency1", "Address1"},
		}

		complaintUseCase.getRowsFromExcel = func(file *multipart.FileHeader) ([][]string, error) {
			return rows, nil
		}

		err := complaintUseCase.Import(fileHeader)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrColumnsDoesntMatch, err)
	})

	t.Run("failed invalid user id", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		fileHeader := &multipart.FileHeader{}

		rows := [][]string{
			{"UserID", "CategoryID", "RegencyID", "Address", "Description", "Status", "Type", "Date", "Files"},
			{"a", "2", "Regency1", "Address1", "Description1", "Verifikasi", "Type1", "05-05-2024", "file1.jpg"},
		}

		complaintUseCase.getRowsFromExcel = func(file *multipart.FileHeader) ([][]string, error) {
			return rows, nil
		}

		err := complaintUseCase.Import(fileHeader)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInvalidIDFormat, err)

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("failed invalid category id", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		fileHeader := &multipart.FileHeader{}

		rows := [][]string{
			{"UserID", "CategoryID", "RegencyID", "Address", "Description", "Status", "Type", "Date", "Files"},
			{"1", "a", "Regency1", "Address1", "Description1", "Verifikasi", "Type1", "05-05-2024", "file1.jpg"},
		}

		complaintUseCase.getRowsFromExcel = func(file *multipart.FileHeader) ([][]string, error) {
			return rows, nil
		}

		err := complaintUseCase.Import(fileHeader)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInvalidCategoryIDFormat, err)

		mockComplaintRepo.AssertExpectations(t)
	})
}

func TestIncreaseTotalLikes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("IncreaseTotalLikes", "1").Return(nil)

		err := complaintUseCase.IncreaseTotalLikes("1")
		assert.NoError(t, err)

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("IncreaseTotalLikes", "1").Return(constants.ErrInternalServerError)

		err := complaintUseCase.IncreaseTotalLikes("1")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockComplaintRepo.AssertExpectations(t)
	})
}

func TestDecreaseTotalLikes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("DecreaseTotalLikes", "1").Return(nil)

		err := complaintUseCase.DecreaseTotalLikes("1")
		assert.NoError(t, err)

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("DecreaseTotalLikes", "1").Return(constants.ErrInternalServerError)

		err := complaintUseCase.DecreaseTotalLikes("1")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockComplaintRepo.AssertExpectations(t)
	})
}

func TestGetComplaintIDsByUserID(t *testing.T) {
	t.Run("success empty", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetComplaintIDsByUserID", 1).Return([]string{}, nil)

		result, err := complaintUseCase.GetComplaintIDsByUserID(1)
		assert.NoError(t, err)
		assert.Equal(t, []string{}, result)

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("success not empty", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetComplaintIDsByUserID", 1).Return([]string{"1", "2"}, nil)

		result, err := complaintUseCase.GetComplaintIDsByUserID(1)
		assert.NoError(t, err)
		assert.Equal(t, []string{"1", "2"}, result)

		mockComplaintRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockComplaintRepo := new(MockComplaintRepo)
		mockComplaintFileRepo := new(MockComplaintFileRepo)

		complaintUseCase := NewComplaintUseCase(mockComplaintRepo, mockComplaintFileRepo)

		mockComplaintRepo.On("GetComplaintIDsByUserID", 1).Return([]string(nil), constants.ErrInternalServerError)

		result, err := complaintUseCase.GetComplaintIDsByUserID(1)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, []string(nil), result)

		mockComplaintRepo.AssertExpectations(t)
	})
}
