package routes

import (
	"e-complaint-api/controllers/admin"
	"e-complaint-api/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AdminController *admin.AdminController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	var jwt = echojwt.JWT([]byte(os.Getenv("JWT_SECRET")))

	// Route For Super Admin
	superAdmin := e.Group("/api/v1")
	superAdmin.Use(jwt, middlewares.IsSuperAdmin)
	superAdmin.POST("/admins", r.AdminController.CreateAccount)
	superAdmin.POST("/admins/login", r.AdminController.Login)
	superAdmin.GET("/admins", r.AdminController.GetAllAdmins)
	superAdmin.GET("/admins/:id", r.AdminController.GetAdminByID)
	superAdmin.DELETE("/admins/:id", r.AdminController.DeleteAdmin)
	superAdmin.PUT("/admins/:id", r.AdminController.UpdateAdmin)

	// Route For Admin
	admin := e.Group("/api/v1")
	admin.GET("/admins", r.AdminController.GetAllAdmins)
	admin.GET("/admins/:id", r.AdminController.GetAdminByID)
	admin.POST("/admins/login", r.AdminController.Login)

	// Route For User

	// Route For Admin and User

	// Route For Public
}
