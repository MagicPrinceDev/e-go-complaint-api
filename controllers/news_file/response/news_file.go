package response

import "e-complaint-api/entities"

type NewsFile struct {
	ID     int    `json:"id"`
	NewsID int    `json:"news_id"`
	Path   string `json:"path"`
}

func FromEntitiesToResponse(file *entities.NewsFile) *NewsFile {
	return &NewsFile{
		ID:     file.ID,
		NewsID: file.NewsID,
		Path:   file.Path,
	}
}
