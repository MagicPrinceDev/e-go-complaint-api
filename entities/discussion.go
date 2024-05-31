package entities

type Discussion struct {
	ID int `gorm:"primaryKey"`
}

type DiscussionRepositoryInterface interface {
}

type DiscussionUseCaseInterface interface {
}
