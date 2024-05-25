package request

type UpdatePassword struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
}

func (up *UpdatePassword) ToEntities() (string, string) {
	return up.OldPassword, up.NewPassword
}
