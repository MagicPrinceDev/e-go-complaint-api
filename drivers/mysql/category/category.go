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
			return entities.Category{}, constants.ErrCategoryNotFound
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
			return nil, constants.ErrCategoryNotFound
		}
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepo) DeleteCategory(id int) error {
	complaints := r.DB.Where("category_id = ?", id).Find(&entities.Complaint{})
	if complaints.RowsAffected > 0 {
		return constants.ErrCategoryHasBeenUsed
	}

	news := r.DB.Where("category_id = ?", id).Find(&entities.News{})
	if news.RowsAffected > 0 {
		return constants.ErrCategoryHasBeenUsed
	}

	if err := r.DB.Delete(&entities.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
