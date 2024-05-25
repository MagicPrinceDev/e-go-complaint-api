package routes

import (
	"e-complaint-api/controllers/admin"
	"e-complaint-api/controllers/user"
	"e-complaint-api/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AdminController *admin.AdminController
	UserController  *user.UserController
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
	superAdmin.PUT("/admins/:id/change-password", r.AdminController.UpdatePassword)

	// Route For Admin
	admin := e.Group("/api/v1")
	admin.GET("/admins", r.AdminController.GetAllAdmins)
	admin.GET("/admins/:id", r.AdminController.GetAdminByID)
	admin.GET("/users", r.UserController.GetAllUsers)
	admin.POST("/admins/login", r.AdminController.Login)
	admin.PUT("/admins/:id", r.AdminController.UpdateAdmin)
	admin.PUT("/admins/:id/change-password", r.AdminController.UpdatePassword)

	// Route For User
	user := e.Group("/api/v1")
	user.POST("/users/login", r.UserController.Login)
	user.POST("/users/register", r.UserController.Register)
	user.Use(jwt, middlewares.IsUser)
	user.PUT("/users/:id", r.UserController.UpdateUser)

	// Route For Admin and User
	au := e.Group("/api/v1")
	au.GET("/users/:id", r.UserController.GetUserByID)
	au.PUT("/user/:id/change-password", r.UserController.UpdatePassword)
	au.DELETE("/users/:id", r.UserController.DeleteUser)

	// Route For Public
}
