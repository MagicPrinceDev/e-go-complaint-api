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

func (u *ComplaintLikeUseCase) ToggleLike(complaintLike *entities.ComplaintLike) (string, error) {
	existingComplaintLike, err := u.repo.FindByUserAndComplaint(complaintLike.UserID, complaintLike.ComplaintID)
	if err != nil {
		return "", err
	}

	var likeStatus string
	if existingComplaintLike == nil {
		err = u.repo.Likes(complaintLike)
		likeStatus = "liked"
	} else {
		err = u.repo.Unlike(existingComplaintLike)
		likeStatus = "unliked"
		*complaintLike = *existingComplaintLike
	}

	if err != nil {
		return "", err
	}

	return likeStatus, nil
}
