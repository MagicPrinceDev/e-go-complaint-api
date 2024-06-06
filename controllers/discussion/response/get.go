package response

import "e-complaint-api/entities"

type UserGet struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	ProfilePhoto    string `json:"profile_photo"`
	Comment         string `json:"comment,omitempty"`
	UpdatedAt       string `json:"update_at,omitempty"`
}

type DiscussionGet struct {
	ID        int       `json:"id"`
	User      *UserGet  `json:"user,omitempty"`
	Admin     *AdminGet `json:"admin,omitempty"`
	UpdatedAt string    `json:"update_at"`
}

type AdminGet struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
	Comment         string `json:"comment,omitempty"`
	UpdatedAt       string `json:"update_at,omitempty"`
}

func FromEntitiesGetToResponse(data *entities.Discussion) *DiscussionGet {
	var user *UserGet
	var admin *AdminGet

	if data.AdminID != nil {
		admin = &AdminGet{
			ID:              data.Admin.ID,
			Name:            data.Admin.Name,
			TelephoneNumber: data.Admin.TelephoneNumber,
			Email:           data.Admin.Email,
			IsSuperAdmin:    data.Admin.IsSuperAdmin,
			ProfilePhoto:    data.Admin.ProfilePhoto,
			Comment:         data.Comment,
			UpdatedAt:       data.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	} else {
		user = &UserGet{
			ID:              data.User.ID,
			Name:            data.User.Name,
			TelephoneNumber: data.User.TelephoneNumber,
			Email:           data.User.Email,
			ProfilePhoto:    data.User.ProfilePhoto,
			Comment:         data.Comment,
			UpdatedAt:       data.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &DiscussionGet{
		ID:        data.ID,
		User:      user,
		Admin:     admin,
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
