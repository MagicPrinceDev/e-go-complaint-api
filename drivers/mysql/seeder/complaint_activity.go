package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedComplaintActivity(db *gorm.DB) {
	if err := db.First(&entities.ComplaintActivity{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		DiscussionID1 := 1
		DiscussionID2 := 2
		DiscussionID3 := 3
		DiscussionID4 := 4
		DiscussionID5 := 5
		DiscussionID6 := 6
		LikeID1 := 1
		LikeID2 := 2
		LikeID3 := 3
		LikeID4 := 4
		LikeID5 := 5
		LikeID6 := 6
		LikeID7 := 7
		LikeID8 := 8

		complaintActivities := []entities.ComplaintActivity{
			{
				ComplaintID:  "C-81j9aK9280",
				DiscussionID: &DiscussionID1,
			},
			{
				ComplaintID:  "C-81j9aK9280",
				DiscussionID: &DiscussionID2,
			},
			{
				ComplaintID:  "C-271j9ak280",
				DiscussionID: &DiscussionID3,
			},
			{
				ComplaintID:  "C-271j9ak280",
				DiscussionID: &DiscussionID4,
			},
			{
				ComplaintID:  "C-81jas92581",
				DiscussionID: &DiscussionID5,
			},
			{
				ComplaintID:  "C-81jas92581",
				DiscussionID: &DiscussionID6,
			},
			{
				ComplaintID: "C-81j9aK9280",
				LikeID:      &LikeID1,
			},
			{
				ComplaintID: "C-81j9aK9280",
				LikeID:      &LikeID2,
			},
			{
				ComplaintID: "C-81j9aK9280",
				LikeID:      &LikeID3,
			},
			{
				ComplaintID: "C-8kshis9280",
				LikeID:      &LikeID4,
			},
			{
				ComplaintID: "C-8kshis9280",
				LikeID:      &LikeID5,
			},
			{
				ComplaintID: "C-81jas92581",
				LikeID:      &LikeID6,
			},
			{
				ComplaintID: "C-81jas92581",
				LikeID:      &LikeID7,
			},
			{
				ComplaintID: "C-271j9ak280",
				LikeID:      &LikeID8,
			},
		}

		if err := db.CreateInBatches(&complaintActivities, len(complaintActivities)).Error; err != nil {
			panic(err)
		}
	}
}
