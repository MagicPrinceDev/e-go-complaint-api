package request

import "e-complaint-api/entities"

type Chat struct {
	Message string `json:"message" form:"message" query:"message"`
}

func (c *Chat) ToEntities() *entities.Chatbot {
	return &entities.Chatbot{
		UserMessage: c.Message,
	}
}
