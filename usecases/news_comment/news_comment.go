package news_comment

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
)

type NewsCommentUseCase struct {
	repo entities.NewsCommentRepositoryInterface
}

func NewNewsCommentUseCase(repo entities.NewsCommentRepositoryInterface) *NewsCommentUseCase {
	return &NewsCommentUseCase{
		repo: repo,
	}
}

func (ncu *NewsCommentUseCase) CommentNews(newsComment *entities.NewsComment) error {
	err := ncu.repo.CommentNews(newsComment)

	if newsComment.Comment == "" {
		return constants.ErrCommentCannotBeEmpty
	}

	if err != nil {
		return err
	}

	return nil
}

func (ncu *NewsCommentUseCase) GetById(id int) (*entities.NewsComment, error) {
	newsComment, err := ncu.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return newsComment, nil
}

func (ncu *NewsCommentUseCase) GetByNewsId(newsId int) ([]entities.NewsComment, error) {
	newsComment, err := ncu.repo.GetByNewsId(newsId)
	if err != nil {
		return nil, err
	}

	return newsComment, nil
}

func (ncu *NewsCommentUseCase) UpdateComment(newsComment *entities.NewsComment) error {
	err := ncu.repo.UpdateComment(newsComment)
	if err != nil {
		return err
	}

	return nil
}

func (ncu *NewsCommentUseCase) DeleteComment(id int) error {
	err := ncu.repo.DeleteComment(id)
	if err != nil {
		return err
	}

	return nil
}
