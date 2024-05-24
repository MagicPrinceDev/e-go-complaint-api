package entities

type Village struct {
	ID         string `json:"id" gorm:"primaryKey"`
	DistrictID string `json:"district_id" gorm:"foreignKey:DistrictID;references:ID;type:varchar;size:191"`
	Name       string `json:"name"`
}

type VillageRepositoryInterface interface {
	GetByDistrictID(districtID string) ([]Village, error)
	GetByID(id string) (Village, error)
}

type VillageIndonesiaAreaAPIInterface interface {
	GetVillagesDataFromAPI([]string) ([]Village, error)
}

type VillageUseCaseInterface interface {
	GetByDistrictID(districtID string) ([]Village, error)
	GetByID(id string) (Village, error)
}
