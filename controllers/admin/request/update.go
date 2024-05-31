package request

import "e-complaint-api/entities"

type UpdateAccount struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	TelephoneNumber string `json:"telephone_number" form:"telephone_number" `
	Password        string `json:"password" form:"password"`
}

func (req *UpdateAccount) ToEntities() *entities.Admin {
	return &entities.Admin{
		Name:            req.Name,
		Email:           req.Email,
		TelephoneNumber: req.TelephoneNumber,
		Password:        req.Password,
	}
}
