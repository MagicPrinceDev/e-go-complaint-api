package request

import "e-complaint-api/entities"

type CreateAccount struct {
	Name            string `form:"name"`
	Username        string `form:"username"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	TelephoneNumber string `form:"telephone_number"`
}

func (r *CreateAccount) ToEntities() *entities.Admin {
	return &entities.Admin{
		Username:        r.Username,
		Email:           r.Email,
		Password:        r.Password,
		Name:            r.Name,
		TelephoneNumber: r.TelephoneNumber,
	}
}
