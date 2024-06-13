package main

import (
	"e-complaint-api/config"
	"e-complaint-api/drivers/mailtrap"
	"e-complaint-api/drivers/mysql"
	"e-complaint-api/drivers/openai_api"
	"e-complaint-api/routes"
	"os"

	gcs_api "e-complaint-api/drivers/google_cloud_storage"

	admin_cl "e-complaint-api/controllers/admin"
	admin_rp "e-complaint-api/drivers/mysql/admin"
	admin_uc "e-complaint-api/usecases/admin"

	complaint_cl "e-complaint-api/controllers/complaint"
	complaint_rp "e-complaint-api/drivers/mysql/complaint"
	complaint_uc "e-complaint-api/usecases/complaint"

	complaint_file_rp "e-complaint-api/drivers/mysql/complaint_file"
	complaint_file_uc "e-complaint-api/usecases/complaint_file"

	complaint_process_cl "e-complaint-api/controllers/complaint_process"
	complaint_process_rp "e-complaint-api/drivers/mysql/complaint_process"
	complaint_process_uc "e-complaint-api/usecases/complaint_process"

	user_cl "e-complaint-api/controllers/user"
	user_rp "e-complaint-api/drivers/mysql/user"
	user_uc "e-complaint-api/usecases/user"

	category_cl "e-complaint-api/controllers/category"
	category_rp "e-complaint-api/drivers/mysql/category"
	category_uc "e-complaint-api/usecases/category"

	discussion_cl "e-complaint-api/controllers/discussion"
	discussion_rp "e-complaint-api/drivers/mysql/discussion"
	discussion_uc "e-complaint-api/usecases/discussion"

	regency_cl "e-complaint-api/controllers/regency"
	regency_rp "e-complaint-api/drivers/mysql/regency"
	regency_uc "e-complaint-api/usecases/regency"

	news_cl "e-complaint-api/controllers/news"
	news_rp "e-complaint-api/drivers/mysql/news"
	news_uc "e-complaint-api/usecases/news"

	news_file_rp "e-complaint-api/drivers/mysql/news_file"
	news_file_uc "e-complaint-api/usecases/news_file"

	complaint_like "e-complaint-api/controllers/complaint_like"
	complaint_like_rp "e-complaint-api/drivers/mysql/complaint_like"
	complaint_like_uc "e-complaint-api/usecases/complaint_like"

	complaint_activity "e-complaint-api/controllers/complaint_activity"
	complaint_activity_rp "e-complaint-api/drivers/mysql/complaint_activity"
	complaint_activity_uc "e-complaint-api/usecases/complaint_activity"

	chatbot_cl "e-complaint-api/controllers/chatbot"
	chatbot_rp "e-complaint-api/drivers/mysql/chatbot"
	chatbot_uc "e-complaint-api/usecases/chatbot"

	faq_rp "e-complaint-api/drivers/mysql/faq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// For local development only
	// config.LoadEnv()

	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())

	e := echo.New()
	e.Use(middleware.CORS())

	adminRepo := admin_rp.NewAdminRepo(DB)
	adminUsecase := admin_uc.NewAdminUseCase(adminRepo)
	AdminController := admin_cl.NewAdminController(adminUsecase)

	mailTrapApi := mailtrap.NewMailTrapApi(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_FROM"),
	)
	userRepo := user_rp.NewUserRepo(DB)
	userUsecase := user_uc.NewUserUseCase(userRepo, mailTrapApi)
	UserController := user_cl.NewUserController(userUsecase)

	complaintFileGCSAPI := gcs_api.NewFileHandlingAPI(os.Getenv("GCS_CREDENTIALS"), "complaint-files/")
	complaintFileRepo := complaint_file_rp.NewComplaintFileRepo(DB)
	complaintFileUsecase := complaint_file_uc.NewComplaintFileUseCase(complaintFileRepo, complaintFileGCSAPI)

	complaintRepo := complaint_rp.NewComplaintRepo(DB)
	complaintUsecase := complaint_uc.NewComplaintUseCase(complaintRepo, complaintFileRepo)
	ComplaintController := complaint_cl.NewComplaintController(complaintUsecase, complaintFileUsecase)

	complaintProcessRepo := complaint_process_rp.NewComplaintProcessRepo(DB)
	complaintProcessUsecase := complaint_process_uc.NewComplaintProcessUseCase(complaintProcessRepo, complaintRepo)
	ComplaintProcessController := complaint_process_cl.NewComplaintProcessController(complaintUsecase, complaintProcessUsecase)

	categoryRepo := category_rp.NewCategoryRepo(DB)
	categoryUsecase := category_uc.NewCategoryUseCase(categoryRepo)
	CategoryController := category_cl.NewCategoryController(categoryUsecase)

	regencyRepo := regency_rp.NewRegencyRepo(DB)
	regencyUsecase := regency_uc.NewRegencyUseCase(regencyRepo)
	RegencyController := regency_cl.NewRegencyController(regencyUsecase)

	NewsFileGCSAPIInterface := gcs_api.NewFileHandlingAPI(os.Getenv("GCS_CREDENTIALS"), "news-files/")
	NewsFileRepo := news_file_rp.NewNewsFileRepo(DB)
	NewsFileUsecase := news_file_uc.NewNewsFileUseCase(NewsFileRepo, NewsFileGCSAPIInterface)

	newsRepo := news_rp.NewNewsRepo(DB)
	newsUsecase := news_uc.NewNewsUseCase(newsRepo)
	NewsController := news_cl.NewNewsController(newsUsecase, NewsFileUsecase)

	complaintActivityRepo := complaint_activity_rp.NewComplaintActivityRepo(DB)
	complaintActivityUsecase := complaint_activity_uc.NewComplaintActivityUseCase(complaintActivityRepo)
	ComplaintActivityController := complaint_activity.NewComplaintActivityController(complaintActivityUsecase, complaintUsecase)

	discussionRepo := discussion_rp.NewDiscussionRepo(DB)
	discussionUsecase := discussion_uc.NewDiscussionUseCase(discussionRepo)
	DiscussionController := discussion_cl.NewDiscussionController(discussionUsecase, complaintUsecase, complaintActivityUsecase)

	complaintLikeRepo := complaint_like_rp.NewComplaintLikeRepository(DB)
	complaintLikeUsecase := complaint_like_uc.NewComplaintLikeUseCase(complaintLikeRepo)
	ComplaintLikeController := complaint_like.NewComplaintLikeController(complaintLikeUsecase, complaintUsecase, complaintActivityUsecase)

	openAIAPI := openai_api.NewOpenAIAPI(os.Getenv("OPENAI_API_KEY"))
	faqRepo := faq_rp.NewFaqRepo(DB)
	chatbotRepo := chatbot_rp.NewChatbotRepo(DB)
	chatbotUsecase := chatbot_uc.NewChatbotUseCase(chatbotRepo, faqRepo, complaintRepo, openAIAPI)
	ChatbotController := chatbot_cl.NewChatbotController(chatbotUsecase)

	routes := routes.RouteController{
		AdminController:             AdminController,
		UserController:              UserController,
		ComplaintController:         ComplaintController,
		CategoryController:          CategoryController,
		ComplaintProcessController:  ComplaintProcessController,
		DiscussionController:        DiscussionController,
		NewsController:              NewsController,
		RegencyController:           RegencyController,
		ComplaintLikeController:     ComplaintLikeController,
		ComplaintActivityController: ComplaintActivityController,
		ChatbotController:           ChatbotController,
	}

	routes.InitRoute(e)
	e.Start(":8000")
}
