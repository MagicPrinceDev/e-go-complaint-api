package response

import "e-complaint-api/entities"

type Regency struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FromEntitiesToResponse(regeny *entities.Regency) *Regency {
	return &Regency{
		ID:   regeny.ID,
		Name: regeny.Name,
	}
}
