package discussion

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
)

type DiscussionUseCase struct {
	repository entities.DiscussionRepositoryInterface
}

func NewDiscussionUseCase(repository entities.DiscussionRepositoryInterface) *DiscussionUseCase {
	return &DiscussionUseCase{
		repository: repository,
	}
}

func (u *DiscussionUseCase) Create(discussion *entities.Discussion) error {
	if discussion.Comment == "" {
		return constants.ErrCommentCannotBeEmpty
	}
	err := u.repository.Create(discussion)
	if err != nil {
		return err
	}

	return nil
}

func (u *DiscussionUseCase) GetById(id int) (*entities.Discussion, error) {
	discussion, err := u.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return discussion, nil
}

func (u *DiscussionUseCase) GetByComplaintID(complaintID string) (*[]entities.Discussion, error) {
	discussions, err := u.repository.GetByComplaintID(complaintID)
	if err != nil {
		return nil, err
	}

	return discussions, nil
}

func (u *DiscussionUseCase) Update(discussion *entities.Discussion) error {
	if discussion.Comment == "" {
		return constants.ErrCommentCannotBeEmpty
	}
	err := u.repository.Update(discussion)
	if err != nil {
		return err
	}

	return nil
}

func (u *DiscussionUseCase) Delete(id int) error {
	err := u.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
