package entities

import (
	"time"
)

type ComplaintActivity struct {
	ID           int    `gorm:"primaryKey"`
	ComplaintID  string `gorm:"type:varchar;size:15;not null"`
	DiscussionID *int
	LikeID       *int
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `gorm:"autoUpdateTime"`
	Complaint    Complaint     `gorm:"foreignKey:ComplaintID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Discussion   Discussion    `gorm:"foreignKey:DiscussionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Like         ComplaintLike `gorm:"foreignKey:LikeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ComplaintActivityRepositoryInterface interface {
	GetByComplaintIDs(complaintIDs []string, activityType string) ([]ComplaintActivity, error)
	Create(complaintActivity *ComplaintActivity) error
	Delete(complaintActivity ComplaintActivity) error
	Update(complaintActivity ComplaintActivity) error
}

type ComplaintActivityUseCaseInterface interface {
	GetByComplaintIDs(complaintIDs []string, activityType string) ([]ComplaintActivity, error)
	Create(complaintActivity *ComplaintActivity) (ComplaintActivity, error)
	Delete(complaintActivity ComplaintActivity) error
	Update(complaintActivity ComplaintActivity) error
}
