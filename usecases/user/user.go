package user

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/middlewares"
	"e-complaint-api/utils"
	"errors"
	"mime/multipart"
	"strings"
)

type UserUseCase struct {
	repository   entities.UserRepositoryInterface
	emailTrapApi entities.MailTrapAPIInterface
	gcsAPI       entities.UserGCSAPIInterface
}

func NewUserUseCase(repository entities.UserRepositoryInterface, emailTrapApi entities.MailTrapAPIInterface, gcsAPI entities.UserGCSAPIInterface) *UserUseCase {
	return &UserUseCase{
		repository:   repository,
		emailTrapApi: emailTrapApi,
		gcsAPI:       gcsAPI,
	}
}

func (u *UserUseCase) Register(user *entities.User) (entities.User, error) {
	if user.Email == "" || user.Password == "" || user.Name == "" || user.TelephoneNumber == "" {
		return entities.User{}, constants.ErrAllFieldsMustBeFilled
	}

	if len(user.Password) < 8 {
		return entities.User{}, constants.ErrPasswordMustBeAtLeast8Characters
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
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) UpdateUser(id int, user *entities.User) (entities.User, error) {
	if user.Email == "" || user.Name == "" || user.TelephoneNumber == "" {
		return entities.User{}, constants.ErrAllFieldsMustBeFilled
	}

	existingUser, err := u.repository.GetUserByID(id)
	if err != nil {
		return entities.User{}, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.TelephoneNumber = user.TelephoneNumber

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

func (u *UserUseCase) UpdateProfilePhoto(id int, profilePhoto *multipart.FileHeader) error {
	filepaths, err := u.gcsAPI.Upload([]*multipart.FileHeader{profilePhoto})
	if err != nil {
		return err
	}

	err = u.repository.UpdateProfilePhoto(id, filepaths[0])
	if err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (u *UserUseCase) Delete(id int) error {
	_, err := u.repository.GetUserByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return constants.ErrUserNotFound
		}
		return constants.ErrInternalServerError
	}

	err = u.repository.Delete(id)
	if err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (u *UserUseCase) UpdatePassword(id int, newPassword string) error {
	if newPassword == "" {
		return constants.ErrAllFieldsMustBeFilled
	}

	if len(newPassword) < 8 {
		return constants.ErrPasswordMustBeAtLeast8Characters
	}

	hash, _ := utils.HashPassword(newPassword)
	return u.repository.UpdatePassword(id, hash)
}

func (u *UserUseCase) SendOTP(email, otp_type string) error {
	if email == "" {
		return constants.ErrAllFieldsMustBeFilled
	}

	otp := utils.GenerateOTP(5)

	err := u.repository.SendOTP(email, otp)
	if err != nil {
		return err
	}

	err = u.emailTrapApi.SendOTP(email, otp, otp_type)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) VerifyOTP(email, otp, otp_type string) error {
	if email == "" || otp == "" {
		return constants.ErrAllFieldsMustBeFilled
	}

	if otp_type == "forgot_password" {
		err := u.repository.VerifyOTPForgotPassword(email, otp)
		if err != nil {
			return err
		}
	} else {
		err := u.repository.VerifyOTPRegister(email, otp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *UserUseCase) UpdatePasswordForgot(email, newPassword string) error {
	if email == "" || newPassword == "" {
		return constants.ErrAllFieldsMustBeFilled
	}

	if len(newPassword) < 8 {
		return constants.ErrPasswordMustBeAtLeast8Characters
	}

	hash, _ := utils.HashPassword(newPassword)
	err := u.repository.UpdatePasswordForgot(email, hash)
	if err != nil {
		return err
	}

	return nil
}
