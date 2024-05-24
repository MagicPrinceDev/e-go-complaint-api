package request

import "e-complaint-api/entities"

type UpdateAccount struct {
	Name            string `json:"name" form:"name"`
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	TelephoneNumber string `json:"telephone_number" form:"telephone_number" `
}

func (req *UpdateAccount) ToEntities() *entities.Admin {
	return &entities.Admin{
		Name:            req.Name,
		Username:        req.Username,
		Email:           req.Email,
		TelephoneNumber: req.TelephoneNumber,
	}
}
