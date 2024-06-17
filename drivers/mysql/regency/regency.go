package regency

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type RegencyRepo struct {
	DB *gorm.DB
}

func NewRegencyRepo(db *gorm.DB) *RegencyRepo {
	return &RegencyRepo{DB: db}
}

func (r *RegencyRepo) GetAll() ([]entities.Regency, error) {
	var regencies []entities.Regency
	if err := r.DB.Find(&regencies).Error; err != nil {
		return nil, err
	}

	return regencies, nil
}
