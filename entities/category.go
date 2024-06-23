package entities

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          int            `gorm:"primaryKey"`
	Name        string         `gorm:"not null;type:varchar(255)"`
	Description string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type CategoryRepositoryInterface interface {
	GetAll() ([]Category, error)
	GetByID(id int) (Category, error)
	CreateCategory(category *Category) (*Category, error)
	UpdateCategory(id int, category *Category) (*Category, error)
	DeleteCategory(id int) error
}

type CategoryUseCaseInterface interface {
	GetAll() ([]Category, error)
	GetByID(id int) (Category, error)
	CreateCategory(category *Category) (*Category, error)
	UpdateCategory(id int, category *Category) (*Category, error)
	DeleteCategory(id int) error
}
