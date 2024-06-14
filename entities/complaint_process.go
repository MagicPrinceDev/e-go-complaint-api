package entities

import (
	"time"

	"gorm.io/gorm"
)

type ComplaintProcess struct {
	ID          int            `gorm:"primaryKey"`
	ComplaintID string         `gorm:"not null;type:varchar;size:15;"`
	AdminID     int            `gorm:"not null"`
	Status      string         `gorm:"not null;type:enum('Pending', 'Verifikasi', 'On Progress', 'Selesai', 'Ditolak')"`
	Message     string         `gorm:"type:text"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Admin       Admin          `gorm:"foreignKey:AdminID;references:ID"`
}

type ComplaintProcessRepositoryInterface interface {
	Create(complaintProcesses *ComplaintProcess) error
	GetByComplaintID(complaintID string) ([]ComplaintProcess, error)
	Update(complaintProcesses *ComplaintProcess) error
	Delete(complaintID string, complaintProcessID int) (string, error)
}

type ComplaintProcessUseCaseInterface interface {
	Create(complaintProcesses *ComplaintProcess) (ComplaintProcess, error)
	GetByComplaintID(complaintID string) ([]ComplaintProcess, error)
	Update(complaintProcesses *ComplaintProcess) (ComplaintProcess, error)
	Delete(complaintID string, complaintProcessID int) (string, error)
}
