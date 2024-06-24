package response

import (
	"e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
)

type Get struct {
	ID          int               `json:"id"`
	User        *response.GetUser `json:"user,omitempty"`
	UserMessage string            `json:"user_message,omitempty"`
	BotResponse string            `json:"bot_response"`
	CreatedAt   string            `json:"created_at"`
}

func GetFromEntitiesToResponse(data *entities.Chatbot) *Get {
	return &Get{
		ID:          data.ID,
		User:        response.GetUsersFromEntitiesToResponse(&data.User),
		UserMessage: data.UserMessage,
		BotResponse: data.BotResponse,
		CreatedAt:   data.CreatedAt.Format("2 January 2006 15:04:05"),
	}
}
