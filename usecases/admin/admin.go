package admin

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/middlewares"
	"e-complaint-api/utils"
	"strings"
)

type AdminUseCase struct {
	repository entities.AdminRepositoryInterface
}

func NewAdminUseCase(repository entities.AdminRepositoryInterface) *AdminUseCase {
	return &AdminUseCase{
		repository: repository,
	}
}

func (u *AdminUseCase) CreateAccount(admin *entities.Admin) (entities.Admin, error) {
	if admin.Name == "" || admin.Email == "" || admin.Password == "" || admin.Username == "" || admin.TelephoneNumber == "" {
		return entities.Admin{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.CreateAccount(admin)

	if err != nil {
		if strings.HasSuffix(err.Error(), "email'") {
			return entities.Admin{}, constants.ErrEmailAlreadyExists
		} else if strings.HasSuffix(err.Error(), "username'") {
			return entities.Admin{}, constants.ErrUsernameAlreadyExists
		} else {
			return entities.Admin{}, constants.ErrInternalServerError
		}
	}

	return *admin, nil
}

func (u *AdminUseCase) Login(admin *entities.Admin) (entities.Admin, error) {
	if admin.Username == "" || admin.Password == "" {
		return entities.Admin{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Login(admin)
	if err != nil {
		return entities.Admin{}, constants.ErrInvalidUsernameOrPassword
	}

	if admin.IsSuperAdmin {
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Username, "super_admin")
	} else {
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Username, "admin")
	}

	return *admin, nil
}

func (u *AdminUseCase) GetAllAdmins() ([]entities.Admin, error) {
	adminPtrs, err := u.repository.GetAllAdmins()
	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	adminValues := make([]entities.Admin, len(adminPtrs))
	for i, admin := range adminPtrs {
		adminValues[i] = *admin
	}

	return adminValues, nil
}

func (u *AdminUseCase) GetAdminByID(id int) (*entities.Admin, error) {
	admin, err := u.repository.GetAdminByID(id)
	if err != nil {
		return nil, constants.ErrInternalServerError
	}
	return admin, nil
}

func (u *AdminUseCase) DeleteAdmin(id int) error {
	err := u.repository.DeleteAdmin(id)
	if err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (u *AdminUseCase) UpdateAdmin(id int, admin *entities.Admin) (entities.Admin, error) {
	existingAdmin, err := u.repository.GetAdminByID(id)
	if err != nil {
		return entities.Admin{}, constants.ErrInternalServerError
	}

	// Ensure existing data remains if no new data is provided
	if admin.Name != "" {
		existingAdmin.Name = admin.Name
	}
	if admin.Email != "" {
		existingAdmin.Email = admin.Email
	}
	if admin.Password != "" {
		hash, _ := utils.HashPassword(admin.Password)
		existingAdmin.Password = hash
	}
	if admin.Username != "" {
		existingAdmin.Username = admin.Username
	}
	if admin.TelephoneNumber != "" {
		existingAdmin.TelephoneNumber = admin.TelephoneNumber
	}
	if admin.ProfilePhoto != "" {
		existingAdmin.ProfilePhoto = admin.ProfilePhoto
	}

	err = u.repository.UpdateAdmin(existingAdmin)
	if err != nil {
		return entities.Admin{}, constants.ErrInternalServerError
	}

	return *existingAdmin, nil
}
