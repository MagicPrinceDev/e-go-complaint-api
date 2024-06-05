package response

import "e-complaint-api/entities"

type GetDiscussion struct {
	ID        int    `json:"id"`
	User      *User  `json:"user,omitempty"`
	Admin     *Admin `json:"admin,omitempty"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

func FromEntitiesGetToResponse(data *entities.Discussion) *GetDiscussion {
	var user *User
	var admin *Admin

	if data.AdminID != nil {
		admin = AdminFromEntitiesToResponse(&data.Admin)
	} else {
		user = UserFromEntitiesToResponse(&data.User)
	}

	return &GetDiscussion{
		ID:        data.ID,
		User:      user,
		Admin:     admin,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
