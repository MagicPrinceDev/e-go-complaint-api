package request

import "e-complaint-api/entities"

type Login struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *Login) ToEntities() *entities.Admin {
	return &entities.Admin{
		Email:    r.Email,
		Password: r.Password,
	}
}
