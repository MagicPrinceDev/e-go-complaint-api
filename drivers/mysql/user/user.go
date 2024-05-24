package user

import (
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"

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

	if err := r.DB.Where("username = ?", user.Username).First(&userDB).Error; err != nil {
		return errors.New("username or password is incorrect")
	}

	if !utils.CheckPasswordHash(user.Password, userDB.Password) {
		return errors.New("username or password is incorrect")
	}

	(*user).ID = userDB.ID
	(*user).Name = userDB.Name
	(*user).Username = userDB.Username

	return nil
}
