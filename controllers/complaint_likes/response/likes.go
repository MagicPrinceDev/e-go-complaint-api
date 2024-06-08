package response

type LikeResponse struct {
	Message string `json:"message"`
}

func NewLikeResponse(likeStatus bool) *LikeResponse {
	message := "Success Unlike"
	if likeStatus {
		message = "Success Like"
	}
	return &LikeResponse{
		Message: message,
	}
}
