package base

type BaseDeletedResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func NewDeletedResponse(message string) *BaseDeletedResponse {
	return &BaseDeletedResponse{
		Status:  true,
		Message: message,
	}
}
