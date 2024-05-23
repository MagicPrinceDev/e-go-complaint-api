package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int    `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	Username        string `gorm:"unique;not null"`
	Password        string `gorm:"not null"`
	Email           string `gorm:"unique"`
	TelephoneNumber string
	ProfilePhoto    string         `gorm:"default:profiles/default.jpg"`
	Token           string         `gorm:"-"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type UserRepositoryInterface interface {
	Register(user *User) error
	Login(user *User) error
}

type MailTrapAPIInterface interface {
	SendEmail(emailReceiver string) error
}

type UserUseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
}
