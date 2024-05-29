package response

import "e-complaint-api/entities"

type Update struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func UpdateUserFromEntitiesToResponse(user *entities.Admin) *Update {
	return &Update{
		ID:              user.ID,
		Name:            user.Name,
		Username:        user.Username,
		Email:           user.Email,
		TelephoneNumber: user.TelephoneNumber,
	}
}
