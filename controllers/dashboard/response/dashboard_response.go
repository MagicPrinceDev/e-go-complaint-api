package response

type DashboardResponse struct {
	TotalComplaints     int64                       `json:"totalComplaints"`
	ComplaintsByStatus  map[string]int64            `json:"complaintsByStatus"`
	UsersByYearAndMonth map[string][]MonthData      `json:"usersByYearAndMonth"`
	LatestComplaints    []NumberedComplaintResponse `json:"latestComplaints"`
}

type MonthData struct {
	Month string `json:"month"`
	Count int64  `json:"count"`
}

type User struct {
	Name string `json:"name"`
}

type Complaint struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}

type Category struct {
	Name string `json:"name"`
}

type NumberedComplaintResponse struct {
	No        int       `json:"no"`
	User      User      `json:"user"`
	Complaint Complaint `json:"complaint"`
	Category  Category  `json:"category"`
}

func NewDashboardResponse(totalComplaints int64, complaintsByStatus map[string]int64, usersByYearAndMonth map[string][]MonthData, latestComplaints []NumberedComplaintResponse) *DashboardResponse {
	return &DashboardResponse{
		TotalComplaints:     totalComplaints,
		ComplaintsByStatus:  complaintsByStatus,
		UsersByYearAndMonth: usersByYearAndMonth,
		LatestComplaints:    latestComplaints,
	}
}
