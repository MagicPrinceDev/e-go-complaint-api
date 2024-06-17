package request

type VerifyOTP struct {
	Email string `json:"email" form:"email"`
	OTP   string `json:"otp" form:"otp"`
}
