package entities

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int            `gorm:"primaryKey"`
	Name            string         `gorm:"not null;type:varchar(255)"`
	Email           string         `gorm:"unique;not null;type:varchar(255)"`
	Password        string         `gorm:"not null;type:varchar(255)"`
	TelephoneNumber string         `gorm:"not null;type:varchar(20)"`
	ProfilePhoto    string         `gorm:"default:profile-photos/default.jpg;type:varchar(255)"`
	Token           string         `gorm:"-"`
	Otp             string         `gorm:"default:null;type:varchar(5)"`
	OtpExpiredAt    time.Time      `gorm:"default:null"`
	EmailVerified   bool           `gorm:"default:false"`
	ForgotVerified  bool           `gorm:"default:false"`
	Discussion      []Discussion   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NewsComment     []NewsComment  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type UserRepositoryInterface interface {
	Register(user *User) error
	Login(user *User) error
	GetAllUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(id int, user *User) error
	UpdateProfilePhoto(id int, profilePhoto string) error
	Delete(id int) error
	UpdatePassword(id int, newPassword string) error
	SendOTP(email, otp string) error
	VerifyOTPRegister(email, otp string) error
	VerifyOTPForgotPassword(email, otp string) error
	UpdatePasswordForgot(email, newPassword string) error
}

type MailTrapAPIInterface interface {
	SendOTP(email, otp, otp_type string) error
}

type UserGCSAPIInterface interface {
	Upload(files []*multipart.FileHeader) ([]string, error)
}

type UserUseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
	GetAllUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(id int, user *User) (User, error)
	UpdateProfilePhoto(id int, profilePhoto *multipart.FileHeader) error
	Delete(id int) error
	UpdatePassword(id int, newPassword string) error
	SendOTP(email, otp_type string) error
	VerifyOTP(email, otp, otp_type string) error
	UpdatePasswordForgot(email, newPassword string) error
}
