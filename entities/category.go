package entities

type Category struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
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
