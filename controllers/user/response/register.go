package response

import "e-complaint-api/entities"

type Register struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func RegisterFromEntitiesToResponse(user *entities.User) *Register {
	return &Register{
		Name:            user.Name,
		Username:        user.Username,
		Email:           user.Email,
		TelephoneNumber: user.TelephoneNumber,
	}
}
