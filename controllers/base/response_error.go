package base

type BaseErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) *BaseErrorResponse {
	return &BaseErrorResponse{
		Status:  false,
		Message: message,
	}
}
