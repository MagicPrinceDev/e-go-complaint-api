package response

import (
	"e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
)

type Get struct {
	ID          int              `json:"id"`
	User        response.GetUser `json:"user"`
	UserMessage string           `json:"user_message"`
	BotResponse string           `json:"bot_response"`
	CreatedAt   string           `json:"created_at"`
}

func GetFromEntitiesToResponse(data *entities.Chatbot) *Get {
	return &Get{
		ID:          data.ID,
		User:        *response.GetUsersFromEntitiesToResponse(&data.User),
		UserMessage: data.UserMessage,
		BotResponse: data.BotResponse,
		CreatedAt:   data.CreatedAt.Format("3 January 2006 15:04:05"),
	}
}
