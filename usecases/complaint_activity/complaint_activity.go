package complaint_activity

import (
	"e-complaint-api/entities"
)

type ComplaintActivityUseCase struct {
	repo entities.ComplaintActivityRepositoryInterface
}

func NewComplaintActivityUseCase(repo entities.ComplaintActivityRepositoryInterface) *ComplaintActivityUseCase {
	return &ComplaintActivityUseCase{
		repo: repo,
	}
}

func (u *ComplaintActivityUseCase) GetByComplaintIDs(complaintIDs []string, activityType string) ([]entities.ComplaintActivity, error) {
	complaintActivities, err := u.repo.GetByComplaintIDs(complaintIDs, activityType)
	if err != nil {
		return nil, err
	}

	return complaintActivities, nil
}
