package news

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type NewsRepo struct {
	DB *gorm.DB
}

func NewNewsRepo(db *gorm.DB) *NewsRepo {
	return &NewsRepo{DB: db}
}

func (r *NewsRepo) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.News, error) {
	var news []entities.News
	query := r.DB

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query = query.Order(sortBy + " " + sortType)

	if err := query.Limit(limit).Offset((page - 1) * limit).Preload("Admin").Preload("Category").Preload("Files").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}

func (r *NewsRepo) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var totalData int64

	query := r.DB.Model(&entities.News{})

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&totalData).Error; err != nil {
		return entities.Metadata{}, err
	}

	metadata := entities.Metadata{
		TotalData: int(totalData),
	}

	return metadata, nil
}
