package seeder

import (
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	if err := db.First(&entities.Admin{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		hash, _ := utils.HashPassword("password")
		admins := []entities.Admin{
			{
				Name:            "Super Admin",
				Password:        hash,
				Email:           "super_admin@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    true,
			},
			{
				Name:            "Admin Pandeglang",
				Password:        hash,
				Email:           "admin_pandeglang@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
			},
			{
				Name:            "Admin Lebak",
				Password:        hash,
				Email:           "admin_lebak@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
			},
			{
				Name:            "Admin Serang",
				Password:        hash,
				Email:           "admin_serang@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
			},
		}

		if err := db.CreateInBatches(&admins, len(admins)).Error; err != nil {
			panic(err)
		}
	}
}
