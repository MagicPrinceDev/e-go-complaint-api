package entities

import "time"

type ComplaintLike struct {
	ID          int       `gorm:"primaryKey"`
	ComplaintID string    `gorm:"type:varchar(15);index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID      int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type ComplaintLikeRepositoryInterface interface {
	Unlike(complaintLike *ComplaintLike) error
	Likes(complaintLike *ComplaintLike) error
	ToggleLike(complaintLike *ComplaintLike) error
	FindByUserAndComplaint(userID int, complaintID string) (*ComplaintLike, error)
}

type ComplaintLikeUseCaseInterface interface {
	Likes(complaintLike *ComplaintLike) error
	Unlike(complaintLike *ComplaintLike) error
	ToggleLike(complaintLike *ComplaintLike) error
	FindByUserAndComplaint(userID int, complaintID string) (*ComplaintLike, error)
}
