package entities

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID              int    `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	Username        string `gorm:"unique;not null"`
	Password        string `gorm:"not null"`
	Email           string `gorm:"unique"`
	TelephoneNumber string
	IsSuperAdmin    bool           `gorm:"default:false"`
	ProfilePhoto    string         `gorm:"default:profiles/default.jpg"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Token           string         `gorm:"-"`
}

type AdminRepositoryInterface interface {
	CreateAccount(admin *Admin) error
	Login(admin *Admin) error
}

type AdminUseCaseInterface interface {
	CreateAccount(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
}
