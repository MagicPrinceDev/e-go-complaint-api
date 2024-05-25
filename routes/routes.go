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

	// Route For Admin
	admin := e.Group("/api/v1")
	admin.POST("/admins/login", r.AdminController.Login)

	// Route For User
	user := e.Group("/api/v1")
	user.POST("/users/login", r.UserController.Login)
	user.POST("/users/register", r.UserController.Register)
	user.POST("/complaints", r.ComplaintController.Create)
	user.PUT("/complaints/:id", r.ComplaintController.Update)

	// Route For All Authenticated User
	all_user := e.Group("/api/v1")
	all_user.Use(jwt)
	all_user.GET("/complaints", r.ComplaintController.GetPaginated)
	all_user.GET("/complaints/:id", r.ComplaintController.GetByID)
	all_user.DELETE("/complaints/:id", r.ComplaintController.Delete)
	all_user.GET("/categories", r.CategoryController.GetAll)

	// Route For Public
}
