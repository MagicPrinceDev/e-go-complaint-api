package seeder

import (
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	if err := db.First(&entities.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		hash, _ := utils.HashPassword("password")
		users := []entities.User{
			{
				Name:            "Putra Ramadhan",
				Password:        hash,
				Email:           "putraramadhan@gmail.com",
				TelephoneNumber: "081234567890",
				EmailVerified:   true,
				ProfilePhoto:    "profile-photos/example_1.jpg",
			},
			{
				Name:            "Andika Saputra",
				Password:        hash,
				Email:           "andikaputra@gmail.com",
				TelephoneNumber: "081234567890",
				EmailVerified:   true,
				ProfilePhoto:    "profile-photos/example_2.jpg",
			},
			{
				Name:            "Muhammad Iqbal",
				Password:        hash,
				Email:           "muhammadiqbal@gmail.com",
				TelephoneNumber: "081234567890",
				EmailVerified:   true,
				ProfilePhoto:    "profile-photos/example_3.jpg",
			},
		}

		if err := db.CreateInBatches(&users, len(users)).Error; err != nil {
			panic(err)
		}
	}
}
