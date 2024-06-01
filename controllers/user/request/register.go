package request

import "e-complaint-api/entities"

type Register struct {
	Name            string `form:"name" json:"name"`
	Email           string `form:"email" json:"email"`
	TelephoneNumber string `form:"telephone_number" json:"telephone_number"`
	Password        string `form:"password" json:"password"`
}

func (r *Register) ToEntities() *entities.User {
	return &entities.User{
		Name:            r.Name,
		Email:           r.Email,
		TelephoneNumber: r.TelephoneNumber,
		Password:        r.Password,
	}
}
