package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedComplaintProcess(db *gorm.DB) {
	if err := db.First(&entities.ComplaintProcess{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		complaintProcesses := []entities.ComplaintProcess{
			{
				ComplaintID: "C-8ksh&s9280",
				AdminID:     2,
				Status:      "verifikasi",
				Message:     "Aduan anda telah diverifikasi oleh kami",
			},
			{
				ComplaintID: "C-8ksh&s9280",
				AdminID:     2,
				Status:      "on progress",
				Message:     "Aduan anda sedang dalam proses penanganan",
			},
			{
				ComplaintID: "C-8ksh&s9280",
				AdminID:     2,
				Status:      "selesai",
				Message:     "Aduan anda telah selesai ditangani",
			},
			{
				ComplaintID: "C-81jas92581",
				AdminID:     3,
				Status:      "verifikasi",
				Message:     "Aduan anda telah diverifikasi oleh kami",
			},
			{
				ComplaintID: "C-271j9ak280",
				AdminID:     3,
				Status:      "verifikasi",
				Message:     "Aduan anda telah diverifikasi oleh kami",
			},
			{
				ComplaintID: "C-271j9ak280",
				AdminID:     3,
				Status:      "on progress",
				Message:     "Sedang dalam proses penanganan",
			},
			{
				ComplaintID: "C-123j9ak280",
				AdminID:     4,
				Status:      "ditolak",
				Message:     "Aduan anda ditolak karena tidak sesuai dengan ketentuan yang berlaku",
			},
		}

		for _, complaintProcess := range complaintProcesses {
			db.Create(&complaintProcess)
		}
	}
}
