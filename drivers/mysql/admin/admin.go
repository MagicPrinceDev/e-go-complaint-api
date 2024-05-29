package admin

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"errors"

	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{DB: db}
}

func (r *AdminRepo) CreateAccount(admin *entities.Admin) error {
	hash, _ := utils.HashPassword(admin.Password)
	(*admin).Password = hash

	if err := r.DB.Create(&admin).Error; err != nil {
		return err
	}

	return nil
}

func (r *AdminRepo) Login(admin *entities.Admin) error {
	var adminDB entities.Admin

	if err := r.DB.Where("username = ?", admin.Username).First(&adminDB).Error; err != nil {
		return errors.New("username or password is incorrect")
	}

	if !utils.CheckPasswordHash(admin.Password, adminDB.Password) {
		return errors.New("username or password is incorrect")
	}

	(*admin).ID = adminDB.ID
	(*admin).Username = adminDB.Username
	(*admin).IsSuperAdmin = adminDB.IsSuperAdmin

	return nil
}

func (r *AdminRepo) GetAllAdmins() ([]*entities.Admin, error) {
	var admins []*entities.Admin
	err := r.DB.Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *AdminRepo) GetAdminByID(id int) (*entities.Admin, error) {
	admin := &entities.Admin{}
	result := r.DB.First(admin, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, constants.ErrAdminNotFound
		}
		return nil, result.Error
	}
	return admin, nil
}

func (r *AdminRepo) DeleteAdmin(id int) error {
	if err := r.DB.Delete(&entities.Admin{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *AdminRepo) UpdateAdmin(id int, admin *entities.Admin) error {
	if err := r.DB.Model(&entities.Admin{}).Where("id = ?", id).Updates(&admin).Error; err != nil {
		return err
	}

	return nil
}

func (r *AdminRepo) UpdatePassword(id int, newPassword string) error {
	return r.DB.Model(&entities.Admin{}).Where("id = ?", id).Updates(&entities.Admin{Password: newPassword}).Error
}

func (r *AdminRepo) GetAdminByEmail(email string) (*entities.Admin, error) {
	var admin entities.Admin
	if err := r.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepo) GetAdminByUsername(username string) (*entities.Admin, error) {
	var admin entities.Admin
	if err := r.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}
