package category

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
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
