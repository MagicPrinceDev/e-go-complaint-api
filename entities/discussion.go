package entities

import (
	"gorm.io/gorm"
	"time"
)

type Discussion struct {
	ID          int            `gorm:"primaryKey"`
	Comment     string         `gorm:"not null"`
	UserID      int            `gorm:"not null"`
	User        User           `gorm:"foreignKey:UserID;references:ID"`
	ComplaintID int            `gorm:"not null"`
	Complaint   Complaint      `gorm:"foreignKey:ComplaintID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AdminID     int            `gorm:"not null"`
	Admin       Admin          `gorm:"foreignKey:AdminID;references:ID"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type DiscussionRepositoryInterface interface {
}

type DiscussionUseCaseInterface interface {
}
