package request

type UpdatePassword struct {
	NewPassword string `json:"new_password" form:"new_password"`
}
