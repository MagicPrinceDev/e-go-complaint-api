package response

import "e-complaint-api/entities"

type GetUser struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
	ProfilePhoto    string `json:"profile_photo"`
}

func GetUsersFromEntitiesToResponse(users *entities.User) *GetUser {
	return &GetUser{
		ID:              users.ID,
		Email:           users.Email,
		Name:            users.Name,
		TelephoneNumber: users.TelephoneNumber,
		ProfilePhoto:    users.ProfilePhoto,
	}
}
