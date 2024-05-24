package response

import "e-complaint-api/entities"

type GetAllAdmins struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
}

func GetAdminsFromEntitiesToResponse(admin *entities.Admin) *GetAllAdmins {
	return &GetAllAdmins{
		ID:              admin.ID,
		Name:            admin.Name,
		Username:        admin.Username,
		Email:           admin.Email,
		TelephoneNumber: admin.TelephoneNumber,
		IsSuperAdmin:    admin.IsSuperAdmin,
		ProfilePhoto:    admin.ProfilePhoto,
	}
}
