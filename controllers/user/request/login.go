package request

import "e-complaint-api/entities"

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (r *Login) ToEntities() *entities.User {
	return &entities.User{
		Username: r.Username,
		Password: r.Password,
	}
}
