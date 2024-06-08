package complaint_likes

import (
	"e-complaint-api/entities"
	"gorm.io/gorm"
)

type ComplaintLikeRepo struct {
	DB *gorm.DB
}

func NewComplaintLikeRepo(db *gorm.DB) *ComplaintLikeRepo {
	return &ComplaintLikeRepo{DB: db}
}

func (r *ComplaintLikeRepo) Likes(complaintLike []*entities.ComplaintLike) error {
	if err := r.DB.Create(complaintLike).Error; err != nil {
		return err
	}
	return nil
}

func (r *ComplaintLikeRepo) Unlike(complaintLike []*entities.ComplaintLike) error {
	if err := r.DB.Delete(complaintLike).Error; err != nil {
		return err
	}
	return nil
}
