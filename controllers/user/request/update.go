package request

import "e-complaint-api/entities"

type UpdateUser struct {
	Name            string `json:"name" form:"name"`
	Username        string `json:"username" form:"username" `
	Email           string `json:"email" form:"email"`
	TelephoneNumber string `json:"telephone_number" form:"telephone_number"`
}

func (u *UpdateUser) ToEntities() *entities.User {
	return &entities.User{
		Name:            u.Name,
		Username:        u.Username,
		Email:           u.Email,
		TelephoneNumber: u.TelephoneNumber,
	}
}
