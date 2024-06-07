package user

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/middlewares"
	"e-complaint-api/utils"
	"errors"
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
	if user.Email == "" || user.Password == "" || user.Name == "" || user.TelephoneNumber == "" {
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
			return entities.User{}, err
		}
	}

	return *user, nil
}

func (u *UserUseCase) Login(user *entities.User) (entities.User, error) {
	if user.Email == "" || user.Password == "" {
		return entities.User{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Login(user)

	(*user).Token = middlewares.GenerateTokenJWT(user.ID, user.Name, user.Email, "user")

	if err != nil {
		return entities.User{}, err
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
	existingUser, err := u.repository.GetUserByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			return constants.ErrNotFound
		}
		return constants.ErrInternalServerError
	}

	if existingUser == nil {
		return constants.ErrNotFound
	}

	err = u.repository.Delete(id)
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

	if oldPassword == "" || newPassword == "" {
		return constants.ErrAllFieldsMustBeFilled
	}

	if !utils.CheckPasswordHash(oldPassword, existingUser.Password) {
		return constants.ErrOldPasswordDoesntMatch
	}

	hash, _ := utils.HashPassword(newPassword)
	return u.repository.UpdatePassword(id, hash)
}

func (u *UserUseCase) SendOTP(email string) error {
	if email == "" {
		return constants.ErrAllFieldsMustBeFilled
	}

	otp := utils.GenerateOTP(5)

	err := u.repository.SendOTP(email, otp)
	if err != nil {
		return err
	}

	err = u.emailTrapApi.SendOTP(email, otp)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) VerifyOTP(email, otp string) error {
	if email == "" || otp == "" {
		return constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.VerifyOTP(email, otp)
	if err != nil {
		return err
	}

	return nil
}
