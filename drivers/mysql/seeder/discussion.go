package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedDiscussion(db *gorm.DB) {
	if err := db.First(&entities.Discussion{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		UserId1 := 1
		UserId2 := 2
		UserId3 := 3
		AdminID2 := 2
		discussions := []entities.Discussion{
			{
				UserID:      &UserId1,
				ComplaintID: "C-81j9aK9280",
				Comment:     "Min kenapa belum diverifikasi ya?",
			},
			{
				UserID:      &UserId3,
				ComplaintID: "C-81j9aK9280",
				Comment:     "Iya nih, kok lama banget verifikasinya?",
			},
			{
				UserID:      &UserId2,
				ComplaintID: "C-271j9ak280",
				Comment:     "Min kenapa progressnya lama sekali ya?",
			},
			{
				UserID:      &UserId1,
				ComplaintID: "C-271j9ak280",
				Comment:     "Iya nih, kok lama banget progressnya?",
			},
			{
				UserID:      &UserId1,
				ComplaintID: "C-81jas92581",
				Comment:     "Min kenapa belum diprogress juga ya?",
			},
			{
				AdminID:     &AdminID2,
				ComplaintID: "C-81jas92581",
				Comment:     "Mohon bersabar ya, progress akan segera dilakukan 1x24 jam kedepan",
			},
		}
		if err := db.CreateInBatches(&discussions, len(discussions)).Error; err != nil {
			panic(err)
		}
	}
}
