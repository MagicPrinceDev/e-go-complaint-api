package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedNewsLike(db *gorm.DB) {
	if err := db.First(&entities.NewsLike{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		newsLikes := []entities.NewsLike{
			{
				NewsID: 1,
				UserID: 1,
			},
			{
				NewsID: 1,
				UserID: 2,
			},
			{
				NewsID: 1,
				UserID: 3,
			},
			{
				NewsID: 2,
				UserID: 1,
			},
			{
				NewsID: 2,
				UserID: 2,
			},
			{
				NewsID: 3,
				UserID: 1,
			},
			{
				NewsID: 3,
				UserID: 2,
			},
			{
				NewsID: 3,
				UserID: 3,
			},
		}

		if err := db.CreateInBatches(newsLikes, len(newsLikes)).Error; err != nil {
			panic(err)
		}
	}
}
