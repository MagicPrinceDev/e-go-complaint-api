package routes

import (
	"e-complaint-api/controllers/admin"
	"e-complaint-api/controllers/category"
	"e-complaint-api/controllers/complaint"
	"e-complaint-api/controllers/user"
	"e-complaint-api/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AdminController     *admin.AdminController
	UserController      *user.UserController
	ComplaintController *complaint.ComplaintController
	CategoryController  *category.CategoryController
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
  admin.Use(jwt, middlewares.IsAdmin)
	admin.GET("/admins", r.AdminController.GetAllAdmins)
	admin.GET("/admins/:id", r.AdminController.GetAdminByID)
	admin.GET("/users", r.UserController.GetAllUsers)
	admin.POST("/admins/login", r.AdminController.Login)
	admin.PUT("/admins/:id", r.AdminController.UpdateAdmin)
	admin.PUT("/admins/:id/change-password", r.AdminController.UpdatePassword)

	// Route For User
	user := e.Group("/api/v1")
  user.Use(jwt, middlewares.IsUser)
	user.POST("/users/login", r.UserController.Login)
	user.POST("/users/register", r.UserController.Register)
  user.POST("/complaints", r.ComplaintController.Create)
	user.PUT("/complaints/:id", r.ComplaintController.Update)
	user.PUT("/users/:id", r.UserController.UpdateUser)
  user.PUT("/users/:id/change-password", r.UserController.UpdatePassword)

	// Route For All Authenticated User
	auth_user := e.Group("/api/v1")
  auth_user.Use(jwt)
	auth_user.GET("/users/:id", r.UserController.GetUserByID)
	auth_user.DELETE("/users/:id", r.UserController.DeleteUser)
	auth_user.GET("/complaints", r.ComplaintController.GetPaginated)
	auth_user.GET("/complaints/:id", r.ComplaintController.GetByID)
	auth_user.DELETE("/complaints/:id", r.ComplaintController.Delete)
	auth_user.GET("/categories", r.CategoryController.GetAll)


	// Route For Public
}
