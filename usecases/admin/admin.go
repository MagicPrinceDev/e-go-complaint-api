package admin

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/middlewares"
	"errors"
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
	if admin.Name == "" || admin.Email == "" || admin.Password == "" || admin.TelephoneNumber == "" {
		return entities.Admin{}, constants.ErrAllFieldsMustBeFilled
	}

	if len(admin.Password) < 8 {
		return entities.Admin{}, constants.ErrPasswordMustBeAtLeast8Characters
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
	if admin.Email == "" || admin.Password == "" {
		return entities.Admin{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Login(admin)
	if err != nil {
		return entities.Admin{}, constants.ErrInvalidUsernameOrPassword
	}

	if admin.IsSuperAdmin {
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Name, admin.Email, "super_admin")
	} else {
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Name, admin.Email, "admin")
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
	if admin == nil {
		return nil, err
	}

	return admin, nil
}

func (u *AdminUseCase) DeleteAdmin(id int) error {
	admin, _ := u.repository.GetAdminByID(id)

	if admin == nil {
		return constants.ErrAdminNotFound
	}

	err := u.repository.DeleteAdmin(id)
	if err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (u *AdminUseCase) UpdateAdmin(id int, admin *entities.Admin) (entities.Admin, error) {
	existingAdmin, err := u.repository.GetAdminByID(id)
	if err != nil {
		if errors.Is(err, constants.ErrAdminNotFound) {
			return entities.Admin{}, constants.ErrAdminNotFound
		}
		return entities.Admin{}, constants.ErrInternalServerError
	}

	// Check if the email is already taken by another admin
	if admin.Email != "" && admin.Email != existingAdmin.Email {
		conflictingAdmin, err := u.repository.GetAdminByEmail(admin.Email)
		if err == nil && conflictingAdmin != nil && conflictingAdmin.ID != id {
			return entities.Admin{}, constants.ErrEmailAlreadyExists
		}
	}

	isUpdated := false

	// Ensure existing data remains if no new data is provided
	if admin.Name != "" && admin.Name != existingAdmin.Name {
		existingAdmin.Name = admin.Name
		isUpdated = true
	}
	if admin.Email != "" && admin.Email != existingAdmin.Email {
		existingAdmin.Email = admin.Email
		isUpdated = true
	}
	if admin.TelephoneNumber != "" && admin.TelephoneNumber != existingAdmin.TelephoneNumber {
		existingAdmin.TelephoneNumber = admin.TelephoneNumber
		isUpdated = true
	}
	if admin.Password != "" {
		existingAdmin.Password = admin.Password
		isUpdated = true
	}

	if !isUpdated {
		return *existingAdmin, constants.ErrNoChangesDetected
	}

	if len(admin.Password) < 8 {
		return entities.Admin{}, constants.ErrPasswordMustBeAtLeast8Characters
	}

	err = u.repository.UpdateAdmin(id, existingAdmin)
	if err != nil {
		return entities.Admin{}, constants.ErrInternalServerError
	}

	return *existingAdmin, nil
}
