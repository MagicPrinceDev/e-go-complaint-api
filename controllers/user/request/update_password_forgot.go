package request

type UpdatePasswordForgot struct {
	Email       string `json:"email" form:"email"`
	NewPassword string `json:"new_password" form:"new_password"`
}
