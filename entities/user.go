package entities

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int    `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	Email           string `gorm:"unique"`
	Password        string `gorm:"not null"`
	TelephoneNumber string
	ProfilePhoto    string         `gorm:"default:profile-photos/default.jpg"`
	Token           string         `gorm:"-"`
	Discussion      []Discussion   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NewsComment     []NewsComment  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Otp             string         `gorm:"default:null"`
	OtpExpiredAt    time.Time      `gorm:"default:null"`
	EmailVerified   bool           `gorm:"default:false"`
	ForgotVerified  bool           `gorm:"default:false"`
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
