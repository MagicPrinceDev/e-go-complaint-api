package request

import "e-complaint-api/entities"

type Register struct {
	Name            string `form:"name"`
	Email           string `form:"email"`
	TelephoneNumber string `form:"telephone_number"`
	Password        string `form:"password"`
}

func (r *Register) ToEntities() *entities.User {
	return &entities.User{
		Name:            r.Name,
		Email:           r.Email,
		TelephoneNumber: r.TelephoneNumber,
		Password:        r.Password,
	}
}
