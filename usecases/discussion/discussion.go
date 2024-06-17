package discussion

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"strconv"
)

type DiscussionUseCase struct {
	discussionRepo entities.DiscussionRepositoryInterface
	faqRepo        entities.FaqRepositoryInterface
	openAIAPI      entities.DiscussionOpenAIAPIInterface
}

func NewDiscussionUseCase(discussionRepo entities.DiscussionRepositoryInterface, faqRepo entities.FaqRepositoryInterface, openAIAPI entities.DiscussionOpenAIAPIInterface) *DiscussionUseCase {
	return &DiscussionUseCase{
		discussionRepo: discussionRepo,
		faqRepo:        faqRepo,
		openAIAPI:      openAIAPI,
	}
}

func (u *DiscussionUseCase) Create(discussion *entities.Discussion) error {
	if discussion.Comment == "" {
		return constants.ErrCommentCannotBeEmpty
	}

	err := u.discussionRepo.Create(discussion)
	if err != nil {
		return err
	}

	return nil
}

func (u *DiscussionUseCase) GetById(id int) (*entities.Discussion, error) {
	discussion, err := u.discussionRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return discussion, nil
}

func (u *DiscussionUseCase) GetByComplaintID(complaintID string) (*[]entities.Discussion, error) {
	discussions, err := u.discussionRepo.GetByComplaintID(complaintID)
	if err != nil {
		return nil, err
	}

	return discussions, nil
}

func (u *DiscussionUseCase) Update(discussion *entities.Discussion) error {
	if discussion.Comment == "" {
		return constants.ErrCommentCannotBeEmpty
	}
	err := u.discussionRepo.Update(discussion)
	if err != nil {
		return err
	}

	return nil
}

func (u *DiscussionUseCase) Delete(id int) error {
	err := u.discussionRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *DiscussionUseCase) GetAnswerRecommendation(complaintID string) (string, error) {
	discussions, err := u.GetByComplaintID(complaintID)
	if err != nil {
		return "", err
	}

	faq, err := u.faqRepo.GetAll()
	if err != nil {
		return "", err
	}

	var prompt []string
	message := "FAQ (Frequently Asked Questions):\n\n"
	for i, f := range faq {
		stringIdx := strconv.Itoa(i + 1)
		faqMessage := stringIdx + ".) Q: " + f.Question + "\nA: " + f.Answer + "\n\n"
		message += faqMessage
	}
	prompt = append(prompt, message)

	if discussions != nil {
		message = "Diskusi Terkait Aduan User:\n"
		for i, d := range *discussions {
			if d.UserID != nil {
				stringIdx := strconv.Itoa(i + 1)
				discussionMessage := stringIdx + ".)User: " + d.Comment + "\n"
				message += discussionMessage
			} else {
				stringIdx := strconv.Itoa(i + 1)
				discussionMessage := stringIdx + ".)Admin: " + d.Comment + "\n"
				message += discussionMessage
			}
		}
		prompt = append(prompt, message)
	}

	prompt = append(prompt, "Anda sebagai admin, berikan respon jawaban terhadap diskusi oleh user di atas yang belum terjawab oleh Admin. Jika ada pertanyaan yang sama atau mirip, jawaban yang diberikan cukup satu kali saja. Jawaban yang anda berikan disesuaikan dengan FAQ yang telah disediakan(Menyocokkan pertanyaan pada Q lalu jawab dengan A yang sesuai).")

	botResponse, err := u.openAIAPI.GetChatCompletion(prompt, "")
	if err != nil {
		return "", err
	}

	return botResponse, nil
}
