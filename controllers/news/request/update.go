package request

import "e-complaint-api/entities"

type Update struct {
	ID         int    `json:"id" form:"id"`
	AdminID    int    `json:"admin_id" form:"admin_id"`
	CategoryID int    `json:"category_id" form:"category_id" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	Content    string `json:"content" form:"content" binding:"required"`
}

func (r *Update) ToEntities() *entities.News {
	return &entities.News{
		ID:         r.ID,
		AdminID:    r.AdminID,
		CategoryID: r.CategoryID,
		Title:      r.Title,
		Content:    r.Content,
	}
}
