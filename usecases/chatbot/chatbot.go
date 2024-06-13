package chatbot

import (
	"e-complaint-api/entities"
	"strconv"
)

type ChatbotUseCase struct {
	chatbotRepo entities.ChatbotRepositoryInterface
	faqRepo     entities.FaqRepositoryInterface
	OpenAIAPI   entities.ChatbotOpenAIAPIInterface
}

func NewChatbotUseCase(chatbotRepo entities.ChatbotRepositoryInterface, faqRepo entities.FaqRepositoryInterface, OpenAIAPI entities.ChatbotOpenAIAPIInterface) *ChatbotUseCase {
	return &ChatbotUseCase{
		chatbotRepo: chatbotRepo,
		faqRepo:     faqRepo,
		OpenAIAPI:   OpenAIAPI,
	}
}

func (u *ChatbotUseCase) GetChatCompletion(chatbot *entities.Chatbot) error {
	faq, err := u.faqRepo.GetAll()
	if err != nil {
		return err
	}

	var prompt []string
	message := "Tolong jawab pertanyaan user dengan mereferensikan FAQ berikut (Cocokkan Q yang diberikan user dengan Q yang ada di FAQ lalu jawab dengan A yang sesuai):\n\n"
	for i, f := range faq {
		stringIdx := strconv.Itoa(i + 1)
		faqMessage := stringIdx + ".) Q: " + f.Question + "\nA: " + f.Answer + "\n\n"
		message += faqMessage
	}
	prompt = append(prompt, message)

	botResponse, err := u.OpenAIAPI.GetChatCompletion(prompt, chatbot.UserMessage)
	if err != nil {
		return err
	}

	(*chatbot).BotResponse = botResponse

	err = u.chatbotRepo.Create(chatbot)
	if err != nil {
		return err
	}

	return nil
}

func (u *ChatbotUseCase) GetHistory(userID int) ([]entities.Chatbot, error) {
	chatbots, err := u.chatbotRepo.GetHistory(userID)
	if err != nil {
		return nil, err
	}

	return chatbots, nil
}
