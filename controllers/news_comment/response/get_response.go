package response

import "e-complaint-api/entities"

type UserGet struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	ProfilePhoto    string `json:"profile_photo"`
}

type AdminGet struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
}

type NewsGet struct {
	ID        int       `json:"id"`
	User      *UserGet  `json:"user,omitempty"`
	Admin     *AdminGet `json:"admin,omitempty"`
	Comment   string    `json:"comment"`
	UpdatedAt string    `json:"update_at"`
}

func FromEntitiesGetToResponse(data *entities.NewsComment) *NewsGet {
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
		}
	} else {
		user = &UserGet{
			ID:              data.User.ID,
			Name:            data.User.Name,
			TelephoneNumber: data.User.TelephoneNumber,
			Email:           data.User.Email,
			ProfilePhoto:    data.User.ProfilePhoto,
		}
	}

	return &NewsGet{
		ID:        data.ID,
		User:      user,
		Admin:     admin,
		Comment:   data.Comment,
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
