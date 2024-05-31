package entities

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	ID         int            `gorm:"primaryKey"`
	AdminID    int            `gorm:"not null"`
	CategoryID int            `gorm:"not null"`
	Title      string         `gorm:"not null"`
	Content    string         `gorm:"not null"`
	TotalLikes int            `gorm:"default:0"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Admin      Admin          `gorm:"foreignKey:AdminID;references:ID"`
	Category   Category       `gorm:"foreignKey:CategoryID;references:ID"`
	Files      []NewsFile     `gorm:"foreignKey:NewsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type NewsRepositoryInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]News, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
}

type NewsUseCaseInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]News, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
}
