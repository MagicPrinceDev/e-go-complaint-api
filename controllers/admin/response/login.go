package response

import "e-complaint-api/entities"

type Login struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func LoginFromEntitiesToResponse(admin *entities.Admin) *Login {
	return &Login{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Token: admin.Token,
	}
}
