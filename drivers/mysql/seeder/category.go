package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) {
	if err := db.First(&entities.Category{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		categories := []entities.Category{
			{
				Name:        "Kesehatan",
				Description: "Kategori pengaduan yang berkaitan dengan kesehatan seperti fasilitas kesehatan, obat-obatan, dan lain-lain",
			},
			{
				Name:        "Pendidikan",
				Description: "Kategori pengaduan yang berkaitan dengan pendidikan seperti fasilitas pendidikan, kurikulum, dan lain-lain",
			},
			{
				Name:        "Kependudukan",
				Description: "Kategori pengaduan yang berkaitan dengan kependudukan seperti administrasi kependudukan, pelayanan kependudukan, dan lain-lain",
			},
			{
				Name:        "Keamanan",
				Description: "Kategori pengaduan yang berkaitan dengan keamanan seperti kejahatan, kecelakaan, dan lain-lain",
			},
			{
				Name:        "Infrastuktur",
				Description: "Kategori pengaduan yang berkaitan dengan infrastuktur seperti jalan, jembatan, dan lain-lain",
			},
			{
				Name:        "Lingkungan",
				Description: "Kategori pengaduan yang berkaitan dengan lingkungan seperti sampah, polusi, dan lain-lain",
			},
			{
				Name:        "Transportasi",
				Description: "Kategori pengaduan yang berkaitan dengan transportasi seperti angkutan umum, jadwal transportasi, dan lain-lain",
			},
		}

		if err := db.CreateInBatches(&categories, len(categories)).Error; err != nil {
			panic(err)
		}
	}
}
