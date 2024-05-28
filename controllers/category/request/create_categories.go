package request

import "e-complaint-api/entities"

type CreateCategories struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func (req *CreateCategories) ToEntities() *entities.Category {
	return &entities.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}
