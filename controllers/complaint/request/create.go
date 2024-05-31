package request

import "e-complaint-api/entities"

type Create struct {
	UserID      int    `json:"user_id" form:"user_id"`
	CategoryID  int    `json:"category_id" form:"category_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	RegencyID   string `json:"regency_id" form:"regency_id" binding:"required"`
	Address     string `json:"address" form:"address" binding:"required"`
	Type        string `json:"type" form:"type" binding:"required"`
}

func (r *Create) ToEntities() *entities.Complaint {
	return &entities.Complaint{
		UserID:      r.UserID,
		CategoryID:  r.CategoryID,
		Description: r.Description,
		RegencyID:   r.RegencyID,
		Address:     r.Address,
		Type:        r.Type,
	}
}
