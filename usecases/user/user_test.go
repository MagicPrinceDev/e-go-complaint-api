package user

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Login(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAllUsers() ([]*entities.User, error) {
	args := m.Called()
	return args.Get(0).([]*entities.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(id int, user *entities.User) error {
	args := m.Called(id, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateProfilePhoto(id int, photo string) error {
	args := m.Called(id, photo)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) UpdatePassword(id int, newPassword string) error {
	args := m.Called(id, newPassword)
	return args.Error(0)
}

func (m *MockUserRepository) SendOTP(email, otp string) error {
	args := m.Called(email, otp)
	return args.Error(0)
}

func (m *MockUserRepository) VerifyOTPRegister(email, otp string) error {
	args := m.Called(email, otp)
	return args.Error(0)
}

func (m *MockUserRepository) VerifyOTPForgotPassword(email, otp string) error {
	args := m.Called(email, otp)
	return args.Error(0)
}

func (m *MockUserRepository) UpdatePasswordForgot(email, newPassword string) error {
	args := m.Called(email, newPassword)
	return args.Error(0)
}

type MockMailTrapAPI struct {
	mock.Mock
}

func (m *MockMailTrapAPI) SendOTP(email, otp, otpType string) error {
	args := m.Called(email, otp, otpType)
	return args.Error(0)
}

type MockUserGCSAPI struct {
	mock.Mock
}

func (m *MockUserGCSAPI) Upload(files []*multipart.FileHeader) ([]string, error) {
	args := m.Called(files)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockUserGCSAPI) Delete(filePaths []string) error {
	args := m.Called(filePaths)
	return args.Error(0)
}

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:           "user@example.com",
			Password:        "password",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("Register", &user).Return(nil)

		result, err := userUseCase.Register(&user)
		assert.NoError(t, err)
		assert.Equal(t, &result, &user)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:           "",
			Password:        "password",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		result, err := userUseCase.Register(&user)
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed email already exists", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:           "user@gmail.com",
			Password:        "password",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("Register", &user).Return(errors.New("Error 1062 email'"))

		result, err := userUseCase.Register(&user)
		assert.Error(t, constants.ErrEmailAlreadyExists, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed username already exists", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:           "user@gmail.com",
			Password:        "password",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("Register", &user).Return(errors.New("Error 1062 username'"))

		result, err := userUseCase.Register(&user)
		assert.Error(t, constants.ErrUsernameAlreadyExists, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:           "user@gmail.com",
			Password:        "password",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("Register", &user).Return(constants.ErrInternalServerError)

		result, err := userUseCase.Register(&user)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed password must be at least 8 characters", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:           "user@gmail.com",
			Password:        "pass",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		result, err := userUseCase.Register(&user)
		assert.Error(t, constants.ErrPasswordMustBeAtLeast8Characters, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:    "user@gmail.com",
			Password: "password",
		}

		mockUserRepository.On("Login", &user).Return(nil)

		result, err := userUseCase.Login(&user)
		assert.NoError(t, err)
		assert.NotEmpty(t, result.Token)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:    "",
			Password: "password",
		}

		result, err := userUseCase.Login(&user)
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed invalid username or password", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			Email:    "user123@gmail.com",
			Password: "password",
		}

		mockUserRepository.On("Login", &user).Return(constants.ErrInvalidUsernameOrPassword)

		result, err := userUseCase.Login(&user)
		assert.Error(t, constants.ErrInvalidUsernameOrPassword, err)
		assert.Equal(t, entities.User{}, result)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		users := []*entities.User{
			{
				ID:              1,
				Email:           "user1@gmail.com",
				Name:            "User 1",
				TelephoneNumber: "081234567890",
				ProfilePhoto:    "profile_photo.jpg",
			},
			{
				ID:              2,
				Email:           "user2@gmail.com",
				Name:            "User 2",
				TelephoneNumber: "081234567891",
				ProfilePhoto:    "profile_photo.jpg",
			},
		}

		mockUserRepository.On("GetAllUsers").Return(users, nil)

		result, err := userUseCase.GetAllUsers()
		assert.NoError(t, err)
		assert.Equal(t, users, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("GetAllUsers").Return(([]*entities.User)(nil), constants.ErrInternalServerError)

		result, err := userUseCase.GetAllUsers()
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Nil(t, result)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			ID:              1,
			Email:           "user1@gmail.com",
			Name:            "User 1",
			TelephoneNumber: "081234567890",
			ProfilePhoto:    "profile_photo.jpg",
		}

		mockUserRepository.On("GetUserByID", 1).Return(&user, nil)

		result, err := userUseCase.GetUserByID(1)
		assert.NoError(t, err)
		assert.Equal(t, &user, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed user not found", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("GetUserByID", 1).Return((*entities.User)(nil), constants.ErrUserNotFound)

		result, err := userUseCase.GetUserByID(1)
		assert.Error(t, constants.ErrUserNotFound, err)
		assert.Nil(t, result)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			ID:              1,
			Email:           "user@gmail.com",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("GetUserByID", 1).Return(&user, nil)
		mockUserRepository.On("UpdateUser", 1, &user).Return(nil)

		result, err := userUseCase.UpdateUser(1, &user)
		assert.NoError(t, err)
		assert.Equal(t, user, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			ID:              1,
			Email:           "",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		result, err := userUseCase.UpdateUser(1, &user)
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed user not found", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			ID:              1,
			Email:           "user@gmail.com",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("GetUserByID", 1).Return((*entities.User)(nil), constants.ErrUserNotFound)

		result, err := userUseCase.UpdateUser(1, &user)
		assert.Error(t, constants.ErrUserNotFound, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			ID:              1,
			Email:           "user@gmail.com",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("GetUserByID", 1).Return(&user, nil)
		mockUserRepository.On("UpdateUser", 1, &user).Return(constants.ErrInternalServerError)

		result, err := userUseCase.UpdateUser(1, &user)
		assert.Error(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed email already exists", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			ID:              1,
			Email:           "user@gmail.com",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("GetUserByID", 1).Return(&user, nil)
		mockUserRepository.On("UpdateUser", 1, &user).Return(errors.New("Error 1062 email'"))

		result, err := userUseCase.UpdateUser(1, &user)
		assert.Error(t, constants.ErrEmailAlreadyExists, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed username already exists", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		user := entities.User{
			ID:              1,
			Email:           "user@gmail.com",
			Name:            "User",
			TelephoneNumber: "081234567890",
		}

		mockUserRepository.On("GetUserByID", 1).Return(&user, nil)
		mockUserRepository.On("UpdateUser", 1, &user).Return(errors.New("Error 1062 username'"))

		result, err := userUseCase.UpdateUser(1, &user)
		assert.Error(t, constants.ErrUsernameAlreadyExists, err)
		assert.Equal(t, entities.User{}, result)

		mockUserRepository.AssertExpectations(t)
	})

}

func TestUpdateProfilePhoto(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		fileHeader := &multipart.FileHeader{
			Filename: "profile_photo.jpg",
		}

		fileHeaders := []*multipart.FileHeader{fileHeader}

		mockUserGCSAPI.On("Upload", fileHeaders).Return([]string{"profile_photo.jpg"}, nil)
		mockUserRepository.On("UpdateProfilePhoto", 1, "profile_photo.jpg").Return(nil)

		err := userUseCase.UpdateProfilePhoto(1, fileHeader)
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed gcs internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		fileHeader := &multipart.FileHeader{
			Filename: "profile_photo.jpg",
		}

		fileHeaders := []*multipart.FileHeader{fileHeader}

		mockUserGCSAPI.On("Upload", fileHeaders).Return(([]string)(nil), constants.ErrInternalServerError)

		err := userUseCase.UpdateProfilePhoto(1, fileHeader)
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		fileHeader := &multipart.FileHeader{
			Filename: "profile_photo.jpg",
		}

		fileHeaders := []*multipart.FileHeader{fileHeader}

		mockUserGCSAPI.On("Upload", fileHeaders).Return([]string{"profile_photo.jpg"}, nil)
		mockUserRepository.On("UpdateProfilePhoto", 1, "profile_photo.jpg").Return(constants.ErrInternalServerError)

		err := userUseCase.UpdateProfilePhoto(1, fileHeader)
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("GetUserByID", 1).Return(&entities.User{}, nil)
		mockUserRepository.On("Delete", 1).Return(nil)

		err := userUseCase.Delete(1)
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("GetUserByID", 1).Return(&entities.User{}, nil)
		mockUserRepository.On("Delete", 1).Return(constants.ErrInternalServerError)

		err := userUseCase.Delete(1)
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error get user by id", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("GetUserByID", 1).Return((*entities.User)(nil), constants.ErrInternalServerError)

		err := userUseCase.Delete(1)
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed user not found", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("GetUserByID", 1).Return((*entities.User)(nil), constants.ErrUserNotFound)

		err := userUseCase.Delete(1)
		assert.Error(t, constants.ErrUserNotFound, err)

		mockUserRepository.AssertExpectations(t)
	})

}

func TestUpdatePassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("UpdatePassword", 1, mock.Anything).Return(nil)

		err := userUseCase.UpdatePassword(1, "password")
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		err := userUseCase.UpdatePassword(1, "")
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed password must be at least 8 characters", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		err := userUseCase.UpdatePassword(1, "pass")
		assert.Error(t, constants.ErrPasswordMustBeAtLeast8Characters, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestSendOTP(t *testing.T) {
	t.Run("success register", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("SendOTP", "user@gmail.com", mock.Anything).Return(nil)
		mockMailTrapAPI.On("SendOTP", "user@gmail.com", mock.Anything, "register").Return(nil)

		err := userUseCase.SendOTP("user@gmail.com", "register")
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("success forgot password", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("SendOTP", "user@gmail.com", mock.Anything).Return(nil)
		mockMailTrapAPI.On("SendOTP", "user@gmail.com", mock.Anything, "forgot_password").Return(nil)

		err := userUseCase.SendOTP("user@gmail.com", "forgot_password")
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		err := userUseCase.SendOTP("", "register")
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("SendOTP", "user@gmail.com", mock.Anything).Return(constants.ErrInternalServerError)

		err := userUseCase.SendOTP("user@gmail.com", "register")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed email not registered", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("SendOTP", "user@gmail.com", mock.Anything).Return(constants.ErrEmailNotRegistered)

		err := userUseCase.SendOTP("user@gmail.com", "register")
		assert.Error(t, constants.ErrEmailNotRegistered, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed mailtrap api error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("SendOTP", "user@gmail.com", mock.Anything).Return(nil)
		mockMailTrapAPI.On("SendOTP", "user@gmail.com", mock.Anything, "register").Return(errors.New("Mailtrap API error"))

		err := userUseCase.SendOTP("user@gmail.com", "register")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestVerifyOTP(t *testing.T) {
	t.Run("success register", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("VerifyOTPRegister", "user@gmail.com", "12345").Return(nil)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "register")
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("success forgot password", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)
		mockUserRepository.On("VerifyOTPForgotPassword", "user@gmail.com", "12345").Return(nil)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "forgot_password")
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		err := userUseCase.VerifyOTP("", "12345", "register")
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed register email not registered", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("VerifyOTPRegister", "user@gmail.com", "12345").Return(constants.ErrEmailNotRegistered)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "register")
		assert.Error(t, constants.ErrEmailNotRegistered, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed forgot password email not registered", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("VerifyOTPForgotPassword", "user@gmail.com", "12345").Return(constants.ErrEmailNotRegistered)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "forgot_password")
		assert.Error(t, constants.ErrEmailNotRegistered, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed register invalid otp", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("VerifyOTPRegister", "user@gmail.com", "12345").Return(constants.ErrInvalidOTP)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "register")
		assert.Error(t, constants.ErrInvalidOTP, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed forgot password invalid otp", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("VerifyOTPForgotPassword", "user@gmail.com", "12345").Return(constants.ErrInvalidOTP)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "forgot_password")
		assert.Error(t, constants.ErrInvalidOTP, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error register", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("VerifyOTPRegister", "user@gmail.com", "12345").Return(constants.ErrInternalServerError)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "register")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error forgot password", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)
		mockUserRepository.On("VerifyOTPForgotPassword", "user@gmail.com", "12345").Return(constants.ErrInternalServerError)

		err := userUseCase.VerifyOTP("user@gmail.com", "12345", "forgot_password")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})

}

func TestUpdatePasswordForgot(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("UpdatePasswordForgot", "user@gmail.com", mock.Anything).Return(nil)

		err := userUseCase.UpdatePasswordForgot("user@gmail.com", "password")
		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed empty field", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		err := userUseCase.UpdatePasswordForgot("", "password")
		assert.Error(t, constants.ErrAllFieldsMustBeFilled, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed user not found", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("UpdatePasswordForgot", "user@gmail.com", mock.Anything).Return(constants.ErrUserNotFound)

		err := userUseCase.UpdatePasswordForgot("user@gmail.com", "password")
		assert.Error(t, constants.ErrUserNotFound, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("UpdatePasswordForgot", "user@gmail.com", mock.Anything).Return(constants.ErrInternalServerError)

		err := userUseCase.UpdatePasswordForgot("user@gmail.com", "password")
		assert.Error(t, constants.ErrInternalServerError, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed otp not verified", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		mockUserRepository.On("UpdatePasswordForgot", "user@gmail.com", mock.Anything).Return(constants.ErrForgotPasswordOTPNotVerified)

		err := userUseCase.UpdatePasswordForgot("user@gmail.com", "password")
		assert.Error(t, constants.ErrForgotPasswordOTPNotVerified, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed password must be at least 8 characters", func(t *testing.T) {
		mockUserRepository := new(MockUserRepository)
		mockMailTrapAPI := new(MockMailTrapAPI)
		mockUserGCSAPI := new(MockUserGCSAPI)
		userUseCase := NewUserUseCase(mockUserRepository, mockMailTrapAPI, mockUserGCSAPI)

		err := userUseCase.UpdatePasswordForgot("user@gmail.com", "pass")
		assert.Error(t, constants.ErrPasswordMustBeAtLeast8Characters, err)

		mockUserRepository.AssertExpectations(t)
	})
}
