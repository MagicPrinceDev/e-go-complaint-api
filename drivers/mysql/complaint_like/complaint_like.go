package complaint_like

import (
	"e-complaint-api/entities"
	"errors"
	"gorm.io/gorm"
)

type ComplaintLikeRepository struct {
	DB *gorm.DB
}

func (clr *ComplaintLikeRepository) ToggleLike(complaintLike *entities.ComplaintLike) error {
	existingLike, err := clr.FindByUserAndComplaint(complaintLike.UserID, complaintLike.ComplaintID)
	if err != nil {
		return err
	}

	if existingLike != nil {
		return clr.Unlike(existingLike)
	}

	return clr.Likes(complaintLike)
}

func NewComplaintLikeRepository(DB *gorm.DB) *ComplaintLikeRepository {
	return &ComplaintLikeRepository{
		DB: DB,
	}
}

func (clr *ComplaintLikeRepository) FindByUserAndComplaint(userID int, complaintID string) (*entities.ComplaintLike, error) {
	var complaintLike entities.ComplaintLike
	result := clr.DB.Where("user_id = ? AND complaint_id = ?", userID, complaintID).First(&complaintLike)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &complaintLike, nil
}
func (clr *ComplaintLikeRepository) Likes(complaintLike *entities.ComplaintLike) error {
	result := clr.DB.Create(complaintLike)
	return result.Error
}

func (clr *ComplaintLikeRepository) Unlike(complaintLike *entities.ComplaintLike) error {
	result := clr.DB.Delete(complaintLike)
	return result.Error
}