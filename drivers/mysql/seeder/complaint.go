package seeder

import (
	"e-complaint-api/entities"
	"errors"
	"time"

	"gorm.io/gorm"
)

func SeedComplaint(db *gorm.DB) {
	if err := db.First(&entities.Complaint{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		complaints := []entities.Complaint{
			{
				ID:          "C-81j9aK9280",
				UserID:      1,
				CategoryID:  1,
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "3601",
				Address:     "Jl. lorem ipsum No. 1 RT 01 RW 01, Kelurahan Lorem Ipsum, Kecamatan Lorem Ipsum, Kota Lorem Ipsum, Provinsi Lorem Ipsum",
				Status:      "Pending",
				Type:        "public",
				Date:        time.Now(),
				TotalLikes:  3,
			},
			{
				ID:          "C-8ksh&s9280",
				UserID:      1,
				CategoryID:  2,
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "3603",
				Address:     "Jl. lorem ipsum No. 1 RT 01 RW 01, Kelurahan Lorem Ipsum, Kecamatan Lorem Ipsum, Kota Lorem Ipsum, Provinsi Lorem Ipsum",
				Status:      "Selesai",
				Type:        "private",
				Date:        time.Now(),
				TotalLikes:  2,
			},
			{
				ID:          "C-81jas92581",
				UserID:      2,
				CategoryID:  4,
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "3673",
				Address:     "Jl. lorem ipsum No. 1 RT 01 RW 01, Kelurahan Lorem Ipsum, Kecamatan Lorem Ipsum, Kota Lorem Ipsum, Provinsi Lorem Ipsum",
				Status:      "Verifikasi",
				Type:        "private",
				Date:        time.Now(),
				TotalLikes:  2,
			},
			{
				ID:          "C-271j9ak280",
				UserID:      3,
				CategoryID:  1,
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "3671",
				Address:     "Jl. lorem ipsum No. 1 RT 01 RW 01, Kelurahan Lorem Ipsum, Kecamatan Lorem Ipsum, Kota Lorem Ipsum, Provinsi Lorem Ipsum",
				Status:      "On Progress",
				Type:        "public",
				Date:        time.Now(),
				TotalLikes:  1,
			},
			{
				ID:          "C-123j9ak280",
				UserID:      3,
				CategoryID:  6,
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "3672",
				Address:     "Jl. lorem ipsum No. 1 RT 01 RW 01, Kelurahan Lorem Ipsum, Kecamatan Lorem Ipsum, Kota Lorem Ipsum, Provinsi Lorem Ipsum",
				Status:      "Ditolak",
				Type:        "public",
				Date:        time.Now(),
				TotalLikes:  0,
			},
		}

		if err := db.CreateInBatches(&complaints, len(complaints)).Error; err != nil {
			panic(err)
		}
	}
}
