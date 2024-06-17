package user

import (
	"e-complaint-api/entities"
	"mime/multipart"

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
