package entities

import "time"

type ComplaintLike struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	ComplaintID string    `gorm:"type:varchar;size:15;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Complaint   Complaint `gorm:"foreignKey:ComplaintID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ComplaintLikeRepositoryInterface interface {
	Unlike(complaintLike *ComplaintLike) error
	Likes(complaintLike *ComplaintLike) error
	FindByUserAndComplaint(userID int, complaintID string) (*ComplaintLike, error)
}

type ComplaintLikeUseCaseInterface interface {
	ToggleLike(complaintLike *ComplaintLike) (string, error)
}
