package user

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/middlewares"
	"strings"
)

type UserUseCase struct {
	repository   entities.UserRepositoryInterface
	emailTrapApi entities.MailTrapAPIInterface
}

func NewUserUseCase(repository entities.UserRepositoryInterface, emailTrapApi entities.MailTrapAPIInterface) *UserUseCase {
	return &UserUseCase{
		repository:   repository,
		emailTrapApi: emailTrapApi,
	}
}

func (u *UserUseCase) Register(user *entities.User) (entities.User, error) {
	if user.Email == "" || user.Password == "" || user.Username == "" || user.Name == "" || user.TelephoneNumber == "" {
		return entities.User{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Register(user)

	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			if strings.HasSuffix(err.Error(), "email'") {
				return entities.User{}, constants.ErrEmailAlreadyExists
			} else if strings.HasSuffix(err.Error(), "username'") {
				return entities.User{}, constants.ErrUsernameAlreadyExists
			}
		} else {
			return entities.User{}, constants.ErrInternalServerError
		}
	}

	err2 := u.emailTrapApi.SendEmail(user.Email)
	if err2 != nil {
		return entities.User{}, constants.ErrInternalServerError
	}

	return *user, nil
}

func (u *UserUseCase) Login(user *entities.User) (entities.User, error) {
	if user.Username == "" || user.Password == "" {
		return entities.User{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Login(user)

	(*user).Token = middlewares.GenerateTokenJWT(user.ID, user.Username, "user")

	if err != nil {
		return entities.User{}, constants.ErrInvalidUsernameOrPassword
	}

	return *user, nil
}
