package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedComplaintFile(db *gorm.DB) {
	if err := db.First(&entities.ComplaintFile{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		complaintFiles := []entities.ComplaintFile{
			{
				ComplaintID: "C-81j9aK9280",
				Path:        "complaint-files/example_kesehatan_1.jpg",
			},
			{
				ComplaintID: "C-81j9aK9280",
				Path:        "complaint-files/example_kesehatan_2.jpg",
			},
			{
				ComplaintID: "C-8kshis9280",
				Path:        "complaint-files/example_pendidikan_1.jpg",
			},
			{
				ComplaintID: "C-8kshis9280",
				Path:        "complaint-files/example_pendidikan_2.jpg",
			},
			{
				ComplaintID: "C-81jas92581",
				Path:        "complaint-files/example_kependudukan_1.jpg",
			},
			{
				ComplaintID: "C-81jas92581",
				Path:        "complaint-files/example_kependudukan_2.jpg",
			},
			{
				ComplaintID: "C-271j9ak280",
				Path:        "complaint-files/example_keamanan_1.jpg",
			},
			{
				ComplaintID: "C-271j9ak280",
				Path:        "complaint-files/example_keamanan_2.jpg",
			},
			{
				ComplaintID: "C-123j9ak280",
				Path:        "complaint-files/example_lingkungan_1.jpg",
			},
			{
				ComplaintID: "C-123j9ak280",
				Path:        "complaint-files/example_lingkungan_2.jpg",
			},
		}

		if err := db.CreateInBatches(&complaintFiles, len(complaintFiles)).Error; err != nil {
			panic(err)
		}
	}
}
