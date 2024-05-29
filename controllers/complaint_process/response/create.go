package response

import (
	admin_response "e-complaint-api/controllers/admin/response"
	"e-complaint-api/entities"
)

type Create struct {
	ID          int                 `json:"id"`
	ComplaintID string              `json:"complaint_id"`
	Admin       *admin_response.Get `json:"admin"`
	Status      string              `json:"status"`
	Message     string              `json:"message"`
}

func CreateFromEntitiesToResponse(data *entities.ComplaintProcess) *Create {
	return &Create{
		ID:          data.ID,
		ComplaintID: data.ComplaintID,
		Admin:       admin_response.GetFromEntitiesToResponse(&data.Admin),
		Status:      data.Status,
		Message:     data.Message,
	}
}