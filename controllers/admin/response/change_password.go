package response

import "e-complaint-api/entities"

type ChangePassword struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ChangePasswordFromEntitiesToResponse(admin *entities.Admin) *ChangePassword {
	return &ChangePassword{
		ID:       admin.ID,
		Username: admin.Username,
		Email:    admin.Email,
	}
}
