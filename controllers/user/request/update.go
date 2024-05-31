package request

import "e-complaint-api/entities"

type UpdateUser struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	TelephoneNumber string `json:"telephone_number" form:"telephone_number"`
}

func (u *UpdateUser) ToEntities() *entities.User {
	return &entities.User{
		Name:            u.Name,
		Email:           u.Email,
		TelephoneNumber: u.TelephoneNumber,
	}
}
