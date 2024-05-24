package seeder

import (
	"e-complaint-api/entities"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func SeedVillageFromAPI(db *gorm.DB, api entities.VillageIndonesiaAreaAPIInterface) {
	if err := db.First(&entities.Village{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var districtIDs []string
		if err := db.Model(&entities.District{}).Pluck("id", &districtIDs).Error; err != nil {
			panic(err)
		}
		fmt.Println(districtIDs)

		villages, err := api.GetVillagesDataFromAPI(districtIDs)
		if err != nil {
			panic(err)
		}

		if err := db.CreateInBatches(villages, len(villages)).Error; err != nil {
			panic(err)
		}
	}
}
