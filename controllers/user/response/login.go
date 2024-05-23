package response

import "e-complaint-api/entities"

type Login struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func LoginFromEntitiesToResponse(user *entities.User) *Login {
	return &Login{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Token:    user.Token,
	}
}
