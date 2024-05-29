package request

import "e-complaint-api/entities"

type Update struct {
	ID          int    `json:"id" form:"id"`
	ComplaintID string `json:"complaint_id" form:"complaint_id"`
	AdminID     int    `json:"admin_id" form:"admin_id"`
	Message     string `json:"message" form:"message"`
}

func (r *Update) ToEntities() *entities.ComplaintProcess {
	return &entities.ComplaintProcess{
		ID:          r.ID,
		ComplaintID: r.ComplaintID,
		AdminID:     r.AdminID,
		Message:     r.Message,
	}
}
