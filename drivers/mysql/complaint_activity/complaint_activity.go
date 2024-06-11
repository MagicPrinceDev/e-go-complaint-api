package complaint_activity

import (
	"e-complaint-api/entities"
	"time"

	"gorm.io/gorm"
)

type ComplaintActivityRepo struct {
	DB *gorm.DB
}

func NewComplaintActivityRepo(db *gorm.DB) *ComplaintActivityRepo {
	return &ComplaintActivityRepo{DB: db}
}

func (r *ComplaintActivityRepo) GetByComplaintIDs(complaintIDs []string, activityType string) ([]entities.ComplaintActivity, error) {
	var complaintActivities []entities.ComplaintActivity

	if activityType == "" {
		if err := r.DB.Preload("Like").Preload("Discussion").Preload("Like.User").Preload("Discussion.Admin").Preload("Discussion.User").Where("complaint_id IN ?", complaintIDs).Find(&complaintActivities).Error; err != nil {
			return nil, err
		}
	} else if activityType == "like" {
		if err := r.DB.Preload("Like").Preload("Like.User").Where("complaint_id IN ?", complaintIDs).Where("like_id IS NOT NULL").Find(&complaintActivities).Error; err != nil {
			return nil, err
		}
	} else if activityType == "discussion" {
		if err := r.DB.Preload("Discussion").Preload("Discussion.Admin").Preload("Discussion.User").Where("complaint_id IN ?", complaintIDs).Where("discussion_id IS NOT NULL").Find(&complaintActivities).Error; err != nil {
			return nil, err
		}
	}

	return complaintActivities, nil
}

func (r *ComplaintActivityRepo) Create(complaintActivity *entities.ComplaintActivity) error {
	if err := r.DB.Create(complaintActivity).Error; err != nil {
		return err
	}

	return nil
}

func (r *ComplaintActivityRepo) Delete(complaintActivity entities.ComplaintActivity) error {
	if complaintActivity.LikeID == nil {
		if err := r.DB.Where("complaint_id = ? AND discussion_id = ?", complaintActivity.ComplaintID, *complaintActivity.DiscussionID).Delete(&complaintActivity).Error; err != nil {
			return err
		}
	} else if complaintActivity.DiscussionID == nil {
		if err := r.DB.Where("complaint_id = ? AND like_id = ?", complaintActivity.ComplaintID, *complaintActivity.LikeID).Delete(&complaintActivity).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *ComplaintActivityRepo) Update(complaintActivity entities.ComplaintActivity) error {
	complaintActivity.UpdatedAt = time.Now()
	if err := r.DB.Where("complaint_id = ? AND discussion_id = ?", complaintActivity.ComplaintID, complaintActivity.DiscussionID).Updates(&complaintActivity).Error; err != nil {
		return err
	}

	return nil
}
