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

func (u *ComplaintLikeUseCase) ToggleLike(complaintLike *entities.ComplaintLike) error {
	existingComplaintLike, err := u.repo.FindByUserAndComplaint(complaintLike.UserID, complaintLike.ComplaintID)
	if err != nil {
		return err
	}

	if existingComplaintLike == nil {
		complaintLike.LikeStatus = true
		err = u.repo.Likes(complaintLike)
	} else {
		existingComplaintLike.LikeStatus = !existingComplaintLike.LikeStatus
		if existingComplaintLike.LikeStatus {
			err = u.repo.Likes(existingComplaintLike)
		} else {
			err = u.repo.Unlike(existingComplaintLike)
		}
		complaintLike.LikeStatus = existingComplaintLike.LikeStatus
	}

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
