package user

import (
	"e-complaint-api/constants"
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
			return nil, constants.ErrNotFound
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

func (r *UserRepo) Delete(id int) error {
	if err := r.DB.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) UpdatePassword(id int, newPassword string) error {
	return r.DB.Model(&entities.User{}).Where("id = ?", id).Updates(&entities.User{Password: newPassword}).Error
}
