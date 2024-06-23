package admin

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) CreateAccount(admin *entities.Admin) error {
	args := m.Called(admin)
	return args.Error(0)
}

func (m *MockAdminRepository) Login(admin *entities.Admin) error {
	args := m.Called(admin)
	return args.Error(0)
}

func (m *MockAdminRepository) GetAllAdmins() ([]*entities.Admin, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Admin), args.Error(1)
}

func (m *MockAdminRepository) GetAdminByID(id int) (*entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func (m *MockAdminRepository) DeleteAdmin(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAdminRepository) UpdateAdmin(id int, admin *entities.Admin) error {
	args := m.Called(id, admin)
	return args.Error(0)
}

func (m *MockAdminRepository) GetAdminByEmail(email string) (*entities.Admin, error) {
	args := m.Called(email)
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		mockAdminRepo.On("CreateAccount", &admin).Return(nil)

		result, err := AdminUseCase.CreateAccount(&admin)
		assert.NoError(t, err)
		assert.Equal(t, admin, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Name:            "",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		result, err := AdminUseCase.CreateAccount(&admin)
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed email already exists", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		mockAdminRepo.On("CreateAccount", &admin).Return(errors.New("Error 1062: Duplicate entry 'admin' for key 'email'"))

		result, err := AdminUseCase.CreateAccount(&admin)
		assert.Error(t, constants.ErrEmailAlreadyExists, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		mockAdminRepo.On("CreateAccount", &admin).Return(constants.ErrInternalServerError)

		result, err := AdminUseCase.CreateAccount(&admin)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed username already exists", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		mockAdminRepo.On("CreateAccount", &admin).Return(errors.New("Error 1062: Duplicate entry 'admin' for key 'username'"))

		result, err := AdminUseCase.CreateAccount(&admin)
		assert.Error(t, constants.ErrUsernameAlreadyExists, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed password must be at least 8 characters", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin",
			TelephoneNumber: "08123456789",
		}

		result, err := AdminUseCase.CreateAccount(&admin)
		assert.Error(t, constants.ErrPasswordMustBeAtLeast8Characters, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success admin", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Email:    "admin@gmail.com",
			Password: "admin",
		}

		mockAdminRepo.On("Login", &admin).Return(nil)

		result, err := AdminUseCase.Login(&admin)
		assert.NoError(t, err)
		assert.Equal(t, admin, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("success super admin", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Email:        "super_admin@gmail.com",
			Password:     "super_admin",
			IsSuperAdmin: true,
		}

		mockAdminRepo.On("Login", &admin).Return(nil)

		result, err := AdminUseCase.Login(&admin)
		assert.NoError(t, err)
		assert.Equal(t, admin, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Email:    "",
			Password: "admin",
		}

		result, err := AdminUseCase.Login(&admin)
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed invalid username or password", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			Email:    "admin@gmail.com",
			Password: "admin",
		}

		mockAdminRepo.On("Login", &admin).Return(constants.ErrInvalidUsernameOrPassword)

		result, err := AdminUseCase.Login(&admin)
		assert.Error(t, constants.ErrInvalidUsernameOrPassword, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})
}

func TestGetAllAdmins(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admins := []*entities.Admin{
			{
				Name:            "admin1",
				Email:           "admin1@gmail.com",
				Password:        "admin1",
				TelephoneNumber: "08123456789",
			},
			{
				Name:            "admin2",
				Email:           "admin2@gmail.com",
				Password:        "admin2",
				TelephoneNumber: "08123456789",
			},
		}

		adminValues := make([]entities.Admin, len(admins))
		for i, admin := range admins {
			adminValues[i] = *admin
		}

		mockAdminRepo.On("GetAllAdmins").Return(admins, nil)

		result, err := AdminUseCase.GetAllAdmins()
		assert.NoError(t, err)
		assert.Equal(t, adminValues, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		mockAdminRepo.On("GetAllAdmins").Return(([]*entities.Admin)(nil), constants.ErrInternalServerError)

		result, err := AdminUseCase.GetAllAdmins()
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Nil(t, result)

		mockAdminRepo.AssertExpectations(t)
	})
}

func TestGetAdminByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			ID:              1,
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return(&admin, nil)

		result, err := AdminUseCase.GetAdminByID(1)
		assert.NoError(t, err)
		assert.Equal(t, &admin, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed admin not found", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		mockAdminRepo.On("GetAdminByID", 1).Return((*entities.Admin)(nil), constants.ErrAdminNotFound)

		result, err := AdminUseCase.GetAdminByID(1)
		assert.Error(t, constants.ErrAdminNotFound, err)
		assert.Nil(t, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		mockAdminRepo.On("GetAdminByID", 1).Return((*entities.Admin)(nil), constants.ErrInternalServerError)

		result, err := AdminUseCase.GetAdminByID(1)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Nil(t, result)

		mockAdminRepo.AssertExpectations(t)
	})
}

func TestDeleteAdmin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		mockAdminRepo.On("GetAdminByID", 1).Return(&entities.Admin{}, nil)
		mockAdminRepo.On("DeleteAdmin", 1).Return(nil)

		err := AdminUseCase.DeleteAdmin(1)
		assert.NoError(t, err)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed admin not found", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		mockAdminRepo.On("GetAdminByID", 1).Return((*entities.Admin)(nil), constants.ErrAdminNotFound)

		err := AdminUseCase.DeleteAdmin(1)
		assert.Error(t, constants.ErrAdminNotFound, err)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		mockAdminRepo.On("GetAdminByID", 1).Return(&entities.Admin{}, nil)
		mockAdminRepo.On("DeleteAdmin", 1).Return(constants.ErrInternalServerError)

		err := AdminUseCase.DeleteAdmin(1)
		assert.Error(t, constants.ErrInternalServerError, err)

		mockAdminRepo.AssertExpectations(t)
	})
}

func TestUpdateAdmin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			ID:              1,
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		updatedAdmin := entities.Admin{
			ID:              1,
			Name:            "updated_admin",
			Email:           "updated_admin@gmail.com",
			Password:        "updated_admin",
			TelephoneNumber: "08123456780",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return(&admin, nil)
		mockAdminRepo.On("GetAdminByEmail", updatedAdmin.Email).Return((*entities.Admin)(nil), nil)
		mockAdminRepo.On("UpdateAdmin", 1, &updatedAdmin).Return(nil)

		result, err := AdminUseCase.UpdateAdmin(1, &updatedAdmin)
		assert.NoError(t, err)
		assert.Equal(t, updatedAdmin, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed admin not found", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		updatedAdmin := entities.Admin{
			ID:              1,
			Name:            "updated_admin",
			Email:           "admin@gmail.com",
			Password:        "updated_admin",
			TelephoneNumber: "08123456780",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return((*entities.Admin)(nil), constants.ErrAdminNotFound)

		result, err := AdminUseCase.UpdateAdmin(1, &updatedAdmin)
		assert.Error(t, constants.ErrAdminNotFound, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error when getting admin by email", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			ID:              1,
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		updatedAdmin := entities.Admin{
			ID:              1,
			Name:            "updated_admin",
			Email:           "admin@gmail.com",
			Password:        "updated_admin",
			TelephoneNumber: "08123456780",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return(&admin, constants.ErrInternalServerError)

		result, err := AdminUseCase.UpdateAdmin(1, &updatedAdmin)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed email already exists", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			ID:              1,
			Name:            "admin",
			Email:           "admin123@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		updatedAdmin := entities.Admin{
			ID:              1,
			Name:            "updated_admin",
			Email:           "admin@gmail.com",
			Password:        "updated_admin",
			TelephoneNumber: "08123456780",
		}

		conflictingAdmin := entities.Admin{
			ID:              2,
			Name:            "other_admin",
			Email:           "admin@gmail.com",
			Password:        "password",
			TelephoneNumber: "08123456780",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return(&admin, nil)
		mockAdminRepo.On("GetAdminByEmail", updatedAdmin.Email).Return(&conflictingAdmin, nil)

		result, err := AdminUseCase.UpdateAdmin(1, &updatedAdmin)
		assert.Equal(t, constants.ErrEmailAlreadyExists, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed no new data provided", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			ID:              1,
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		updatedAdmin := entities.Admin{
			ID:              0,
			Name:            "",
			Email:           "",
			Password:        "",
			TelephoneNumber: "",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return(&admin, nil)

		result, err := AdminUseCase.UpdateAdmin(1, &updatedAdmin)
		assert.Equal(t, constants.ErrNoChangesDetected, err)
		assert.Equal(t, admin, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed internal server error when updating admin", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			ID:              1,
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		updatedAdmin := entities.Admin{
			ID:              1,
			Name:            "updated_admin",
			Email:           "updated_admin@gmail.com",
			Password:        "updated_admin",
			TelephoneNumber: "08123456780",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return(&admin, nil)
		mockAdminRepo.On("GetAdminByEmail", updatedAdmin.Email).Return((*entities.Admin)(nil), nil)
		mockAdminRepo.On("UpdateAdmin", 1, &updatedAdmin).Return(constants.ErrInternalServerError)

		result, err := AdminUseCase.UpdateAdmin(1, &updatedAdmin)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("failed password must be at least 8 characters", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		AdminUseCase := NewAdminUseCase(mockAdminRepo)

		admin := entities.Admin{
			ID:              1,
			Name:            "admin",
			Email:           "admin@gmail.com",
			Password:        "admin12345",
			TelephoneNumber: "08123456789",
		}

		updatedAdmin := entities.Admin{
			ID:              1,
			Name:            "updated_admin",
			Email:           "admin@gmail.com",
			Password:        "admin",
			TelephoneNumber: "08123456780",
		}

		mockAdminRepo.On("GetAdminByID", 1).Return(&admin, nil)

		result, err := AdminUseCase.UpdateAdmin(1, &updatedAdmin)
		assert.Error(t, constants.ErrPasswordMustBeAtLeast8Characters, err)
		assert.Equal(t, entities.Admin{}, result)

		mockAdminRepo.AssertExpectations(t)
	})
}
