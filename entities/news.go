package entities

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	ID         int            `gorm:"primaryKey"`
	AdminID    int            `gorm:"not null"`
	CategoryID int            `gorm:"not null"`
	Title      string         `gorm:"not null;type:varchar(255)"`
	Content    string         `gorm:"not null"`
	TotalLikes int            `gorm:"default:0"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Admin      Admin          `gorm:"foreignKey:AdminID;references:ID"`
	Category   Category       `gorm:"foreignKey:CategoryID;references:ID"`
	Files      []NewsFile     `gorm:"foreignKey:NewsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NewsLike   []NewsLike     `gorm:"foreignKey:NewsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments   []NewsComment  `gorm:"foreignKey:NewsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type NewsRepositoryInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]News, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
	GetByID(id int) (News, error)
	Create(news *News) error
	Delete(id int) error
	Update(news News) (News, error)
}

type NewsUseCaseInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]News, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
	GetByID(id int) (News, error)
	Create(news *News) (News, error)
	Delete(id int) error
	Update(news News) (News, error)
}
