package response

import "e-complaint-api/entities"

type Get struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	TelephoneNumber string `json:"telephone_number"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
}

func GetFromEntitiesToResponse(admin *entities.Admin) *Get {
	return &Get{
		ID:              admin.ID,
		Name:            admin.Name,
		Email:           admin.Email,
		Password:        admin.Password,
		TelephoneNumber: admin.TelephoneNumber,
		IsSuperAdmin:    admin.IsSuperAdmin,
		ProfilePhoto:    admin.ProfilePhoto,
	}
}
