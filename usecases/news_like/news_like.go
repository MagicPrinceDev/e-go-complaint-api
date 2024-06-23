package news_like

import "e-complaint-api/entities"

type NewsLikeUseCase struct {
	repo entities.NewsLikeRepositoryInterface
}

func NewNewsLikeUseCase(repo entities.NewsLikeRepositoryInterface) *NewsLikeUseCase {
	return &NewsLikeUseCase{
		repo: repo,
	}
}

func (u *NewsLikeUseCase) ToggleLike(newsLike *entities.NewsLike) (string, error) {
	like, err := u.repo.FindByUserAndNews(newsLike.UserID, newsLike.NewsID)

	if like == nil {
		err := u.repo.Likes(newsLike)
		if err != nil {
			return "", err
		}

		return "liked", nil
	}

	err = u.repo.Unlike(like)
	if err != nil {
		return "", err
	}

	return "unliked", nil
}

func (u *NewsLikeUseCase) IncreaseTotalLikes(id string) error {
	err := u.repo.IncreaseTotalLikes(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *NewsLikeUseCase) DecreaseTotalLikes(id string) error {
	err := u.repo.DecreaseTotalLikes(id)
	if err != nil {
		return err
	}

	return nil
}
