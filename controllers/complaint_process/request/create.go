package request

import "e-complaint-api/entities"

type Create struct {
	ComplaintID string `json:"complaint_id" form:"complaint_id"`
	AdminID     int    `json:"admin_id" form:"admin_id"`
	Status      string `json:"status" form:"status" validate:"required"`
	Message     string `json:"message" form:"message"`
}

func (r *Create) ToEntities() *entities.ComplaintProcess {
	return &entities.ComplaintProcess{
		ComplaintID: r.ComplaintID,
		AdminID:     r.AdminID,
		Status:      r.Status,
		Message:     r.Message,
	}
}
