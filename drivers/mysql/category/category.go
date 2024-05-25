package category

import (
	"e-complaint-api/entities"

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
