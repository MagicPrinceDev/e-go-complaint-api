package news_file

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type NewsFileRepo struct {
	DB *gorm.DB
}

func NewNewsFileRepo(db *gorm.DB) *NewsFileRepo {
	return &NewsFileRepo{DB: db}
}

func (r *NewsFileRepo) Create(newsFiles []*entities.NewsFile) error {
	if err := r.DB.CreateInBatches(newsFiles, len(newsFiles)).Error; err != nil {
		return err
	}

	return nil
}
