package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedNewsFile(db *gorm.DB) {
	if err := db.First(&entities.NewsFile{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		newsFiles := []entities.NewsFile{
			{
				NewsID: 1,
				Path:   "news_files/example1.jpg",
			},
			{
				NewsID: 1,
				Path:   "news_files/example2.jpg",
			},
			{
				NewsID: 2,
				Path:   "news_files/example3.jpg",
			},
			{
				NewsID: 3,
				Path:   "news_files/example1.jpg",
			},
			{
				NewsID: 3,
				Path:   "news_files/example3.jpg",
			},
		}

		if err := db.CreateInBatches(newsFiles, len(newsFiles)).Error; err != nil {
			panic(err)
		}
	}
}
