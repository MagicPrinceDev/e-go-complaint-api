package user

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/middlewares"
	"e-complaint-api/utils"
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

func (u *UserUseCase) GetAllUsers() ([]*entities.User, error) {
	users, err := u.repository.GetAllUsers()

	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return users, nil
}

func (u *UserUseCase) GetUserByID(id int) (*entities.User, error) {
	user, err := u.repository.GetUserByID(id)

	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return user, nil
}

func (u *UserUseCase) UpdateUser(id int, user *entities.User) (entities.User, error) {
	existingUser, err := u.repository.GetUserByID(id)
	if err != nil {
		return entities.User{}, constants.ErrInternalServerError
	}

	// Ensure existing data remains if no new data is provided
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.TelephoneNumber != "" {
		existingUser.TelephoneNumber = user.TelephoneNumber
	}

	err = u.repository.UpdateUser(id, existingUser)
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

	return *existingUser, nil
}

func (u *UserUseCase) Delete(id int) error {
	err := u.repository.Delete(id)

	if err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (u *UserUseCase) UpdatePassword(id int, oldPassword, newPassword string) error {
	existingUser, err := u.repository.GetUserByID(id)
	if err != nil {
		return constants.ErrInternalServerError
	}

	if !utils.CheckPasswordHash(oldPassword, existingUser.Password) {
		return constants.ErrOldPasswordDoesntMatch
	}

	hash, _ := utils.HashPassword(newPassword)
	return u.repository.UpdatePassword(id, hash)
}
