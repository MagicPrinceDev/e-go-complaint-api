package response

import "e-complaint-api/entities"

type GetAllUser struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
	ProfilePhoto    string `json:"profile_photo"`
}

func GetAllUsersFromEntitiesToResponse(users []*entities.User) []*GetAllUser {
	var usersResponse []*GetAllUser
	for _, user := range users {
		usersResponse = append(usersResponse, &GetAllUser{
			ID:              user.ID,
			Email:           user.Email,
			Name:            user.Name,
			TelephoneNumber: user.TelephoneNumber,
			ProfilePhoto:    user.ProfilePhoto,
		})
	}
	return usersResponse
}
