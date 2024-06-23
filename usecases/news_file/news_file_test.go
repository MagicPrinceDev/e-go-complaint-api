package news_file

import (
	"e-complaint-api/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"testing"
)

type NewsFileMock struct {
	mock.Mock
}

func (m *NewsFileMock) Create(newsFiles []*entities.NewsFile) error {
	args := m.Called(newsFiles)
	return args.Error(0)
}
func (m *NewsFileMock) DeleteByNewsID(newsID int) error {
	args := m.Called(newsID)
	return args.Error(0)
}

type NewsFileGCSAPIMock struct {
	mock.Mock
}

func (m *NewsFileGCSAPIMock) Delete(filePaths []string) error {
	args := m.Called(filePaths)
	return args.Error(0)
}

func (m *NewsFileGCSAPIMock) Upload(files []*multipart.FileHeader) ([]string, error) {
	args := m.Called(files)
	return args.Get(0).([]string), args.Error(1)
}

func TestNewsFileUseCase_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		newsFileMock := new(NewsFileMock)
		newsFileGCSAPIMock := new(NewsFileGCSAPIMock)
		newsFileUseCase := NewNewsFileUseCase(newsFileMock, newsFileGCSAPIMock)

		files := []*multipart.FileHeader{
			&multipart.FileHeader{},
			&multipart.FileHeader{},
		}
		newsID := 1

		newsFileGCSAPIMock.On("Upload", files).Return([]string{"path1", "path2"}, nil)
		newsFileMock.On("Create", mock.Anything).Return(nil)

		_, err := newsFileUseCase.Create(files, newsID)
		assert.Nil(t, err)
	})

	t.Run("upload error", func(t *testing.T) {
		newsFileMock := new(NewsFileMock)
		newsFileGCSAPIMock := new(NewsFileGCSAPIMock)
		newsFileUseCase := NewNewsFileUseCase(newsFileMock, newsFileGCSAPIMock)

		files := []*multipart.FileHeader{
			&multipart.FileHeader{},
			&multipart.FileHeader{},
		}
		newsID := 1

		newsFileGCSAPIMock.On("Upload", files).Return([]string{}, errors.New("upload error"))

		_, err := newsFileUseCase.Create(files, newsID)
		assert.NotNil(t, err)
	})

	t.Run("create error", func(t *testing.T) {
		newsFileMock := new(NewsFileMock)
		newsFileGCSAPIMock := new(NewsFileGCSAPIMock)
		newsFileUseCase := NewNewsFileUseCase(newsFileMock, newsFileGCSAPIMock)

		files := []*multipart.FileHeader{
			&multipart.FileHeader{},
			&multipart.FileHeader{},
		}
		newsID := 1

		newsFileGCSAPIMock.On("Upload", files).Return([]string{"path1", "path2"}, nil)
		newsFileMock.On("Create", mock.Anything).Return(errors.New("create error"))

		_, err := newsFileUseCase.Create(files, newsID)
		assert.NotNil(t, err)
	})
}

func TestNewsFileUseCase_DeleteByNewsID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		newsFileMock := new(NewsFileMock)
		newsFileGCSAPIMock := new(NewsFileGCSAPIMock)
		newsFileUseCase := NewNewsFileUseCase(newsFileMock, newsFileGCSAPIMock)

		newsID := 1

		newsFileMock.On("DeleteByNewsID", newsID).Return(nil)
		newsFileGCSAPIMock.On("Delete", mock.Anything).Return(nil)

		err := newsFileUseCase.DeleteByNewsID(newsID)
		assert.Nil(t, err)
	})

	t.Run("delete error", func(t *testing.T) {
		newsFileMock := new(NewsFileMock)
		newsFileGCSAPIMock := new(NewsFileGCSAPIMock)
		newsFileUseCase := NewNewsFileUseCase(newsFileMock, newsFileGCSAPIMock)

		newsID := 1

		newsFileMock.On("DeleteByNewsID", newsID).Return(errors.New("delete error"))

		err := newsFileUseCase.DeleteByNewsID(newsID)
		assert.NotNil(t, err)
	})
}
