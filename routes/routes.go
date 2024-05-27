package routes

import (
	"e-complaint-api/controllers/admin"
	"e-complaint-api/controllers/category"
	"e-complaint-api/controllers/complaint"
	"e-complaint-api/controllers/complaint_process"
	"e-complaint-api/controllers/user"
	"e-complaint-api/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AdminController            *admin.AdminController
	UserController             *user.UserController
	ComplaintController        *complaint.ComplaintController
	CategoryController         *category.CategoryController
	ComplaintProcessController *complaint_process.ComplaintProcessController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	var jwt = echojwt.JWT([]byte(os.Getenv("JWT_SECRET")))

	// Route For Super Admin
	superAdmin := e.Group("/api/v1")
	superAdmin.Use(jwt, middlewares.IsSuperAdmin)
	superAdmin.POST("/admins", r.AdminController.CreateAccount)
	superAdmin.GET("/admins", r.AdminController.GetAllAdmins)
	superAdmin.GET("/admins/:id", r.AdminController.GetAdminByID)
	superAdmin.DELETE("/admins/:id", r.AdminController.DeleteAdmin)
	superAdmin.PUT("/admins/:id/change-password", r.AdminController.UpdatePassword)

	// Route For Admin
	admin := e.Group("/api/v1")
	admin.POST("/admins/login", r.AdminController.Login)
	admin.Use(jwt, middlewares.IsAdmin)
	admin.GET("/admins", r.AdminController.GetAllAdmins)
	admin.GET("/admins/:id", r.AdminController.GetAdminByID)
	admin.GET("/users", r.UserController.GetAllUsers)
	admin.PUT("/admins/:id", r.AdminController.UpdateAdmin)
	admin.PUT("/admins/:id/change-password", r.AdminController.UpdatePassword)
	admin.POST("/complaints/:complaint_id/processes", r.ComplaintProcessController.Create)
	admin.GET("/complaints/:complaint_id/processes", r.ComplaintProcessController.GetByComplaintID)

	// Route For User
	user := e.Group("/api/v1")
	user.POST("/users/login", r.UserController.Login)
	user.POST("/users/register", r.UserController.Register)
	user.Use(jwt, middlewares.IsUser)
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
