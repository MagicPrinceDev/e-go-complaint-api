package response

import "e-complaint-api/entities"

type Login struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func LoginFromEntitiesToResponse(user *entities.User) *Login {
	return &Login{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
}
