package entities

type Category struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

type CategoryRepositoryInterface interface {
	GetAll() ([]Category, error)
}

type CategoryUseCaseInterface interface {
	GetAll() ([]Category, error)
}
