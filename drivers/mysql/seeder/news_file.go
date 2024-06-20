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
				Path:   "news-files/example_pendidikan_3.jpg",
			},
			{
				NewsID: 1,
				Path:   "news-files/example_pendidikan_4.jpg",
			},
			{
				NewsID: 2,
				Path:   "news-files/example_kesehatan_1.jpg",
			},
			{
				NewsID: 2,
				Path:   "news-files/example_kesehatan_2.jpg",
			},
			{
				NewsID: 3,
				Path:   "news-files/example_keamanan_5.jpg",
			},
			{
				NewsID: 3,
				Path:   "news-files/example_keamanan_6.jpg",
			},
		}

		if err := db.CreateInBatches(newsFiles, len(newsFiles)).Error; err != nil {
			panic(err)
		}
	}
}
