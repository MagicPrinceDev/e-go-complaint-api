package complaint_process

import (
	"e-complaint-api/entities"

	"gorm.io/gorm"
)

type ComplaintProcessRepo struct {
	DB *gorm.DB
}

func NewComplaintProcessRepo(db *gorm.DB) *ComplaintProcessRepo {
	return &ComplaintProcessRepo{DB: db}
}

func (repo *ComplaintProcessRepo) Create(complaintProcesses *entities.ComplaintProcess) error {
	if err := repo.DB.Create(complaintProcesses).Error; err != nil {
		return err
	}
	return nil
}
