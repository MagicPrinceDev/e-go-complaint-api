package request

import "e-complaint-api/entities"

type CreateDiscussion struct {
	Comment string `json:"comment" form:"comment"`
	AdminID *int   `json:"admin_id" form:"admin_id"`
	UserID  *int   `json:"user_id"`
}

func (r *CreateDiscussion) ToEntities(userID int, complaintID string, role string) *entities.Discussion {
	var adminID *int
	var userIDPtr *int

	if role == "admin" || role == "super_admin" {
		adminID = &userID
	} else {
		userIDPtr = &userID
	}

	return &entities.Discussion{
		ComplaintID: complaintID,
		Comment:     r.Comment,
		UserID:      userIDPtr,
		AdminID:     adminID,
	}
}
