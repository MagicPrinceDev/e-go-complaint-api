package chatbot

import (
	"e-complaint-api/entities"
	"strconv"
)

type ChatbotUseCase struct {
	chatbotRepo   entities.ChatbotRepositoryInterface
	faqRepo       entities.FaqRepositoryInterface
	complaintRepo entities.ComplaintRepositoryInterface
	OpenAIAPI     entities.ChatbotOpenAIAPIInterface
}

func NewChatbotUseCase(chatbotRepo entities.ChatbotRepositoryInterface, faqRepo entities.FaqRepositoryInterface, complaintRepo entities.ComplaintRepositoryInterface, OpenAIAPI entities.ChatbotOpenAIAPIInterface) *ChatbotUseCase {
	return &ChatbotUseCase{
		chatbotRepo:   chatbotRepo,
		faqRepo:       faqRepo,
		complaintRepo: complaintRepo,
		OpenAIAPI:     OpenAIAPI,
	}
}

func (u *ChatbotUseCase) GetChatCompletion(chatbot *entities.Chatbot) error {
	faq, err := u.faqRepo.GetAll()
	if err != nil {
		return err
	}

	userComplaint, err := u.complaintRepo.GetByUserID(chatbot.UserID)
	if err != nil {
		return err
	}

	var prompt []string
	message := "FAQ (Frequently Asked Questions):\n\n"
	for i, f := range faq {
		stringIdx := strconv.Itoa(i + 1)
		faqMessage := stringIdx + ".) Q: " + f.Question + "\nA: " + f.Answer + "\n\n"
		message += faqMessage
	}
	prompt = append(prompt, message)

	if userComplaint != nil {
		message = "Riwayat Aduan User:\n\n"
		for i, c := range userComplaint {
			stringIdx := strconv.Itoa(i + 1)
			complaintMessage := stringIdx + ".)ID Complaint: " + c.ID + "\n" + "Deskripsi: " + c.Description + "\n" + "Status: " + c.Status + "\n" + "Tanggal Aduan dibuat: " + c.CreatedAt.Format("3 January 2006 15:04:05") + "\n\n"
			message += complaintMessage
		}
		prompt = append(prompt, message)
	}

	prompt = append(prompt, "Tolong anda sebagai Customer Service untuk memberikan respon kepada user berdasarkan FAQ dan Riwayat Aduan User di atas")

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

func (u *ChatbotUseCase) ClearHistory(userID int) error {
	err := u.chatbotRepo.ClearHistory(userID)
	if err != nil {
		return err
	}

	return nil
}
