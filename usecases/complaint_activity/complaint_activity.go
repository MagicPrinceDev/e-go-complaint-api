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

func (u *ComplaintActivityUseCase) Create(complaintActivity *entities.ComplaintActivity) (entities.ComplaintActivity, error) {
	err := u.repo.Create(complaintActivity)
	if err != nil {
		return entities.ComplaintActivity{}, err
	}

	return *complaintActivity, nil
}

func (u *ComplaintActivityUseCase) Delete(complaintActivity entities.ComplaintActivity) error {
	err := u.repo.Delete(complaintActivity)
	if err != nil {
		return err
	}

	return nil
}

func (u *ComplaintActivityUseCase) Update(complaintActivity entities.ComplaintActivity) error {
	err := u.repo.Update(complaintActivity)
	if err != nil {
		return err
	}

	return nil
}
