package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedComplaintLike(db *gorm.DB) {
	if err := db.First(&entities.ComplaintLike{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		complaintLikes := []entities.ComplaintLike{
			{
				UserID:      1,
				ComplaintID: "C-81j9aK9280",
			},
			{
				UserID:      2,
				ComplaintID: "C-81j9aK9280",
			},
			{
				UserID:      3,
				ComplaintID: "C-81j9aK9280",
			},
			{
				UserID:      2,
				ComplaintID: "C-8kshis9280",
			},
			{
				UserID:      3,
				ComplaintID: "C-8kshis9280",
			},
			{
				UserID:      1,
				ComplaintID: "C-81jas92581",
			},
			{
				UserID:      2,
				ComplaintID: "C-81jas92581",
			},
			{
				UserID:      1,
				ComplaintID: "C-271j9ak280",
			},
		}

		if err := db.CreateInBatches(&complaintLikes, len(complaintLikes)).Error; err != nil {
			panic(err)
		}
	}
}
