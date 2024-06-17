package response

import "e-complaint-api/entities"

type Register struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func RegisterFromEntitiesToResponse(user *entities.User) *Register {
	return &Register{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		TelephoneNumber: user.TelephoneNumber,
	}
}
