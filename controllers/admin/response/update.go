package response

import "e-complaint-api/entities"

type Update struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func UpdateUserFromEntitiesToResponse(admin *entities.Admin) *Update {
	return &Update{
		ID:              admin.ID,
		Name:            admin.Name,
		Email:           admin.Email,
		TelephoneNumber: admin.TelephoneNumber,
	}
}
