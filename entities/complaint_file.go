package entities

import (
	"time"

	"gorm.io/gorm"
)

type ComplaintFile struct {
	ID          int            `gorm:"primaryKey"`
	ComplaintID string         `gorm:"not null;type:varchar;size:15;"`
	Path        string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
