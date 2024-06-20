package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedNewsComment(db *gorm.DB) {
	var newsComment []entities.NewsComment

	userID1 := 1
	userID2 := 2
	userID3 := 3
	adminID2 := 2
	adminID3 := 3

	if err := db.First(&entities.NewsComment{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		newsComment = []entities.NewsComment{
			{
				UserID:  &userID1,
				NewsID:  1,
				Comment: "Terimakasih atas informasinya",
			},
			{
				AdminID: &adminID2,
				NewsID:  1,
				Comment: "Sama-sama, semoga bermanfaat",
			},
			{
				UserID:  &userID2,
				NewsID:  1,
				Comment: "Terimakasih telah memperbaiki fasilitas kesehatan kami",
			},
			{
				UserID:  &userID3,
				NewsID:  2,
				Comment: "Terimakasih telah memperbaiki fasilitas SMA kami",
			},
			{
				AdminID: &adminID3,
				NewsID:  2,
				Comment: "Sama-sama, semoga bermanfaat",
			},
			{
				UserID:  &userID1,
				NewsID:  3,
				Comment: "Terimakasih telah meningkatkan keamanan di lingkungan kami",
			},
			{
				AdminID: &adminID3,
				NewsID:  3,
				Comment: "Sama-sama, semoga lingkungan anda semakin aman",
			},
		}
	}

	if err := db.CreateInBatches(&newsComment, len(newsComment)).Error; err != nil {
		panic(err)
	}
}
