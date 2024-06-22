package dashboard

import (
	"e-complaint-api/controllers/dashboard/response"
	"e-complaint-api/entities"
)

type DashboardRepoInterface interface {
	GetTotalComplaints() (int64, error)
	GetComplaintsByStatus() (map[string]int64, error)
	GetUsersByYearAndMonth() (map[string][]response.MonthData, error)
	GetLatestComplaints(limit int) ([]entities.Complaint, error)
}

type DashboardUsecase struct {
	DashboardRepo DashboardRepoInterface
}

func NewDashboardUseCase(dashboardRepo DashboardRepoInterface) *DashboardUsecase {
	return &DashboardUsecase{DashboardRepo: dashboardRepo}
}

func (uc *DashboardUsecase) GetTotalComplaints() (int64, error) {
	return uc.DashboardRepo.GetTotalComplaints()
}

func (uc *DashboardUsecase) GetComplaintsByStatus() (map[string]int64, error) {
	return uc.DashboardRepo.GetComplaintsByStatus()
}

func (uc *DashboardUsecase) GetUsersByYearAndMonth() (map[string][]response.MonthData, error) {
	return uc.DashboardRepo.GetUsersByYearAndMonth()
}

func (uc *DashboardUsecase) GetLatestComplaints(limit int) ([]entities.Complaint, error) {
	return uc.DashboardRepo.GetLatestComplaints(limit)
}
