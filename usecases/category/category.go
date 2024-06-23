package category

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
)

type CategoryUseCase struct {
	repository entities.CategoryRepositoryInterface
}

func NewCategoryUseCase(repo entities.CategoryRepositoryInterface) *CategoryUseCase {
	return &CategoryUseCase{repository: repo}
}

func (uc *CategoryUseCase) GetAll() ([]entities.Category, error) {
	categories, err := uc.repository.GetAll()
	if err != nil {
		return []entities.Category{}, constants.ErrInternalServerError
	}

	return categories, nil
}

func (uc *CategoryUseCase) GetByID(id int) (entities.Category, error) {
	category, err := uc.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrCategoryNotFound) {
			return entities.Category{}, constants.ErrCategoryNotFound
		}
		return entities.Category{}, constants.ErrInternalServerError
	}

	return category, nil
}

func (uc *CategoryUseCase) CreateCategory(category *entities.Category) (*entities.Category, error) {
	category, err := uc.repository.CreateCategory(category)
	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	if category == nil {
		return nil, nil
	}

	if category.Name == "" || category.Description == "" {
		return nil, constants.ErrAllFieldsMustBeFilled
	}

	return category, nil
}

func (uc *CategoryUseCase) UpdateCategory(id int, newCategory *entities.Category) (*entities.Category, error) {
	existingCategory, err := uc.repository.GetByID(id)

	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	if newCategory.Name == "" && newCategory.Description == "" {
		return nil, constants.ErrAllFieldsMustBeFilled
	}

	if newCategory.Name == "" {
		newCategory.Name = existingCategory.Name
	}

	if newCategory.Description == "" {
		newCategory.Description = existingCategory.Description
	}

	if existingCategory.Name == newCategory.Name && existingCategory.Description == newCategory.Description {
		return nil, constants.ErrNoChangesDetected
	}

	updatedCategory, err := uc.repository.UpdateCategory(id, newCategory)
	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return updatedCategory, nil
}

func (uc *CategoryUseCase) DeleteCategory(id int) error {
	_, err := uc.repository.GetByID(id)
	if err != nil {
		return constants.ErrCategoryNotFound
	}
	err = uc.repository.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
