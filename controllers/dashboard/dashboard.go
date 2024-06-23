package dashboard

import (
	"e-complaint-api/controllers/base"
	"e-complaint-api/controllers/dashboard/response"
	"e-complaint-api/usecases/dashboard"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DashboardController struct {
	DashboardUsecase dashboard.DashboardUsecase
}

func NewDashboardController(dashboardUsecase dashboard.DashboardUsecase) *DashboardController {
	return &DashboardController{DashboardUsecase: dashboardUsecase}
}

func (ctrl *DashboardController) GetDashboardData(c echo.Context) error {
	totalComplaints, _ := ctrl.DashboardUsecase.GetTotalComplaints()
	complaintsByStatus, _ := ctrl.DashboardUsecase.GetComplaintsByStatus()
	usersByYearAndMonth, _ := ctrl.DashboardUsecase.GetUsersByYearAndMonth()
	latestComplaints, _ := ctrl.DashboardUsecase.GetLatestComplaints(5)

	numberedLatestComplaints := make([]response.NumberedComplaintResponse, len(latestComplaints))
	for i, complaint := range latestComplaints {
		numberedLatestComplaints[i] = response.NumberedComplaintResponse{
			No: i + 1,
			User: response.User{
				Name: complaint.User.Name,
			},
			Complaint: response.Complaint{
				Date:   complaint.CreatedAt.Format("2 January 2006 15:04:05"),
				Status: complaint.Status,
			},
			Category: response.Category{
				Name: complaint.Category.Name,
			},
		}
	}

	resp := response.DashboardResponse{
		TotalComplaints:     totalComplaints,
		ComplaintsByStatus:  complaintsByStatus,
		UsersByYearAndMonth: usersByYearAndMonth,
		LatestComplaints:    numberedLatestComplaints,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Dashboard data retrieved successfully", resp))

}
