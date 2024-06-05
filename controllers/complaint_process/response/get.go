package response

import (
	admin_response "e-complaint-api/controllers/admin/response"
	"e-complaint-api/entities"
)

type Get struct {
	ID          int                 `json:"id"`
	ComplaintID string              `json:"complaint_id"`
	Admin       *admin_response.Get `json:"admin"`
	Status      string              `json:"status"`
	Message     string              `json:"message"`
	UpdatedAt   string              `json:"updated_at"`
}

func GetFromEntitiesToResponse(data *entities.ComplaintProcess) *Get {
	return &Get{
		ID:          data.ID,
		ComplaintID: data.ComplaintID,
		Admin:       admin_response.GetFromEntitiesToResponse(&data.Admin),
		Status:      data.Status,
		Message:     data.Message,
		UpdatedAt:   data.UpdatedAt.Format("3 January 2006 15:04:05"),
	}
}
