package response

import "e-complaint-api/entities"

type Get struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetFromEntitiesToResponse(data *entities.Category) *Get {
	return &Get{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
}
