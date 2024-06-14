package response

import "e-complaint-api/entities"

type GetSimple struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
}

func GetSimpleFromEntitiesToResponse(admin *entities.Admin) *GetSimple {
	return &GetSimple{
		ID:              admin.ID,
		Name:            admin.Name,
		Email:           admin.Email,
		TelephoneNumber: admin.TelephoneNumber,
		IsSuperAdmin:    admin.IsSuperAdmin,
		ProfilePhoto:    admin.ProfilePhoto,
	}
}
