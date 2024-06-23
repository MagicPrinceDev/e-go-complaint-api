package complaint_file

import (
	"e-complaint-api/entities"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockComplaintFileRepository struct {
	mock.Mock
}

func (m *MockComplaintFileRepository) Create(complaintFiles []*entities.ComplaintFile) error {
	args := m.Called(complaintFiles)
	return args.Error(0)
}

func (m *MockComplaintFileRepository) DeleteByComplaintID(complaintID string) error {
	args := m.Called(complaintID)
	return args.Error(0)
}

type MockComplaintFileGCSAPI struct {
	mock.Mock
}

func (m *MockComplaintFileGCSAPI) Upload(files []*multipart.FileHeader) ([]string, error) {
	args := m.Called(files)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockComplaintFileGCSAPI) Delete(paths []string) error {
	args := m.Called(paths)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := new(MockComplaintFileRepository)
		gcs_api := new(MockComplaintFileGCSAPI)
		usecase := NewComplaintFileUseCase(repo, gcs_api)

		files := []*multipart.FileHeader{}
		complaintID := "complaint_id"

		repo.On("Create", mock.Anything).Return(nil)
		gcs_api.On("Upload", files).Return([]string{"path"}, nil)

		result, err := usecase.Create(files, complaintID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("failed to upload", func(t *testing.T) {
		repo := new(MockComplaintFileRepository)
		gcs_api := new(MockComplaintFileGCSAPI)
		usecase := NewComplaintFileUseCase(repo, gcs_api)

		files := []*multipart.FileHeader{}
		complaintID := "complaint_id"

		gcs_api.On("Upload", files).Return([]string{}, errors.New("failed to upload"))

		result, err := usecase.Create(files, complaintID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("failed to create", func(t *testing.T) {
		repo := new(MockComplaintFileRepository)
		gcs_api := new(MockComplaintFileGCSAPI)
		usecase := NewComplaintFileUseCase(repo, gcs_api)

		files := []*multipart.FileHeader{}
		complaintID := "complaint_id"

		repo.On("Create", mock.Anything).Return(errors.New("failed to create"))
		gcs_api.On("Upload", files).Return([]string{"path"}, nil)

		result, err := usecase.Create(files, complaintID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestDeleteByComplaintID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := new(MockComplaintFileRepository)
		gcs_api := new(MockComplaintFileGCSAPI)
		usecase := NewComplaintFileUseCase(repo, gcs_api)

		complaintID := "complaint_id"

		repo.On("DeleteByComplaintID", complaintID).Return(nil)

		err := usecase.DeleteByComplaintID(complaintID)

		assert.NoError(t, err)
	})

	t.Run("failed to delete", func(t *testing.T) {
		repo := new(MockComplaintFileRepository)
		gcs_api := new(MockComplaintFileGCSAPI)
		usecase := NewComplaintFileUseCase(repo, gcs_api)

		complaintID := "complaint_id"

		repo.On("DeleteByComplaintID", complaintID).Return(errors.New("failed to delete"))

		err := usecase.DeleteByComplaintID(complaintID)

		assert.Error(t, err)
	})
}
