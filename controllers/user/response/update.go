package response

import "e-complaint-api/entities"

type Update struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func UpdateUserFromEntitiesToResponse(user *entities.User) *Update {
	return &Update{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		TelephoneNumber: user.TelephoneNumber,
	}
}
