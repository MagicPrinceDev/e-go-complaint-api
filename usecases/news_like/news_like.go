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

func (u *NewsLikeUseCase) LikeNews(userID int, newsID int) error {
	newsLike := entities.NewsLike{
		UserID: userID,
		NewsID: newsID,
	}

	err := u.repo.Likes(&newsLike)
	if err != nil {
		return err
	}

	return nil
}

func (u *NewsLikeUseCase) ToggleLike(newsLike *entities.NewsLike) (string, error) {
	like, err := u.repo.FindByUserAndNews(newsLike.UserID, newsLike.NewsID)
	if err != nil {
		return "", err
	}

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

func (u *NewsLikeUseCase) UnlikeNews(userID int, newsID int) error {
	newsLike, err := u.repo.FindByUserAndNews(userID, newsID)
	if err != nil {
		return err
	}

	if newsLike == nil {
		return nil
	}

	err = u.repo.Unlike(newsLike)
	if err != nil {
		return err
	}

	return nil
}

func (u *NewsLikeUseCase) FindByUserAndNews(userID int, newsID int) (*entities.NewsLike, error) {
	newsLike, err := u.repo.FindByUserAndNews(userID, newsID)
	if err != nil {
		return nil, err
	}

	return newsLike, nil
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
