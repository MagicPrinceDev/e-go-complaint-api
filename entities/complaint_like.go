package entities

type ComplaintLike struct {
	ID          int `json:"id"`
	ComplaintID int `json:"complaint_id"`
	UserID      int `json:"user_id"`
}

type ComplaintLikeRepositoryInterface interface {
	Likes(complaintLike []*ComplaintLike) error
	Unlike(complaintLike []*ComplaintLike) error
}

type ComplaintLikeUseCaseInterface interface {
	Likes(complaintLike []*ComplaintLike) error
	Unlike(complaintLike []*ComplaintLike) error
}
