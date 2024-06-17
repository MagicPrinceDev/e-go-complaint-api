package seeder

import (
	"e-complaint-api/entities"
	"errors"
	"gorm.io/gorm"
)

func SeedNewsComment(db *gorm.DB) {
	var newsComment []entities.NewsComment

	userID1 := 1
	adminID2 := 2
	userID3 := 3

	if err := db.First(&entities.NewsComment{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		newsComment = []entities.NewsComment{
			{
				UserID:  &userID1,
				NewsID:  1,
				Comment: "Apa yang terjadi di sana?",
			},
			{
				AdminID: &adminID2,
				NewsID:  1,
				Comment: "Sedang terjadi tanah longsor di sana",
			},
			{
				UserID:  &userID3,
				NewsID:  2,
				Comment: "Apakah banyak korban jiwa?",
			},
			{
				AdminID: &adminID2,
				NewsID:  2,
				Comment: "Sekitar 10 Orang Sedang di evakuasi",
			},
		}
	}

	if err := db.CreateInBatches(&newsComment, len(newsComment)).Error; err != nil {
		panic(err)
	}
}
