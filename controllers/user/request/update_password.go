package request

type UpdatePassword struct {
	NewPassword        string `json:"new_password" form:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password" form:"confirm_new_password"`
}

func (up *UpdatePassword) ToEntities() (string, string) {
	return up.ConfirmNewPassword, up.NewPassword
}
