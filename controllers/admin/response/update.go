package response

import "e-complaint-api/entities"

type Update struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	TelephoneNumber string `json:"telephone_number"`
	UpdatedAt       string `json:"update_at,omitempty"`
}

func UpdateUserFromEntitiesToResponse(admin *entities.Admin) *Update {
	return &Update{
		ID:              admin.ID,
		Name:            admin.Name,
		Email:           admin.Email,
		Password:        admin.Password,
		TelephoneNumber: admin.TelephoneNumber,
		UpdatedAt:       admin.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
