package response

import "e-complaint-api/entities"

type CreateAccount struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func CreateAccountFromEntitiesToResponse(admin *entities.Admin) *CreateAccount {
	return &CreateAccount{
		ID:              admin.ID,
		Name:            admin.Name,
		Email:           admin.Email,
		TelephoneNumber: admin.TelephoneNumber,
	}
}
