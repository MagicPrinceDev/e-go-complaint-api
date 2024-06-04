package seeder

import (
	"e-complaint-api/entities"
	"errors"
	"gorm.io/gorm"
)

func SeedDiscussion(db *gorm.DB) {
	if err := db.First(&entities.Discussion{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		UserId := 1
		AdminID := 1
		discussions := []entities.Discussion{
			{
				UserID:      &UserId,
				ComplaintID: "C-271j9ak280",
				Comment:     "Min kenapa progressnya lama sekali ya",
			},
			{
				UserID:      &UserId,
				ComplaintID: "C-81jas92581",
				Comment:     "Min kenapa progressnya lama sekali ya",
			},
			{
				AdminID:     &AdminID,
				ComplaintID: "C-81jas92581",
				Comment:     "Mohon bersabar ya, sedang dalam proses",
			},
		}
		if err := db.CreateInBatches(&discussions, len(discussions)).Error; err != nil {
			panic(err)
		}
	}
}
