package response

import "e-complaint-api/entities"

type Get struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func GetFromEntitiesToResponse(data *entities.User) *Get {
	return &Get{
		ID:       data.ID,
		Name:     data.Name,
		Username: data.Username,
	}
}
