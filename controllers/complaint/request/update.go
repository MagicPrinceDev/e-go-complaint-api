package request

import (
	"e-complaint-api/entities"
	"time"
)

type Update struct {
	ID          string
	UserID      int    `json:"user_id" form:"user_id"`
	CategoryID  int    `json:"category_id" form:"category_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	RegencyID   string `json:"regency_id" form:"regency_id" binding:"required"`
	Address     string `json:"address" form:"address" binding:"required"`
	Date        string `json:"date" form:"date"`
	Type        string `json:"type" form:"type" binding:"required"`
}

func (r *Update) ToEntities() *entities.Complaint {
	// Parse from dd-mm-yyyy to yyyy-mm-dd
	date, _ := time.Parse("02-01-2006", r.Date)

	return &entities.Complaint{
		ID:          r.ID,
		UserID:      r.UserID,
		CategoryID:  r.CategoryID,
		Description: r.Description,
		RegencyID:   r.RegencyID,
		Address:     r.Address,
		Date:        date,
		Type:        r.Type,
	}
}
