package admin

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/middlewares"
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
