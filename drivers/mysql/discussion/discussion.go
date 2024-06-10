package discussion

import (
	"e-complaint-api/entities"
	"gorm.io/gorm"
)

type DiscussionRepo struct {
	DB *gorm.DB
}

func NewDiscussionRepo(db *gorm.DB) *DiscussionRepo {
	return &DiscussionRepo{DB: db}
}

func (r *DiscussionRepo) Create(discussion *entities.Discussion) error {
	if err := r.DB.Create(discussion).Error; err != nil {
		return err
	}

	return nil
}

func (r *DiscussionRepo) GetById(id int) (*entities.Discussion, error) {
	var discussion entities.Discussion

	if err := r.DB.Preload("User").Preload("Admin").Preload("Complaint").First(&discussion, id).Error; err != nil {
		return nil, err
	}

	return &discussion, nil
}

func (r *DiscussionRepo) GetByComplaintID(complaintID string) (*[]entities.Discussion, error) {
	var discussions []entities.Discussion
	if err := r.DB.Preload("User").Preload("Admin").Preload("Complaint").Where("complaint_id = ?", complaintID).Find(&discussions).Error; err != nil {
		return nil, err
	}
	return &discussions, nil
}

func (r *DiscussionRepo) Update(discussion *entities.Discussion) error {
	if err := r.DB.Save(discussion).Error; err != nil {
		return err
	}
	return nil
}

func (r *DiscussionRepo) Delete(id int) error {
	if err := r.DB.Delete(&entities.Discussion{}, id).Error; err != nil {
		return err
	}
	return nil
}
