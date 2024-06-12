package routes

import (
	"e-complaint-api/controllers/admin"
	"e-complaint-api/controllers/category"
	"e-complaint-api/controllers/complaint"
	complaint_like "e-complaint-api/controllers/complaint_likes"
	"e-complaint-api/controllers/complaint_process"
	"e-complaint-api/controllers/discussion"
	"e-complaint-api/controllers/news"
	"e-complaint-api/controllers/news_comment"
	"e-complaint-api/controllers/news_like"
	"e-complaint-api/controllers/regency"
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
	DiscussionController       *discussion.DiscussionController
	NewsController             *news.NewsController
	RegencyController          *regency.RegencyController
	ComplaintLikeController    *complaint_like.ComplaintLikeController
	NewsLikeController         *news_like.NewsLikeController
	NewsCommentController      *news_comment.NewsCommentController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	var jwt = echojwt.JWT([]byte(os.Getenv("JWT_SECRET")))

	// Route For Super Admin
	superAdmin := e.Group("/api/v1")
	superAdmin.Use(jwt, middlewares.IsSuperAdmin)
	superAdmin.POST("/admins", r.AdminController.CreateAccount)
	superAdmin.DELETE("/admins/:id", r.AdminController.DeleteAdmin)
	superAdmin.PUT("/admins/:id", r.AdminController.UpdateAdmin)

	// Route For Admin & Super Admin
	admin := e.Group("/api/v1")
	admin.POST("/admins/login", r.AdminController.Login)
	admin.Use(jwt, middlewares.IsAdmin)
	admin.GET("/admins", r.AdminController.GetAllAdmins)
	admin.GET("/admins/:id", r.AdminController.GetAdminByID)
	admin.GET("/users", r.UserController.GetAllUsers)
	admin.POST("/complaints/:complaint-id/processes", r.ComplaintProcessController.Create)
	admin.PUT("/complaints/:complaint-id/processes/:process-id", r.ComplaintProcessController.Update)
	admin.POST("/categories", r.CategoryController.CreateCategory)
	admin.PUT("/categories/:id", r.CategoryController.UpdateCategory)
	admin.DELETE("/categories/:id", r.CategoryController.DeleteCategory)
	admin.DELETE("/complaints/:complaint-id/processes/:process-id", r.ComplaintProcessController.Delete)
	admin.POST("/news", r.NewsController.Create)
	admin.DELETE("/news/:id", r.NewsController.Delete)
	admin.PUT("/news/:id", r.NewsController.Update)
	admin.POST("/complaints/import", r.ComplaintController.Import)

	// Route For User
	user := e.Group("/api/v1")
	user.POST("/users/login", r.UserController.Login)
	user.POST("/users/register", r.UserController.Register)
	user.POST("/users/send-otp", r.UserController.SendOTP)
	user.POST("/users/verify-otp", r.UserController.VerifyOTP)
	user.Use(jwt, middlewares.IsUser)
	user.POST("/complaints", r.ComplaintController.Create)
	user.PUT("/complaints/:id", r.ComplaintController.Update)
	user.PUT("/users/update-profile", r.UserController.UpdateUser)
	user.PUT("/users/change-password", r.UserController.UpdatePassword)
	user.GET("/users/complaints", r.ComplaintController.GetByUserID)
	user.POST("/complaints/:complaint-id/likes", r.ComplaintLikeController.ToggleLike)

	// Route For All Authenticated User
	auth_user := e.Group("/api/v1")
	auth_user.Use(jwt)
	auth_user.GET("/users/:id", r.UserController.GetUserByID)
	auth_user.DELETE("/users/:id", r.UserController.DeleteUser)
	auth_user.GET("/complaints", r.ComplaintController.GetPaginated)
	auth_user.GET("/complaints/:id", r.ComplaintController.GetByID)
	auth_user.DELETE("/complaints/:id", r.ComplaintController.Delete)
	auth_user.GET("/complaints/:complaint-id/processes", r.ComplaintProcessController.GetByComplaintID)
	auth_user.GET("/categories", r.CategoryController.GetAll)
	auth_user.GET("/categories/:id", r.CategoryController.GetByID)
	auth_user.DELETE("/complaints/:complaint-id/discussions/:discussion-id", r.DiscussionController.DeleteDiscussion)
	auth_user.GET("/complaints/:complaint-id/discussions", r.DiscussionController.GetDiscussionByComplaintID)
	auth_user.POST("/complaints/:complaint-id/discussions", r.DiscussionController.CreateDiscussion)
	auth_user.PUT("/complaints/:complaint-id/discussions/:discussion-id", r.DiscussionController.UpdateDiscussion)
	auth_user.GET("/news", r.NewsController.GetPaginated)
	auth_user.GET("/news/:id", r.NewsController.GetByID)
	auth_user.GET("/regencies", r.RegencyController.GetAll)
	auth_user.POST("/news/:news-id/likes", r.NewsLikeController.ToggleLike)
	auth_user.POST("/news/:news-id/comments", r.NewsCommentController.CommentNews)
	auth_user.GET("/news/:news-id/comments", r.NewsCommentController.GetCommentNews)
	auth_user.PUT("/news/:news-id/comments/:comment-id", r.NewsCommentController.UpdateComment)
	auth_user.DELETE("/news/:news-id/comments/:comment-id", r.NewsCommentController.DeleteComment)

	// Route For Public
}
