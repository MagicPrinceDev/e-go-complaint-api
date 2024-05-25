package entities

import (
	"time"

	"gorm.io/gorm"
)

type Complaint struct {
	ID          string          `gorm:"primaryKey;length:15"`
	UserID      int             `gorm:"not null"`
	CategoryID  int             `gorm:"not null"`
	Description string          `gorm:"not null"`
	RegencyID   string          `gorm:"not null;type:varchar;size:4"`
	Address     string          `gorm:"not null"`
	Status      string          `gorm:"enum('pending', 'verifikasi', 'on progress', 'selesai', 'ditolak')"`
	Type        string          `gorm:"enum('public', 'private')"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
	User        User            `gorm:"foreignKey:UserID;references:ID"`
	Regency     Regency         `gorm:"foreignKey:RegencyID;references:ID"`
	Files       []ComplaintFile `gorm:"foreignKey:ComplaintID;references:ID"`
	Category    Category        `gorm:"foreignKey:CategoryID;references:ID"`
}

type ComplaintRepositoryInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]Complaint, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
}

type ComplaintUseCaseInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]Complaint, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
}
