package entities

import (
	"gorm.io/gorm"
	"time"
)

type Discussion struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	Comment     string         `gorm:"not null"`
	UserID      *int           `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AdminID     *int           `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ComplaintID string         `gorm:"type:varchar(15);index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User        User           `gorm:"foreignKey:UserID;references:ID"`
	Admin       Admin          `gorm:"foreignKey:AdminID;references:ID"`
	Complaint   Complaint      `gorm:"foreignKey:ComplaintID;references:ID"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type DiscussionRepositoryInterface interface {
	Create(discussion *Discussion) error
	GetById(id int) (*Discussion, error)
	GetByComplaintID(complaintID string) (*[]Discussion, error)
	Update(discussion *Discussion) error
	Delete(id int) error
}

type DiscussionUseCaseInterface interface {
	Create(discussion *Discussion) error
	GetById(id int) (*Discussion, error)
	GetByComplaintID(complaintID string) (*[]Discussion, error)
	Update(discussion *Discussion) error
	Delete(id int) error
}
