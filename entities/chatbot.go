package entities

import (
	"time"

	"gorm.io/gorm"
)

type Chatbot struct {
	ID          int            `gorm:"primaryKey"`
	UserID      int            `gorm:"not null"`
	UserMessage string         `gorm:"not null"`
	BotResponse string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User        User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ChatbotRepositoryInterface interface {
	Create(chatbot *Chatbot) error
	GetHistory(userID int) ([]Chatbot, error)
	ClearHistory(userID int) error
}

type ChatbotOpenAIAPIInterface interface {
	GetChatCompletion(prompt []string, userPrompt string) (string, error)
}

type ChatbotUseCaseInterface interface {
	GetChatCompletion(chatbot *Chatbot) error
	GetHistory(userID int) ([]Chatbot, error)
	ClearHistory(userID int) error
}
