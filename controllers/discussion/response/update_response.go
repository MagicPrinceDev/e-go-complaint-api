package response

import "e-complaint-api/entities"

type UserUpdate struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	ProfilePhoto    string `json:"profile_photo"`
}

type DiscussionUpdate struct {
	ID        int          `json:"id"`
	User      *UserUpdate  `json:"user,omitempty"`
	Admin     *AdminUpdate `json:"admin,omitempty"`
	Comment   string       `json:"comment"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"update_at"`
}

type AdminUpdate struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
}

func FromEntitiesUpdateToResponse(data *entities.Discussion) *DiscussionUpdate {
	var user *UserUpdate
	var admin *AdminUpdate

	if data.AdminID != nil {
		admin = AdminFromEntitiesUpdateToResponse(&data.Admin)
	} else {
		user = UserFromEntitiesUpdateToResponse(&data.User)
	}

	return &DiscussionUpdate{
		ID:        data.ID,
		User:      user,
		Admin:     admin,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt.Format("2 January 2006 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2 January 2006 15:04:05"),
	}
}
func UserFromEntitiesUpdateToResponse(u *entities.User) *UserUpdate {
	return &UserUpdate{
		ID:              u.ID,
		Name:            u.Name,
		TelephoneNumber: u.TelephoneNumber,
		Email:           u.Email,
		ProfilePhoto:    u.ProfilePhoto,
	}
}

func AdminFromEntitiesUpdateToResponse(a *entities.Admin) *AdminUpdate {
	return &AdminUpdate{
		ID:              a.ID,
		Name:            a.Name,
		TelephoneNumber: a.TelephoneNumber,
		Email:           a.Email,
		IsSuperAdmin:    a.IsSuperAdmin,
		ProfilePhoto:    a.ProfilePhoto,
	}
}
