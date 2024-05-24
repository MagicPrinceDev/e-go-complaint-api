package entities

type Regency struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Districts []District `gorm:"foreignKey:RegencyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type RegencyRepositoryInterface interface {
	GetByRegencyID(regencyID string) ([]Regency, error)
	GetByID(id string) (Regency, error)
}

type RegencyIndonesiaAreaAPIInterface interface {
	GetRegenciesDataFromAPI() ([]Regency, error)
}

type RegencyUseCaseInterface interface {
	GetByRegencyID(regencyID string) ([]Regency, error)
	GetByID(id string) (Regency, error)
}
