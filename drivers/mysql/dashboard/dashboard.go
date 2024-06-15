package dashboard

import (
	"e-complaint-api/controllers/dashboard/response"
	"e-complaint-api/entities"
	"gorm.io/gorm"
	"strconv"
)

type DashboardRepo struct {
	DB *gorm.DB
}

func NewDashboardRepo(db *gorm.DB) *DashboardRepo {
	return &DashboardRepo{DB: db}
}

func (repo *DashboardRepo) GetTotalComplaints() (int64, error) {
	var totalComplaints int64
	if err := repo.DB.Model(&entities.Complaint{}).Count(&totalComplaints).Error; err != nil {
		return 0, err
	}
	return totalComplaints, nil
}

func (repo *DashboardRepo) GetComplaintsByStatus() (map[string]int64, error) {
	var results []struct {
		Status string
		Count  int64
	}

	if err := repo.DB.Model(&entities.Complaint{}).Select("status, count(*) as count").Group("status").Scan(&results).Error; err != nil {
		return nil, err
	}

	complaintsByStatus := make(map[string]int64)
	for _, result := range results {
		complaintsByStatus[result.Status] = result.Count
	}
	return complaintsByStatus, nil
}

func (repo *DashboardRepo) GetTotalUsers() (int64, error) {
	var totalUsers int64
	if err := repo.DB.Model(&entities.User{}).Count(&totalUsers).Error; err != nil {
		return 0, err
	}
	return totalUsers, nil
}

func monthNumberToName(monthNumber int) string {
	monthNames := []string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	return monthNames[monthNumber]
}

func (repo *DashboardRepo) GetUsersByYearAndMonth() (map[string][]response.MonthData, error) {
	var results []struct {
		Year  int
		Month int
		Count int64
	}

	if err := repo.DB.Model(&entities.User{}).Select("EXTRACT(YEAR FROM created_at) as year, EXTRACT(MONTH FROM created_at) as month, count(*) as count").Group("year, month").Scan(&results).Error; err != nil {
		return nil, err
	}

	usersByYearAndMonth := make(map[string][]response.MonthData)
	for _, result := range results {
		year := strconv.Itoa(result.Year)
		if _, ok := usersByYearAndMonth[year]; !ok {
			usersByYearAndMonth[year] = make([]response.MonthData, 12)
			for m := 1; m <= 12; m++ {
				usersByYearAndMonth[year][m-1] = response.MonthData{
					Month: monthNumberToName(m),
					Count: 0,
				}
			}
		}
		usersByYearAndMonth[year][result.Month-1].Count = result.Count
	}
	return usersByYearAndMonth, nil
}

func (repo *DashboardRepo) GetLatestComplaints(limit int) ([]entities.Complaint, error) {
	var complaints []entities.Complaint

	if err := repo.DB.Preload("Category").Preload("User").Order("created_at desc").Limit(limit).Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}
