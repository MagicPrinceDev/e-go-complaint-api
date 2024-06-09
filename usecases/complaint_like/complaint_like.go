package complaint_like

import (
	"e-complaint-api/entities"
)

type ComplaintLikeUseCase struct {
	repo entities.ComplaintLikeRepositoryInterface
}

func NewComplaintLikeUseCase(repo entities.ComplaintLikeRepositoryInterface) *ComplaintLikeUseCase {
	return &ComplaintLikeUseCase{
		repo: repo,
	}
}

func (clu *ComplaintLikeUseCase) ToggleLike(complaintLike *entities.ComplaintLike) error {
	existingComplaintLike, err := clu.repo.FindByUserAndComplaint(complaintLike.UserID, complaintLike.ComplaintID)
	if err != nil {
		if err.Error() == "record not found" {
			err = clu.repo.Likes(complaintLike)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	err = clu.repo.Unlike(existingComplaintLike)
	if err != nil {
		return err
	}

	return nil
}

func (clu *ComplaintLikeUseCase) Likes(complaintLike *entities.ComplaintLike) error {
	return clu.repo.Likes(complaintLike)
}

func (clu *ComplaintLikeUseCase) Unlike(complaintLike *entities.ComplaintLike) error {
	return clu.repo.Unlike(complaintLike)
}

func (clu *ComplaintLikeUseCase) FindByUserAndComplaint(userID int, complaintID string) (*entities.ComplaintLike, error) {
	return clu.repo.FindByUserAndComplaint(userID, complaintID)
}
