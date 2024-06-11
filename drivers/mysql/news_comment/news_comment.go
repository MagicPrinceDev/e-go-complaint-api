package news_comment

import (
	"e-complaint-api/entities"
	"gorm.io/gorm"
)

type NewsComment struct {
	DB *gorm.DB
}

func NewNewsComment(db *gorm.DB) *NewsComment {
	return &NewsComment{DB: db}
}

func (r *NewsComment) CommentNews(newsComment *entities.NewsComment) error {
	if err := r.DB.Create(newsComment).Error; err != nil {
		return err
	}
	return nil
}

func (r *NewsComment) GetById(id int) (*entities.NewsComment, error) {
	var newsComment entities.NewsComment
	if err := r.DB.Preload("User").Preload("Admin").Preload("News").First(&newsComment, id).Error; err != nil {
		return nil, err
	}
	return &newsComment, nil
}
