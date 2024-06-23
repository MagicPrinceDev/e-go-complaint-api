package user

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Register(user *entities.User) error {
	hash, _ := utils.HashPassword(user.Password)
	(*user).Password = hash

	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Login(user *entities.User) error {
	var userDB entities.User

	if err := r.DB.Where("email = ?", user.Email).First(&userDB).Error; err != nil {
		return constants.ErrInvalidUsernameOrPassword
	}

	if !utils.CheckPasswordHash(user.Password, userDB.Password) {
		return constants.ErrInvalidUsernameOrPassword
	}

	if !userDB.EmailVerified {
		return constants.ErrEmailNotVerified
	}

	(*user).ID = userDB.ID
	(*user).Name = userDB.Name
	(*user).Email = userDB.Email

	return nil
}

func (r *UserRepo) GetAllUsers() ([]*entities.User, error) {
	var users []*entities.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) GetUserByID(id int) (*entities.User, error) {
	var user entities.User

	if err := r.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) UpdateUser(id int, user *entities.User) error {
	if err := r.DB.Model(&entities.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) UpdateProfilePhoto(id int, profilePhoto string) error {
	if err := r.DB.Model(&entities.User{}).Where("id = ?", id).Update("profile_photo", profilePhoto).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Delete(id int) error {
	if err := r.DB.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) UpdatePassword(id int, newPassword string) error {
	return r.DB.Model(&entities.User{}).Where("id = ?", id).Updates(&entities.User{Password: newPassword}).Error
}

func (r *UserRepo) SendOTP(email, otp string) error {
	otpExpiredAt := time.Now().Add(time.Minute * 10)

	var user entities.User
	if err := r.DB.Model(&entities.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return constants.ErrEmailNotRegistered
	}

	user.Otp = otp
	user.OtpExpiredAt = otpExpiredAt

	if err := r.DB.Save(&user).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *UserRepo) VerifyOTPRegister(email, otp string) error {
	var user entities.User
	if err := r.DB.Model(&entities.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return constants.ErrEmailNotRegistered
	}

	if user.Otp != otp {
		return constants.ErrInvalidOTP
	}

	if user.OtpExpiredAt.Before(time.Now()) {
		return constants.ErrExpiredOTP
	}

	user.EmailVerified = true

	if err := r.DB.Save(&user).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *UserRepo) VerifyOTPForgotPassword(email, otp string) error {
	var user entities.User
	if err := r.DB.Model(&entities.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return constants.ErrUserNotFound
	}

	if user.Otp != otp {
		return constants.ErrInvalidOTP
	}

	if user.OtpExpiredAt.Before(time.Now()) {
		return constants.ErrExpiredOTP
	}

	user.ForgotVerified = true
	if err := r.DB.Save(&user).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *UserRepo) UpdatePasswordForgot(email, newPassword string) error {
	var user entities.User
	if err := r.DB.Model(&entities.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return constants.ErrUserNotFound
	}

	if !user.ForgotVerified {
		return constants.ErrForgotPasswordOTPNotVerified
	}

	user.Password = newPassword
	user.ForgotVerified = false
	if err := r.DB.Save(&user).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}
