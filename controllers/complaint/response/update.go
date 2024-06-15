package response

import (
	category_response "e-complaint-api/controllers/category/response"
	file_response "e-complaint-api/controllers/complaint_file/response"
	regency_response "e-complaint-api/controllers/regency/response"
	user_response "e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
)

type Update struct {
	ID          string                         `json:"id"`
	User        *user_response.Get             `json:"user"`
	Category    *category_response.Get         `json:"category"`
	Regency     *regency_response.Regency      `json:"regency"`
	Address     string                         `json:"address"`
	Description string                         `json:"description"`
	Status      string                         `json:"status"`
	Type        string                         `json:"type"`
	Date        string                         `json:"date"`
	Files       []*file_response.ComplaintFile `json:"files"`
	UpdatedAt   string                         `json:"updated_at"`
}

func UpdateFromEntitiesToResponse(data *entities.Complaint) *Update {
	if data.Type == "private" {
		(*data).User = entities.User{
			ID:    0,
			Name:  "Anonymous",
			Email: "anonymous@anonymous.com",
		}
	}

	var files []*file_response.ComplaintFile
	for _, file := range data.Files {
		files = append(files, &file_response.ComplaintFile{
			ID:          file.ID,
			ComplaintID: file.ComplaintID,
			Path:        file.Path,
		})
	}

	return &Update{
		ID:          data.ID,
		User:        user_response.GetFromEntitiesToResponse(&data.User),
		Category:    category_response.GetFromEntitiesToResponse(&data.Category),
		Regency:     regency_response.FromEntitiesToResponse(&data.Regency),
		Address:     data.Address,
		Description: data.Description,
		Status:      data.Status,
		Type:        data.Type,
		Date:        data.Date.Format("2 January 2006"),
		Files:       files,
		UpdatedAt:   data.UpdatedAt.Format("2 January 2006 15:04:05"),
	}
}
