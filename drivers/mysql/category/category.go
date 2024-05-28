package category

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (r *CategoryRepo) GetAll() ([]entities.Category, error) {
	var categories []entities.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepo) GetByID(id int) (entities.Category, error) {
	var category entities.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Category{}, constants.ErrNotFound
		}
		return entities.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepo) CreateCategory(category *entities.Category) (*entities.Category, error) {
	if err := r.DB.Create(&category).Error; err != nil {
		return &entities.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepo) UpdateCategory(id int, category *entities.Category) (*entities.Category, error) {
	if err := r.DB.Model(&entities.Category{}).Where("id = ?", id).Updates(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.ErrNotFound
		}
		return nil, err
	}

	return category, nil
}
