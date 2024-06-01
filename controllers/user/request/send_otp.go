package request

type SendOTP struct {
	Email string `json:"email" form:"email"`
}
