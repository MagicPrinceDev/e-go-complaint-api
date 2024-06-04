package news

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"time"

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

	if limit != 0 && page != 0 {
		query = query.Limit(limit).Offset((page - 1) * limit)
	}

	if err := query.Preload("Admin").Preload("Category").Preload("Files").Find(&news).Error; err != nil {
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

func (r *NewsRepo) GetByID(id int) (entities.News, error) {
	var news entities.News

	if err := r.DB.Preload("Admin").Preload("Category").Preload("Files").First(&news, id).Error; err != nil {
		return entities.News{}, constants.ErrNewsNotFound
	}

	return news, nil
}

func (r *NewsRepo) Create(news *entities.News) error {
	if err := r.DB.Create(news).Error; err != nil {
		return err
	}

	if err := r.DB.Preload("Admin").Preload("Category").Preload("Files").First(news, news.ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *NewsRepo) Delete(id int) error {
	var news entities.News
	if err := r.DB.First(&news, id).Error; err != nil {
		return constants.ErrNewsNotFound
	}

	news.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := r.DB.Save(&news).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *NewsRepo) Update(news entities.News) (entities.News, error) {
	var oldNews entities.News

	if err := r.DB.First(&oldNews, news.ID).Error; err != nil {
		return entities.News{}, constants.ErrNewsNotFound
	}

	oldNews.Title = news.Title
	oldNews.Content = news.Content
	oldNews.CategoryID = news.CategoryID

	if err := r.DB.Save(&oldNews).Error; err != nil {
		return entities.News{}, constants.ErrInternalServerError
	}

	if err := r.DB.Preload("Admin").Preload("Category").Preload("Files").First(&news, oldNews.ID).Error; err != nil {
		return entities.News{}, constants.ErrInternalServerError
	}

	return news, nil
}
