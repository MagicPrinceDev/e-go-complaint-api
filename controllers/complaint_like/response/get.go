package response

import (
	user_response "e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
)

type Get struct {
	ID        int                   `json:"id"`
	User      user_response.GetUser `json:"user"`
	CreatedAt string                `json:"created_at"`
}

func GetFromEntitiesToResponse(data *entities.ComplaintLike) *Get {
	return &Get{
		ID:        data.ID,
		User:      *user_response.GetUsersFromEntitiesToResponse(&data.User),
		CreatedAt: data.CreatedAt.Format("2 January 2006 15:04:05"),
	}
}
