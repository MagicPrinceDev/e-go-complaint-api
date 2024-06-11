package entities

import "time"

type NewsLike struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	NewsID    int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	News      News      `gorm:"foreignKey:NewsID;references:ID"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
}

type NewsLikeRepositoryInterface interface {
	FindByUserAndNews(userID int, newsID int) (*NewsLike, error)
	Likes(newsLike *NewsLike) error
	Unlike(newsLike *NewsLike) error
	IncreaseTotalLikes(id string) error
	DecreaseTotalLikes(id string) error
}

type NewsLikeUseCaseInterface interface {
	ToggleLike(complaintLike *NewsLike) (string, error)
	IncreaseTotalLikes(id string) error
	DecreaseTotalLikes(id string) error
}
