package complaint_process

import (
	"e-complaint-api/constants"
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
		return nil, constants.ErrInternalServerError
	}

	if len(complaintProcesses) == 0 {
		return nil, constants.ErrComplaintProcessNotFound
	}

	return complaintProcesses, nil
}

func (repo *ComplaintProcessRepo) Update(complaintProcesses *entities.ComplaintProcess) error {
	var oldComplaintProcess entities.ComplaintProcess
	if err := repo.DB.Where("complaint_id = ? AND id = ?", complaintProcesses.ComplaintID, complaintProcesses.ID).First(&oldComplaintProcess).Error; err != nil {
		return constants.ErrComplaintProcessNotFound
	}

	oldComplaintProcess.Message = complaintProcesses.Message
	oldComplaintProcess.AdminID = complaintProcesses.AdminID
	if err := repo.DB.Save(&oldComplaintProcess).Error; err != nil {
		return constants.ErrInternalServerError
	}

	if err := repo.DB.Preload("Admin").First(complaintProcesses).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (repo *ComplaintProcessRepo) Delete(complaintID string, complaintProcessID int) (string, error) {
	var complaintProcess entities.ComplaintProcess
	if err := repo.DB.Where("complaint_id = ?", complaintID).Order("created_at desc").First(&complaintProcess).Error; err != nil {
		return "", constants.ErrComplaintProcessNotFound
	}

	if complaintProcess.ID != complaintProcessID {
		return "", constants.ErrComplaintProcessCannotBeDeleted
	}

	complaintProcess.DeletedAt = gorm.DeletedAt{Time: complaintProcess.CreatedAt, Valid: true}
	if err := repo.DB.Save(&complaintProcess).Error; err != nil {
		return "", constants.ErrInternalServerError
	}

	return complaintProcess.Status, nil
}
