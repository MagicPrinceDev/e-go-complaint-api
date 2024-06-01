package request

import "e-complaint-api/entities"

type CreateAccount struct {
	Name            string `form:"name" json:"name"`
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	TelephoneNumber string `form:"telephone_number" json:"telephone_number"`
}

func (r *CreateAccount) ToEntities() *entities.Admin {
	return &entities.Admin{
		Email:           r.Email,
		Password:        r.Password,
		Name:            r.Name,
		TelephoneNumber: r.TelephoneNumber,
	}
}
