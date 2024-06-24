package response

import "e-complaint-api/entities"

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	ProfilePhoto    string `json:"profile_photo"`
}

type Discussion struct {
	ID        int    `json:"id"`
	User      *User  `json:"user,omitempty"`
	Admin     *Admin `json:"admin,omitempty"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

type Admin struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
}

func FromEntitiesToResponse(data *entities.Discussion) *Discussion {
	var user *User
	var admin *Admin

	if data.AdminID != nil {
		admin = AdminFromEntitiesToResponse(&data.Admin)
	} else if data.UserID != nil {
		user = UserFromEntitiesToResponse(&data.User)
	}

	return &Discussion{
		ID:        data.ID,
		User:      user,
		Admin:     admin,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt.Format("2 January 2006 15:04:05"),
	}
}

func UserFromEntitiesToResponse(u *entities.User) *User {
	return &User{
		ID:              u.ID,
		Name:            u.Name,
		TelephoneNumber: u.TelephoneNumber,
		Email:           u.Email,
		ProfilePhoto:    u.ProfilePhoto,
	}
}

func AdminFromEntitiesToResponse(a *entities.Admin) *Admin {
	return &Admin{
		ID:              a.ID,
		Name:            a.Name,
		TelephoneNumber: a.TelephoneNumber,
		Email:           a.Email,
		IsSuperAdmin:    a.IsSuperAdmin,
		ProfilePhoto:    a.ProfilePhoto,
	}
}
