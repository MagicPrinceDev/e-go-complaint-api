package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedRegencyFromAPI(db *gorm.DB, api entities.RegencyIndonesiaAreaAPIInterface) {
	if err := db.First(&entities.Regency{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		regencies, err := api.GetRegenciesDataFromAPI()
		if err != nil {
			panic(err)
		}

		if err := db.CreateInBatches(regencies, len(regencies)).Error; err != nil {
			panic(err)
		}
	}
}
