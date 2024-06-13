package request

import "e-complaint-api/entities"

type CommentNews struct {
	AdminID *int
	UserID  *int
	Comment string `json:"comment" form:"comment"`
	NewsID  int
}

func (r *CommentNews) ToEntities(userID int, NewsID int, role string) *entities.NewsComment {
	var adminID *int
	var userIDPtr *int

	if role == "admin" || role == "super_admin" {
		adminID = &userID
	} else {
		userIDPtr = &userID
	}

	return &entities.NewsComment{
		NewsID:  NewsID,
		Comment: r.Comment,
		UserID:  userIDPtr,
		AdminID: adminID,
	}
}
