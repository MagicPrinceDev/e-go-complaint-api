package entities

type Regency struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type RegencyRepositoryInterface interface {
	GetAll() ([]Regency, error)
}

type RegencyIndonesiaAreaAPIInterface interface {
	GetRegenciesDataFromAPI() ([]Regency, error)
}

type RegencyUseCaseInterface interface {
	GetAll() ([]Regency, error)
}
