package seeder

import (
	"e-complaint-api/entities"
	"errors"
	"gorm.io/gorm"
)

func SeedDiscussion(db *gorm.DB) {
	if err := db.First(&entities.Discussion{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		discussions := []entities.Discussion{
			{
				ID:      1,
				User:    entities.User{},
				Comment: "Min kenapa progressnya lama sekali ya",
			},
			{
				ID:      2,
				Admin:   entities.Admin{},
				Comment: "Halo, untuk progress memang memakan waktu yang cukup lama ya. Dikarenakan proses yang kompleks",
			},
		}
		if err := db.CreateInBatches(&discussions, len(discussions)).Error; err != nil {
			panic(err)
		}

	}
}
