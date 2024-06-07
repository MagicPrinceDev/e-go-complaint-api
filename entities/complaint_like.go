package entities

type ComplaintLike struct {
	ID          int `json:"id"`
	ComplaintID int `json:"complaint_id"`
	UserID      int `json:"user_id"`
}
