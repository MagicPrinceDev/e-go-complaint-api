package main

import (
	"e-complaint-api/config"
	"e-complaint-api/drivers/mailtrap"
	"e-complaint-api/drivers/mysql"
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

	user_cl "e-complaint-api/controllers/user"
	user_rp "e-complaint-api/drivers/mysql/user"
	user_uc "e-complaint-api/usecases/user"

	category_cl "e-complaint-api/controllers/category"
	category_rp "e-complaint-api/drivers/mysql/category"
	category_uc "e-complaint-api/usecases/category"

	"github.com/labstack/echo/v4"
)

func main() {
	// For local development only
	// config.LoadEnv()

	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())

	e := echo.New()

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

	complaintFileGCSAPI := gcs_api.NewFileHandlingAPI(os.Getenv("GCS_CREDENTIALS"), "complaint_files/")
	complaintFileRepo := complaint_file_rp.NewComplaintFileRepo(DB)
	complaintFileUsecase := complaint_file_uc.NewComplaintFileUseCase(complaintFileRepo, complaintFileGCSAPI)

	complaintRepo := complaint_rp.NewComplaintRepo(DB)
	complaintUsecase := complaint_uc.NewComplaintUseCase(complaintRepo)
	ComplaintController := complaint_cl.NewComplaintController(complaintUsecase, complaintFileUsecase)

	categoryRepo := category_rp.NewCategoryRepo(DB)
	categoryUsecase := category_uc.NewCategoryUseCase(categoryRepo)
	CategoryController := category_cl.NewCategoryController(categoryUsecase)

	routes := routes.RouteController{
		AdminController:     AdminController,
		UserController:      UserController,
		ComplaintController: ComplaintController,
		CategoryController:  CategoryController,
	}

	routes.InitRoute(e)
	e.Start(":8000")
}
