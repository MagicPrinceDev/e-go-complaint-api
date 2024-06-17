package response

import (
	admin_response "e-complaint-api/controllers/admin/response"
	category_response "e-complaint-api/controllers/category/response"
	file_response "e-complaint-api/controllers/news_file/response"
	"e-complaint-api/entities"
)

type Create struct {
	ID         int                       `json:"id"`
	Admin      admin_response.Get        `json:"admin"`
	Category   category_response.Get     `json:"category"`
	Title      string                    `json:"title"`
	Content    string                    `json:"content"`
	TotalLikes int                       `json:"total_likes"`
	Files      []*file_response.NewsFile `json:"files"`
	CreatedAt  string                    `json:"created_at"`
}

func CreateFromEntitiesToResponse(data *entities.News) *Create {
	return &Create{
		ID:         data.ID,
		Admin:      *admin_response.GetFromEntitiesToResponse(&data.Admin),
		Category:   *category_response.GetFromEntitiesToResponse(&data.Category),
		Title:      data.Title,
		Content:    data.Content,
		TotalLikes: data.TotalLikes,
		CreatedAt:  data.CreatedAt.Format("2 January 2006 15:04:05"),
	}
}
