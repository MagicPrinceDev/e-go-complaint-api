package response

import (
	category_response "e-complaint-api/controllers/category/response"
	file_response "e-complaint-api/controllers/complaint_file/response"
	regency_response "e-complaint-api/controllers/regency/response"
	user_response "e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
)

type AdminGet struct {
	ID          string                        `json:"id"`
	User        user_response.GetUser         `json:"user"`
	Category    category_response.Get         `json:"category"`
	Regency     regency_response.Regency      `json:"regency"`
	Address     string                        `json:"address"`
	Description string                        `json:"description"`
	Status      string                        `json:"status"`
	Type        string                        `json:"type"`
	Files       []file_response.ComplaintFile `json:"files"`
	Date        string                        `json:"date"`
	TotalLikes  int                           `json:"total_likes"`
	UpdatedAt   string                        `json:"updated_at"`
}

func AdminGetFromEntitiesToResponse(data *entities.Complaint) *AdminGet {
	var files []file_response.ComplaintFile
	for _, file := range data.Files {
		files = append(files, file_response.ComplaintFile{
			ID:          file.ID,
			ComplaintID: file.ComplaintID,
			Path:        file.Path,
		})
	}

	return &AdminGet{
		ID:          data.ID,
		User:        *user_response.GetUsersFromEntitiesToResponse(&data.User),
		Category:    *category_response.GetFromEntitiesToResponse(&data.Category),
		Regency:     *regency_response.FromEntitiesToResponse(&data.Regency),
		Address:     data.Address,
		Description: data.Description,
		Status:      data.Status,
		Type:        data.Type,
		Files:       files,
		Date:        data.Date.Format("2 January 2006"),
		TotalLikes:  data.TotalLikes,
		UpdatedAt:   data.UpdatedAt.Format("2 January 2006 15:04:05"),
	}
}
