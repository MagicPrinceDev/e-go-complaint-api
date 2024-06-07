package complaint_likes

import "e-complaint-api/entities"

type ComplaintLikeUseCase struct {
	complaintLikeRepo entities.ComplaintLikeRepositoryInterface
}

func NewComplaintLikeUseCase(complaintLikeRepo entities.ComplaintLikeRepositoryInterface) *ComplaintLikeUseCase {
	return &ComplaintLikeUseCase{
		complaintLikeRepo: complaintLikeRepo,
	}
}

func (cl *ComplaintLikeUseCase) Likes(complaintLike []*entities.ComplaintLike) error {
	err := cl.complaintLikeRepo.Likes(complaintLike)
	if err != nil {
		return err
	}
	return nil
}

func (cl *ComplaintLikeUseCase) Unlike(complaintLike []*entities.ComplaintLike) error {
	err := cl.complaintLikeRepo.Unlike(complaintLike)
	if err != nil {
		return err
	}
	return nil
}
