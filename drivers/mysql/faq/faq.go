package faq

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type FaqRepo struct {
	DB *gorm.DB
}

func NewFaqRepo(db *gorm.DB) *FaqRepo {
	return &FaqRepo{DB: db}
}

func (r *FaqRepo) GetAll() ([]entities.Faq, error) {
	var faqs []entities.Faq
	if err := r.DB.Find(&faqs).Error; err != nil {
		return nil, err
	}
	return faqs, nil
}
