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

	if err := repo.DB.Preload("Admin").First(complaintProcesses).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ComplaintProcessRepo) GetByComplaintID(complaintID string) ([]entities.ComplaintProcess, error) {
	var complaintProcesses []entities.ComplaintProcess
	if err := repo.DB.Where("complaint_id = ?", complaintID).Preload("Admin").Find(&complaintProcesses).Error; err != nil {
		return nil, err
	}

	return complaintProcesses, nil
}

func (repo *ComplaintProcessRepo) Update(complaintProcesses *entities.ComplaintProcess) error {
	var oldComplaintProcess entities.ComplaintProcess
	if err := repo.DB.First(&oldComplaintProcess, complaintProcesses.ID).Error; err != nil {
		return err
	}

	oldComplaintProcess.Message = complaintProcesses.Message
	oldComplaintProcess.AdminID = complaintProcesses.AdminID
	if err := repo.DB.Save(&oldComplaintProcess).Error; err != nil {
		return err
	}

	if err := repo.DB.Preload("Admin").First(complaintProcesses).Error; err != nil {
		return err
	}

	return nil
}
