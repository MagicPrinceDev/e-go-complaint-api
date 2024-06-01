package seeder

import (
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	if err := db.First(&entities.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		hash, _ := utils.HashPassword("user")
		users := []entities.User{
			{
				Name:            "User 1",
				Password:        hash,
				Email:           "user1@gmail.com",
				TelephoneNumber: "081234567890",
			},
			{
				Name:            "User 2",
				Password:        hash,
				Email:           "user2@gmail.com",
				TelephoneNumber: "081234567890",
			},
			{
				Name:            "User 3",
				Password:        hash,
				Email:           "user3@gmail.com",
				TelephoneNumber: "081234567890",
			},
		}

		if err := db.CreateInBatches(&users, len(users)).Error; err != nil {
			panic(err)
		}
	}
}
