package entities

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID              int            `gorm:"primaryKey"`
	Name            string         `gorm:"not null;type:varchar(255)"`
	Email           string         `gorm:"unique;type:varchar(255)"`
	Password        string         `gorm:"not null;type:varchar(255)"`
	TelephoneNumber string         `gorm:"type:varchar(20)"`
	IsSuperAdmin    bool           `gorm:"default:false"`
	ProfilePhoto    string         `gorm:"default:profile-photos/admin-default.jpg;type:varchar(255)"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Token           string         `gorm:"-"`
	Discussion      []Discussion   `gorm:"foreignKey:AdminID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NewsComment     []NewsComment  `gorm:"foreignKey:AdminID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type AdminRepositoryInterface interface {
	CreateAccount(admin *Admin) error
	Login(admin *Admin) error
	GetAllAdmins() ([]*Admin, error)
	GetAdminByID(id int) (*Admin, error)
	DeleteAdmin(id int) error
	UpdateAdmin(id int, user *Admin) error
	GetAdminByEmail(email string) (*Admin, error)
}

type AdminUseCaseInterface interface {
	CreateAccount(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
	GetAllAdmins() ([]Admin, error)
	GetAdminByID(id int) (*Admin, error)
	DeleteAdmin(id int) error
	UpdateAdmin(id int, user *Admin) (Admin, error)
}
