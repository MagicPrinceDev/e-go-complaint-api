package entities

type Faq struct {
	ID       int    `gorm:"primaryKey"`
	Question string `gorm:"not null;type:text"`
	Answer   string `gorm:"not null;type:text"`
}

type FaqRepositoryInterface interface {
	GetAll() ([]Faq, error)
}

type FaqUseCaseInterface interface {
	GetAll() ([]Faq, error)
}
