package entities

import (
	"mime/multipart"
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

type NewsFileRepositoryInterface interface {
	Create(newsFiles []*NewsFile) error
	// DeleteByNewsID(newsID int) error
}

type NewsFileGCSAPIInterface interface {
	Upload(files []*multipart.FileHeader) ([]string, error)
	// Delete(filePaths []string) error
}

type NewsFileUseCaseInterface interface {
	Create(files []*multipart.FileHeader, newsID int) ([]NewsFile, error)
	// DeleteByNewsID(newsID int) error
}
