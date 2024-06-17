package chatbot

import (
	"e-complaint-api/controllers/base"
	chatbot_request "e-complaint-api/controllers/chatbot/request"
	chatbot_response "e-complaint-api/controllers/chatbot/response"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ChatbotController struct {
	chatbotUseCase entities.ChatbotUseCaseInterface
}

func NewChatbotController(chatbotUseCase entities.ChatbotUseCaseInterface) *ChatbotController {
	return &ChatbotController{
		chatbotUseCase: chatbotUseCase,
	}
}

func (cc *ChatbotController) GetHistory(c echo.Context) error {
	userID, _ := utils.GetIDFromJWT(c)

	history, err := cc.chatbotUseCase.GetHistory(userID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var historyResponse []*chatbot_response.Get
	historyResponse = append(historyResponse, &chatbot_response.Get{
		BotResponse: "Halo, saya adalah KAVA(KeluhProv AI Virtual Assistant). Ada yang bisa saya bantu?",
	})

	if len(history) == 0 {
		historyResponse[0].CreatedAt = time.Now().Format("3 January 2006 15:04:05")
	} else {
		historyResponse[0].CreatedAt = history[0].CreatedAt.Format("3 January 2006 15:04:05")
	}

	for _, h := range history {
		historyResponse = append(historyResponse, chatbot_response.GetFromEntitiesToResponse(&h))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get chatbot history", historyResponse))
}

func (cc *ChatbotController) GetChatCompletion(c echo.Context) error {
	userID, _ := utils.GetIDFromJWT(c)

	var request chatbot_request.Chat
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(err.Error()))
	}

	chatbot := request.ToEntities()
	chatbot.UserID = userID

	err := cc.chatbotUseCase.GetChatCompletion(chatbot)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	chatbotResponse := chatbot_response.GetFromEntitiesToResponse(chatbot)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get chat completion", chatbotResponse))
}

func (cc *ChatbotController) ClearHistory(c echo.Context) error {
	userID, _ := utils.GetIDFromJWT(c)

	err := cc.chatbotUseCase.ClearHistory(userID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success clear chatbot history", nil))
}
