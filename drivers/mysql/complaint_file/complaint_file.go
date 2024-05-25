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

func (r *ComplaintFileRepo) DeleteByComplaintID(complaintID string) error {
	var complaintFiles []entities.ComplaintFile
	if err := r.DB.Where("complaint_id = ?", complaintID).Find(&complaintFiles).Error; err != nil {
		return err
	}

	for _, cf := range complaintFiles {
		cf.DeletedAt = gorm.DeletedAt{Time: cf.UpdatedAt, Valid: true}
		if err := r.DB.Save(&cf).Error; err != nil {
			return err
		}
	}

	return nil
}
