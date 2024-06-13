package chatbot

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type ChatbotRepo struct {
	DB *gorm.DB
}

func NewChatbotRepo(db *gorm.DB) *ChatbotRepo {
	return &ChatbotRepo{DB: db}
}

func (r *ChatbotRepo) Create(chatbot *entities.Chatbot) error {
	if err := r.DB.Create(chatbot).Error; err != nil {
		return err
	}

	if err := r.DB.Preload("User").First(chatbot, chatbot.ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *ChatbotRepo) GetHistory(userID int) ([]entities.Chatbot, error) {
	var chatbots []entities.Chatbot
	if err := r.DB.Preload("User").Where("user_id = ?", userID).Find(&chatbots).Error; err != nil {
		return nil, err
	}
	return chatbots, nil
}
