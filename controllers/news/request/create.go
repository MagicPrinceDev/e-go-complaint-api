package request

import "e-complaint-api/entities"

type Create struct {
	AdminID    int    `json:"admin_id" form:"admin_id"`
	CategoryID int    `json:"category_id" form:"category_id" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	Content    string `json:"content" form:"content" binding:"required"`
}

func (r *Create) ToEntities() *entities.News {
	return &entities.News{
		AdminID:    r.AdminID,
		CategoryID: r.CategoryID,
		Title:      r.Title,
		Content:    r.Content,
	}
}
