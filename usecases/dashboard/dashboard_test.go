package dashboard_test

import (
	"e-complaint-api/controllers/dashboard/response"
	"e-complaint-api/entities"
	"e-complaint-api/usecases/dashboard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strconv"
	"testing"
)

type MockDashboardRepo struct {
	mock.Mock
}

func (m *MockDashboardRepo) GetTotalComplaints() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDashboardRepo) GetComplaintsByStatus() (map[string]int64, error) {
	args := m.Called()
	return args.Get(0).(map[string]int64), args.Error(1)
}

func (m *MockDashboardRepo) GetUsersByYearAndMonth() (map[string][]response.MonthData, error) {
	args := m.Called()
	return args.Get(0).(map[string][]response.MonthData), args.Error(1)
}

func (m *MockDashboardRepo) GetLatestComplaints(limit int) ([]entities.Complaint, error) {
	args := m.Called(limit)
	return args.Get(0).([]entities.Complaint), args.Error(1)
}

func TestDashboardUsecase_GetTotalComplaints(t *testing.T) {
	mockRepo := new(MockDashboardRepo)
	mockRepo.On("GetTotalComplaints").Return(int64(10), nil)

	uc := dashboard.NewDashboardUseCase(mockRepo)
	total, err := uc.GetTotalComplaints()

	assert.NoError(t, err)
	assert.Equal(t, int64(10), total)
}

func TestDashboardUsecase_GetComplaintsByStatus(t *testing.T) {
	mockRepo := new(MockDashboardRepo)
	mockRepo.On("GetComplaintsByStatus").Return(map[string]int64{"open": 5, "closed": 5}, nil)

	uc := dashboard.NewDashboardUseCase(mockRepo)
	statuses, err := uc.GetComplaintsByStatus()

	assert.NoError(t, err)
	assert.Equal(t, map[string]int64{"open": 5, "closed": 5}, statuses)
}

func TestDashboardUsecase_GetUsersByYearAndMonth(t *testing.T) {
	mockRepo := new(MockDashboardRepo)
	mockRepo.On("GetUsersByYearAndMonth").Return(map[string][]response.MonthData{"2022": {{Month: "January", Count: 10}}}, nil)

	uc := dashboard.NewDashboardUseCase(mockRepo)
	users, err := uc.GetUsersByYearAndMonth()

	assert.NoError(t, err)
	assert.Equal(t, map[string][]response.MonthData{"2022": {{Month: "January", Count: 10}}}, users)
}

func TestDashboardUsecase_GetLatestComplaints(t *testing.T) {
	mockRepo := new(MockDashboardRepo)
	mockRepo.On("GetLatestComplaints", 5).Return([]entities.Complaint{{ID: strconv.Itoa(1)}, {ID: strconv.Itoa(2)}, {ID: strconv.Itoa(3)}, {ID: strconv.Itoa(4)}, {ID: strconv.Itoa(5)}}, nil)

	uc := dashboard.NewDashboardUseCase(mockRepo)
	complaints, err := uc.GetLatestComplaints(5)

	assert.NoError(t, err)
	assert.Equal(t, []entities.Complaint{{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}, {ID: "5"}}, complaints)
}
