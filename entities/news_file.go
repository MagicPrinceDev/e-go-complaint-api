package entities

import (
	"time"

	"gorm.io/gorm"
)

type NewsFile struct {
	ID        int            `gorm:"primaryKey"`
	NewsID    int            `gorm:"not null"`
	Path      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
