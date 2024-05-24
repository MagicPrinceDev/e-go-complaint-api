package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedDistrictFromAPI(db *gorm.DB, api entities.DistrictIndonesiaAreaAPIInterface) {
	if err := db.First(&entities.District{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var regencyIDs []string
		if err := db.Model(&entities.Regency{}).Pluck("id", &regencyIDs).Error; err != nil {
			panic(err)
		}

		districts, err := api.GetDistrictsDataFromAPI(regencyIDs)
		if err != nil {
			panic(err)
		}

		if err := db.CreateInBatches(districts, len(districts)).Error; err != nil {
			panic(err)
		}
	}
}
