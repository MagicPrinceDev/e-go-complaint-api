package entities

type District struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	RegencyID string    `json:"regency_id" gorm:"foreignKey:RegencyID;references:ID;type:varchar;size:191"`
	Name      string    `json:"name"`
	Villages  []Village `gorm:"foreignKey:DistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type DistrictRepositoryInterface interface {
	GetByDistrictID(DistrictID string) ([]District, error)
	GetByID(id string) (District, error)
}

type DistrictIndonesiaAreaAPIInterface interface {
	GetDistrictsDataFromAPI([]string) ([]District, error)
}

type DistrictUseCaseInterface interface {
	GetByDistrictID(DistrictID string) ([]District, error)
	GetByID(id string) (District, error)
}
