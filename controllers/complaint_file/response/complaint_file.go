package response

import "e-complaint-api/entities"

type ComplaintFile struct {
	ID          int    `json:"id"`
	ComplaintID string `json:"complaint_id"`
	Path        string `json:"path"`
}

func FromEntitiesToResponse(file *entities.ComplaintFile) *ComplaintFile {
	return &ComplaintFile{
		ID:          file.ID,
		ComplaintID: file.ComplaintID,
		Path:        file.Path,
	}
}
