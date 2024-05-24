package main

import (
	"e-complaint-api/config"
	admin_cl "e-complaint-api/controllers/admin"
	"e-complaint-api/drivers/mysql"
	admin_rp "e-complaint-api/drivers/mysql/admin"
	"e-complaint-api/routes"
	admin_uc "e-complaint-api/usecases/admin"

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

	routes := routes.RouteController{
		AdminController: AdminController,
	}

	routes.InitRoute(e)
	e.Start(":8000")
}
