package entities

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type Complaint struct {
	ID            string             `gorm:"primaryKey;type:varchar(15)"`
	UserID        int                `gorm:"not null"`
	CategoryID    int                `gorm:"not null"`
	RegencyID     string             `gorm:"not null;type:varchar;size:4;"`
	Address       string             `gorm:"not null"`
	Description   string             `gorm:"not null"`
	Status        string             `gorm:"type:enum('Pending', 'Verifikasi', 'On Progress', 'Selesai', 'Ditolak');default:'Pending'"`
	Type          string             `gorm:"type:enum('public', 'private')"`
	Date          time.Time          `gorm:"type:date"`
	TotalLikes    int                `gorm:"default:0"`
	CreatedAt     time.Time          `gorm:"autoCreateTime"`
	UpdatedAt     time.Time          `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt     `gorm:"index"`
	User          User               `gorm:"foreignKey:UserID;references:ID"`
	Regency       Regency            `gorm:"foreignKey:RegencyID;references:ID"`
	Files         []ComplaintFile    `gorm:"foreignKey:ComplaintID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category      Category           `gorm:"foreignKey:CategoryID;references:ID"`
	Process       []ComplaintProcess `gorm:"foreignKey:ComplaintID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Discussion    []Discussion       `gorm:"foreignKey:ComplaintID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ComplaintLike []ComplaintLike    `gorm:"foreignKey:ComplaintID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ComplaintRepositoryInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]Complaint, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
	GetByID(id string) (Complaint, error)
	GetByUserID(userId int) ([]Complaint, error)
	Create(complaint *Complaint) error
	Delete(id string, userId int) error
	AdminDelete(id string) error
	Update(complaint Complaint) (Complaint, error)
	UpdateStatus(id string, status string) error
	GetStatus(id string) (string, error)
	Import(complaints []Complaint) error
	IncreaseTotalLikes(id string) error
	DecreaseTotalLikes(id string) error
	GetComplaintIDsByUserID(userID int) ([]string, error)
}

type ComplaintUseCaseInterface interface {
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]Complaint, error)
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
	GetByID(id string) (Complaint, error)
	GetByUserID(userId int) ([]Complaint, error)
	Create(complaint *Complaint) (Complaint, error)
	Delete(id string, userId int, role string) error
	Update(complaint Complaint) (Complaint, error)
	UpdateStatus(id string, status string) error
	Import(file *multipart.FileHeader) error
	IncreaseTotalLikes(id string) error
	DecreaseTotalLikes(id string) error
	GetComplaintIDsByUserID(userID int) ([]string, error)
}
