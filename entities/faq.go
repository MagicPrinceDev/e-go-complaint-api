package entities

type Faq struct {
	ID       int    `gorm:"primaryKey"`
	Question string `gorm:"not null"`
	Answer   string `gorm:"not null"`
}

type FaqRepositoryInterface interface {
	GetAll() ([]Faq, error)
}

type FaqUseCaseInterface interface {
	GetAll() ([]Faq, error)
}
