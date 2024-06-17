package dashboard

import (
	"e-complaint-api/controllers/dashboard/response"
	"e-complaint-api/drivers/mysql/dashboard"
	"e-complaint-api/entities"
)

type DashboardUsecase struct {
	DashboardRepo dashboard.DashboardRepo
}

func NewDashboardUseCase(dashboardRepo dashboard.DashboardRepo) *DashboardUsecase {
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
