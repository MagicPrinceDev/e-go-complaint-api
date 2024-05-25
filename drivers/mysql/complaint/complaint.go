package complaint

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type ComplaintRepo struct {
	DB *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) *ComplaintRepo {
	return &ComplaintRepo{DB: db}
}

func (r *ComplaintRepo) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	var complaints []entities.Complaint
	query := r.DB

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("description LIKE ?", "%"+search+"% OR id LIKE ?", "%"+search+"% OR address LIKE ?", "%"+search+"%")
	}

	query = query.Order(sortBy + " " + sortType)

	if err := query.Limit(limit).Offset((page - 1) * limit).Preload("User").Preload("Regency").Preload("Category").Preload("Files").Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}

func (r *ComplaintRepo) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var totalData int64

	query := r.DB.Model(&entities.Complaint{})

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("description LIKE ?", "%"+search+"% OR id LIKE ?", "%"+search+"% OR address LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&totalData).Error; err != nil {
		return entities.Metadata{}, err
	}

	metadata := entities.Metadata{
		TotalData: int(totalData),
	}

	return metadata, nil
}
