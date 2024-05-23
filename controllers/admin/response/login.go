package response

import "e-complaint-api/entities"

type Login struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func LoginFromEntitiesToResponse(admin *entities.Admin) *Login {
	return &Login{
		ID:       admin.ID,
		Username: admin.Username,
		Token:    admin.Token,
	}
}
