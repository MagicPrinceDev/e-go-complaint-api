package complaint

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"time"

	"gorm.io/gorm"
)

type ComplaintRepo struct {
	DB *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) *ComplaintRepo {
	return &ComplaintRepo{DB: db}
}

func (r *ComplaintRepo) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	var complaints []entities.Complaint
	query := r.DB

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("description LIKE ? OR address LIKE ? OR id LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	query = query.Order(sortBy + " " + sortType)

	if limit != 0 && page != 0 {
		query = query.Limit(limit).Offset((page - 1) * limit)
	}

	if err := query.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}

func (r *ComplaintRepo) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var totalData int64

	query := r.DB.Model(&entities.Complaint{})

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("description LIKE ? OR address LIKE ? OR id LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&totalData).Error; err != nil {
		return entities.Metadata{}, err
	}

	metadata := entities.Metadata{
		TotalData: int(totalData),
	}

	return metadata, nil
}

func (r *ComplaintRepo) GetByID(id string) (entities.Complaint, error) {
	var complaint entities.Complaint

	if err := r.DB.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Where("id = ?", id).First(&complaint).Error; err != nil {
		return entities.Complaint{}, constants.ErrComplaintNotFound
	}

	return complaint, nil
}

func (r *ComplaintRepo) GetByUserID(userId int) ([]entities.Complaint, error) {
	var complaints []entities.Complaint

	if err := r.DB.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Where("user_id = ?", userId).Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}

func (r *ComplaintRepo) Create(complaint *entities.Complaint) error {
	if err := r.DB.Create(complaint).Error; err != nil {
		return err
	}

	if err := r.DB.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Where("id = ?", complaint.ID).First(complaint).Error; err != nil {
		return err
	}

	return nil
}

func (r *ComplaintRepo) Delete(id string, userId int) error {
	var complaint entities.Complaint

	if err := r.DB.Where("id = ?", id).First(&complaint).Error; err != nil {
		return constants.ErrComplaintNotFound
	}

	if complaint.UserID != userId {
		return constants.ErrUnauthorized
	}

	complaint.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := r.DB.Save(&complaint).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *ComplaintRepo) AdminDelete(id string) error {
	var complaint entities.Complaint

	if err := r.DB.Where("id = ?", id).First(&complaint).Error; err != nil {
		return constants.ErrComplaintNotFound
	}

	complaint.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := r.DB.Save(&complaint).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *ComplaintRepo) Update(complaint entities.Complaint) (entities.Complaint, error) {
	var oldComplaint entities.Complaint

	if err := r.DB.Where("id = ?", complaint.ID).First(&oldComplaint).Error; err != nil {
		return entities.Complaint{}, constants.ErrComplaintNotFound
	}

	if oldComplaint.UserID != complaint.UserID {
		return entities.Complaint{}, constants.ErrUnauthorized
	}

	oldComplaint.Description = complaint.Description
	oldComplaint.Type = complaint.Type
	oldComplaint.CategoryID = complaint.CategoryID
	oldComplaint.RegencyID = complaint.RegencyID
	oldComplaint.Address = complaint.Address
	oldComplaint.Date = complaint.Date

	if err := r.DB.Save(&oldComplaint).Error; err != nil {
		return entities.Complaint{}, err
	}

	if err := r.DB.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Where("id = ?", oldComplaint.ID).First(&complaint).Error; err != nil {
		return entities.Complaint{}, constants.ErrInternalServerError
	}

	return complaint, nil
}

func (r *ComplaintRepo) UpdateStatus(id string, status string) error {
	var complaint entities.Complaint

	if err := r.DB.Where("id = ?", id).First(&complaint).Error; err != nil {
		return constants.ErrComplaintNotFound
	}

	complaint.Status = status
	if err := r.DB.Save(&complaint).Error; err != nil {
		return err
	}

	return nil
}

func (r *ComplaintRepo) GetStatus(id string) (string, error) {
	var status string

	if err := r.DB.Model(&entities.Complaint{}).Select("status").Where("id = ?", id).Scan(&status).Error; err != nil {
		return "", constants.ErrComplaintNotFound
	}

	return status, nil
}

func (r *ComplaintRepo) Import(complaints []entities.Complaint) error {
	if err := r.DB.CreateInBatches(complaints, len(complaints)).Error; err != nil {
		return err
	}

	return nil
}

func (r *ComplaintRepo) IncreaseTotalLikes(id string) error {
	if err := r.DB.Model(&entities.Complaint{}).Where("id = ?", id).Update("total_likes", gorm.Expr("total_likes + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

func (r *ComplaintRepo) DecreaseTotalLikes(id string) error {
	if err := r.DB.Model(&entities.Complaint{}).Where("id = ?", id).Update("total_likes", gorm.Expr("total_likes - ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

func (r *ComplaintRepo) GetComplaintIDsByUserID(userID int) ([]string, error) {
	var complaintIDs []string

	if err := r.DB.Model(&entities.Complaint{}).Select("id").Where("user_id = ?", userID).Find(&complaintIDs).Error; err != nil {
		return nil, err
	}

	return complaintIDs, nil
}
