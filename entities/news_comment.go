package entities

import (
	"gorm.io/gorm"
	"time"
)

type NewsComment struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	UserID    *int           `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AdminID   *int           `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	NewsID    int            `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comment   string         `gorm:"type:text"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	News      News           `gorm:"foreignKey:NewsID;references:ID"`
	User      User           `gorm:"foreignKey:UserID;references:ID"`
	Admin     Admin          `gorm:"foreignKey:AdminID;references:ID"`
}

type NewsCommentRepositoryInterface interface {
	CommentNews(newsComment *NewsComment) error
	GetById(id int) (*NewsComment, error)
	GetByNewsId(newsId int) ([]NewsComment, error)
	UpdateComment(newsComment *NewsComment) error
	DeleteComment(id int) error
}

type NewsCommentUseCaseInterface interface {
	CommentNews(newsComment *NewsComment) error
	GetById(id int) (*NewsComment, error)
	GetByNewsId(newsId int) ([]NewsComment, error)
	UpdateComment(newsComment *NewsComment) error
	DeleteComment(id int) error
}
