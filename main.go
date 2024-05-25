package main

import (
	"e-complaint-api/config"
	"e-complaint-api/drivers/mailtrap"
	"e-complaint-api/drivers/mysql"
	"e-complaint-api/routes"
	"os"

	admin_cl "e-complaint-api/controllers/admin"
	admin_rp "e-complaint-api/drivers/mysql/admin"
	admin_uc "e-complaint-api/usecases/admin"

	complaint_cl "e-complaint-api/controllers/complaint"
	complaint_rp "e-complaint-api/drivers/mysql/complaint"
	complaint_uc "e-complaint-api/usecases/complaint"

	user_cl "e-complaint-api/controllers/user"
	user_rp "e-complaint-api/drivers/mysql/user"
	user_uc "e-complaint-api/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	// For local development only
	config.LoadEnv()

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

	complaintRepo := complaint_rp.NewComplaintRepo(DB)
	complaintUsecase := complaint_uc.NewComplaintUseCase(complaintRepo)
	ComplaintController := complaint_cl.NewComplaintController(complaintUsecase)

	routes := routes.RouteController{
		AdminController:     AdminController,
		UserController:      UserController,
		ComplaintController: ComplaintController,
	}

	routes.InitRoute(e)
	e.Start(":8080")
}
