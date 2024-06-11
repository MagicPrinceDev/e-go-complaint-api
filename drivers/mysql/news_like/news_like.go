package news_like

import (
	"e-complaint-api/entities"
	"errors"
	"gorm.io/gorm"
)

type NewsLikeRepo struct {
	DB *gorm.DB
}

func NewNewsLikeRepo(db *gorm.DB) *NewsLikeRepo {
	return &NewsLikeRepo{DB: db}
}

func (r *NewsLikeRepo) FindByUserAndNews(userID int, newsID int) (*entities.NewsLike, error) {
	var newsLike entities.NewsLike
	result := r.DB.Where("user_id = ? AND news_id = ?", userID, newsID).First(&newsLike)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &newsLike, nil
}

func (r *NewsLikeRepo) Likes(newsLike *entities.NewsLike) error {
	if err := r.DB.Create(newsLike).Error; err != nil {
		return err
	}

	return nil
}

func (r *NewsLikeRepo) Unlike(newsLike *entities.NewsLike) error {
	db := r.DB
	if err := db.Where("user_id = ? AND news_id = ?", newsLike.UserID, newsLike.NewsID).Delete(&entities.NewsLike{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *NewsLikeRepo) IncreaseTotalLikes(id string) error {
	if err := r.DB.Model(&entities.News{}).Where("id = ?", id).Update("total_likes", gorm.Expr("total_likes + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (r *NewsLikeRepo) DecreaseTotalLikes(id string) error {
	if err := r.DB.Model(&entities.News{}).Where("id = ?", id).Update("total_likes", gorm.Expr("total_likes - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
