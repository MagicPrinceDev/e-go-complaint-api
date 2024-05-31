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

func (r *NewsFileRepo) DeleteByNewsID(newsID int) error {
	var newsFiles []entities.NewsFile
	if err := r.DB.Where("news_id = ?", newsID).Find(&newsFiles).Error; err != nil {
		return err
	}

	for _, nf := range newsFiles {
		nf.DeletedAt = gorm.DeletedAt{Time: nf.UpdatedAt, Valid: true}
		if err := r.DB.Save(&nf).Error; err != nil {
			return err
		}
	}

	return nil
}
