package category

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCategory struct {
	mock.Mock
}

func (m *MockCategory) GetAll() ([]entities.Category, error) {
	args := m.Called()
	return args.Get(0).([]entities.Category), args.Error(1)
}

func (m *MockCategory) GetByID(id int) (entities.Category, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategory) CreateCategory(category *entities.Category) (*entities.Category, error) {
	args := m.Called(category)
	return args.Get(0).(*entities.Category), args.Error(1)
}

func (m *MockCategory) UpdateCategory(id int, newCategory *entities.Category) (*entities.Category, error) {
	args := m.Called(id, newCategory)
	return args.Get(0).(*entities.Category), args.Error(1)
}

func (m *MockCategory) DeleteCategory(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCategoryUseCase_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := []entities.Category{
			{
				ID:          1,
				Name:        "Category 1",
				Description: "Description 1",
			},
			{
				ID:          2,
				Name:        "Category 2",
				Description: "Description 2",
			},
		}

		mockCategory.On("GetAll").Return(category, nil)
		result, err := mockUseCase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, category, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		mockCategory.On("GetAll").Return([]entities.Category{}, constants.ErrInternalServerError)
		result, err := mockUseCase.GetAll()
		assert.Error(t, err)
		assert.Equal(t, []entities.Category{}, result)
		mockCategory.AssertExpectations(t)

	})

}

func TestCategoryUseCase_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			ID:          1,
			Name:        "Category 1",
			Description: "Description 1",
		}

		mockCategory.On("GetByID", 1).Return(category, nil)
		result, err := mockUseCase.GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, category, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("category not found", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		mockCategory.On("GetByID", 1).Return(entities.Category{}, constants.ErrCategoryNotFound)
		result, err := mockUseCase.GetByID(1)
		assert.Error(t, err)
		assert.Equal(t, entities.Category{}, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		mockCategory.On("GetByID", 1).Return(entities.Category{}, constants.ErrInternalServerError)
		result, err := mockUseCase.GetByID(1)
		assert.Error(t, err)
		assert.Equal(t, entities.Category{}, result)
		mockCategory.AssertExpectations(t)
	})

}

func TestCategoryUseCase_CreateCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := new(MockCategory)
		useCase := NewCategoryUseCase(mock)
		category := entities.Category{
			Name:        "Category 1",
			Description: "Description 1",
		}
		mock.On("CreateCategory", &category).Return(&category, nil)
		result, err := useCase.CreateCategory(&category)
		assert.NoError(t, err)
		assert.Equal(t, &category, result)
		mock.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mock := new(MockCategory)
		useCase := NewCategoryUseCase(mock)
		category := entities.Category{
			Name:        "Category 1",
			Description: "Description 1",
		}
		mock.On("CreateCategory", &category).Return((*entities.Category)(nil), constants.ErrInternalServerError)
		result, err := useCase.CreateCategory(&category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mock.AssertExpectations(t)
	})

	t.Run("all fields must be filled", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "",
			Description: "",
		}

		mockCategory.On("CreateCategory", &category).Return((*entities.Category)(nil), constants.ErrAllFieldsMustBeFilled)
		result, err := mockUseCase.CreateCategory(&category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("repository returns nil without error", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "Category 1",
			Description: "Description 1",
		}

		mockCategory.On("CreateCategory", &category).Return((*entities.Category)(nil), nil)
		result, err := mockUseCase.CreateCategory(&category)
		assert.NoError(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("name is empty", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "",
			Description: "Description 1",
		}

		mockCategory.On("CreateCategory", &category).Return(&category, nil)
		result, err := mockUseCase.CreateCategory(&category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("description is empty", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "Category 1",
			Description: "",
		}

		mockCategory.On("CreateCategory", &category).Return(&category, nil)
		result, err := mockUseCase.CreateCategory(&category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

}

func TestCategoryUseCase_UpdateCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := new(MockCategory)
		useCase := NewCategoryUseCase(mock)
		category := entities.Category{
			ID:          1,
			Name:        "Update Category 1",
			Description: "Description Update 1",
		}

		updatedCategory := entities.Category{
			ID:          1,
			Name:        "Updated Category 1",
			Description: "Updated Description 1",
		}

		mock.On("GetByID", 1).Return(category, nil)
		mock.On("UpdateCategory", 1, &updatedCategory).Return(&updatedCategory, nil)
		result, err := useCase.UpdateCategory(1, &updatedCategory)
		assert.NoError(t, err)
		assert.Equal(t, &updatedCategory, result)
		mock.AssertExpectations(t)
	})

	t.Run("fields must be filled", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "",
			Description: "",
		}

		mockCategory.On("GetByID", 1).Return(entities.Category{}, nil)
		result, err := mockUseCase.UpdateCategory(1, &category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("category not found", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "Category 1",
			Description: "Description 1",
		}

		mockCategory.On("GetByID", 1).Return(entities.Category{}, constants.ErrCategoryNotFound)
		result, err := mockUseCase.UpdateCategory(1, &category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("no changes detected", func(t *testing.T) {
		mock := new(MockCategory)
		useCase := NewCategoryUseCase(mock)

		category := entities.Category{
			ID:          1,
			Name:        "Update Category 1",
			Description: "Description Update 1",
		}

		updatedCategory := entities.Category{
			ID:          1,
			Name:        "Updated Category 1",
			Description: "Description Update 1",
		}

		mock.On("GetByID", 1).Return(category, nil)
		mock.On("UpdateCategory", 1, &updatedCategory).Return(&updatedCategory, errors.New("no changes detected"))
		result, err := useCase.UpdateCategory(1, &updatedCategory)
		assert.Error(t, err)
		assert.Nil(t, result)
		mock.AssertExpectations(t)
	})

	t.Run("name is empty", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "",
			Description: "Description 1",
		}

		mockCategory.On("GetByID", 1).Return(category, nil)
		result, err := mockUseCase.UpdateCategory(1, &category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("description is empty", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "Category 1",
			Description: "",
		}

		mockCategory.On("GetByID", 1).Return(category, nil)
		result, err := mockUseCase.UpdateCategory(1, &category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("name and description are empty", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "",
			Description: "",
		}

		mockCategory.On("GetByID", 1).Return(category, nil)
		result, err := mockUseCase.UpdateCategory(1, &category)
		assert.Error(t, err)
		assert.Nil(t, result)
		mockCategory.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		category := entities.Category{
			Name:        "Category 1",
			Description: "Description 1",
		}

		mockCategory.On("GetByID", 1).Return(entities.Category{}, constants.ErrInternalServerError)
		result, err := mockUseCase.UpdateCategory(1, &category)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestCategoryUseCase_DeleteCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := new(MockCategory)
		useCase := NewCategoryUseCase(mock)

		mock.On("GetByID", 1).Return(entities.Category{}, nil)
		mock.On("DeleteCategory", 1).Return(nil)
		err := useCase.DeleteCategory(1)
		assert.NoError(t, err)
		mock.AssertExpectations(t)
	})

	t.Run("category not found", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		mockCategory.On("GetByID", 1).Return(entities.Category{}, constants.ErrCategoryNotFound)
		err := mockUseCase.DeleteCategory(1)
		assert.Error(t, err)
		mockCategory.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		mockCategory.On("GetByID", 1).Return(entities.Category{}, constants.ErrInternalServerError)
		err := mockUseCase.DeleteCategory(1)
		assert.Error(t, err)
		mockCategory.AssertExpectations(t)
	})

	t.Run("delete category error", func(t *testing.T) {
		mockCategory := new(MockCategory)
		mockUseCase := NewCategoryUseCase(mockCategory)

		mockCategory.On("GetByID", 1).Return(entities.Category{}, nil)
		mockCategory.On("DeleteCategory", 1).Return(constants.ErrInternalServerError)
		err := mockUseCase.DeleteCategory(1)
		assert.Error(t, err)
		mockCategory.AssertExpectations(t)
	})
}
