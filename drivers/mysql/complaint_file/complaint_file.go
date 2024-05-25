package complaint_file

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type ComplaintFileRepo struct {
	DB *gorm.DB
}

func NewComplaintFileRepo(db *gorm.DB) *ComplaintFileRepo {
	return &ComplaintFileRepo{DB: db}
}

func (r *ComplaintFileRepo) Create(complaintFiles []*entities.ComplaintFile) error {
	if err := r.DB.CreateInBatches(complaintFiles, len(complaintFiles)).Error; err != nil {
		return err
	}

	return nil
}
