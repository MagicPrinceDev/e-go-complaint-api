package response

import (
	like_response "e-complaint-api/controllers/complaint_like/response"
	discussion_response "e-complaint-api/controllers/discussion/response"
	"e-complaint-api/entities"
)

type Get struct {
	ID         int                                `json:"id"`
	Discussion *discussion_response.DiscussionGet `json:"discussion,omitempty"`
	Like       *like_response.Get                 `json:"like,omitempty"`
	UpdatedAt  string                             `json:"updated_at"`
}

func GetFromEntitiesToResponse(data *entities.ComplaintActivity) *Get {
	if data.LikeID == nil {
		return &Get{
			ID:         data.ID,
			Discussion: discussion_response.FromEntitiesGetToResponse(&data.Discussion),
			UpdatedAt:  data.UpdatedAt.Format("2 January 2006 15:04:05"),
		}
	} else {
		return &Get{
			ID:        data.ID,
			Like:      like_response.GetFromEntitiesToResponse(&data.Like),
			UpdatedAt: data.UpdatedAt.Format("2 January 2006 15:04:05"),
		}
	}
}
