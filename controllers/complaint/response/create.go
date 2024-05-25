package response

import (
	regencyResponse "e-complaint-api/controllers/admin/regency/response"
	categoryResponse "e-complaint-api/controllers/category/response"
	fileResponse "e-complaint-api/controllers/complaint_file/response"
	userResponse "e-complaint-api/controllers/user/response"
	"e-complaint-api/entities"
)

type Create struct {
	ID          string                        `json:"id"`
	User        *userResponse.Get             `json:"user"`
	Category    *categoryResponse.Get         `json:"category"`
	Regency     *regencyResponse.Regency      `json:"regency"`
	Address     string                        `json:"address"`
	Description string                        `json:"description"`
	Status      string                        `json:"status"`
	Type        string                        `json:"type"`
	Files       []*fileResponse.ComplaintFile `json:"files"`
}

func CreateFromEntitiesToResponse(data *entities.Complaint) *Create {
	if data.Type == "private" {
		(*data).User = entities.User{
			ID:       0,
			Name:     "Anonymous",
			Username: "Anonymous",
		}
	}

	return &Create{
		ID:          data.ID,
		User:        userResponse.GetFromEntitiesToResponse(&data.User),
		Category:    categoryResponse.GetFromEntitiesToResponse(&data.Category),
		Regency:     regencyResponse.FromEntitiesToResponse(&data.Regency),
		Address:     data.Address,
		Description: data.Description,
		Status:      data.Status,
		Type:        data.Type,
	}
}
