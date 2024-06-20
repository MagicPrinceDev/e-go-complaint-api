package complaint

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type ComplaintUseCase struct {
	complaintRepo     entities.ComplaintRepositoryInterface
	complaintFileRepo entities.ComplaintFileRepositoryInterface
	getRowsFromExcel  func(file *multipart.FileHeader) ([][]string, error)
}

func NewComplaintUseCase(complaintRepo entities.ComplaintRepositoryInterface, complaintFileRepo entities.ComplaintFileRepositoryInterface) *ComplaintUseCase {
	return &ComplaintUseCase{
		complaintRepo:     complaintRepo,
		complaintFileRepo: complaintFileRepo,
		getRowsFromExcel:  utils.GetRowsFromExcel,
	}
}

func (u *ComplaintUseCase) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	if limit != 0 && page == 0 {
		return nil, constants.ErrPageMustBeFilled
	} else if limit == 0 && page != 0 {
		return nil, constants.ErrLimitMustBeFilled
	}

	if sortBy == "" {
		sortBy = "created_at"
	}

	if sortType == "" {
		sortType = "DESC"
	}

	complaints, err := u.complaintRepo.GetPaginated(limit, page, search, filter, sortBy, sortType)
	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return complaints, nil
}

func (u *ComplaintUseCase) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var pagination entities.Pagination
	metaData, err := u.complaintRepo.GetMetaData(limit, page, search, filter)

	if err != nil {
		return entities.Metadata{}, constants.ErrInternalServerError
	}

	if limit != 0 && page != 0 {
		pagination.FirstPage = 1
		pagination.LastPage = (metaData.TotalData + limit - 1) / limit
		pagination.CurrentPage = page
		if metaData.TotalData == 0 {
			pagination.TotalDataPerPage = 0
			pagination.LastPage = 1
		} else if pagination.CurrentPage == pagination.LastPage {
			pagination.TotalDataPerPage = metaData.TotalData - (pagination.LastPage-1)*limit
		} else {
			pagination.TotalDataPerPage = limit
		}

		if page > 1 {
			pagination.PrevPage = page - 1
		} else {
			pagination.PrevPage = 0
		}

		if page < pagination.LastPage {
			pagination.NextPage = page + 1
		} else {
			pagination.NextPage = 0
		}
	} else {
		pagination.FirstPage = 1
		pagination.LastPage = 1
		pagination.CurrentPage = 1
		pagination.TotalDataPerPage = metaData.TotalData
		pagination.PrevPage = 0
		pagination.NextPage = 0
	}
	metaData.Pagination = pagination

	return metaData, nil
}

func (u *ComplaintUseCase) GetByID(id string) (entities.Complaint, error) {
	complaint, err := u.complaintRepo.GetByID(id)
	if err != nil {
		return entities.Complaint{}, err
	}

	return complaint, nil
}

func (u *ComplaintUseCase) GetByUserID(userId int) ([]entities.Complaint, error) {
	complaints, err := u.complaintRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func (u *ComplaintUseCase) Create(complaint *entities.Complaint) (entities.Complaint, error) {
	if complaint.CategoryID == 0 || complaint.UserID == 0 || complaint.RegencyID == "" || complaint.Description == "" || complaint.Address == "" || complaint.Type == "" || complaint.Date.IsZero() {
		return entities.Complaint{}, constants.ErrAllFieldsMustBeFilled
	}
	(*complaint).ID = utils.GenerateID("C-", 10)

	err := u.complaintRepo.Create(complaint)
	if err != nil {
		if strings.HasSuffix(err.Error(), "REFERENCES `regencies` (`id`))") {
			return entities.Complaint{}, constants.ErrRegencyNotFound
		} else if strings.HasSuffix(err.Error(), "REFERENCES `categories` (`id`))") {
			return entities.Complaint{}, constants.ErrCategoryNotFound
		} else {
			return entities.Complaint{}, constants.ErrInternalServerError
		}
	}

	return *complaint, nil
}

func (u *ComplaintUseCase) Delete(id string, userId int, role string) error {
	if role == "admin" || role == "super_admin" {
		err := u.complaintRepo.AdminDelete(id)
		if err != nil {
			return err
		}
	} else {
		err := u.complaintRepo.Delete(id, userId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *ComplaintUseCase) Update(complaint entities.Complaint) (entities.Complaint, error) {
	if complaint.CategoryID == 0 || complaint.UserID == 0 || complaint.RegencyID == "" || complaint.Description == "" || complaint.Address == "" || complaint.Type == "" || complaint.Date.IsZero() {
		return entities.Complaint{}, constants.ErrAllFieldsMustBeFilled
	}

	complaint, err := u.complaintRepo.Update(complaint)
	if err != nil {
		if strings.HasSuffix(err.Error(), "REFERENCES `regencies` (`id`))") {
			return entities.Complaint{}, constants.ErrRegencyNotFound
		} else if strings.HasSuffix(err.Error(), "REFERENCES `categories` (`id`))") {
			return entities.Complaint{}, constants.ErrCategoryNotFound
		} else {
			return entities.Complaint{}, err
		}
	}

	return complaint, nil
}

func (u *ComplaintUseCase) UpdateStatus(id string, status string) error {
	if status != "Pending" && status != "Verifikasi" && status != "On Progress" && status != "Selesai" && status != "Ditolak" {
		return constants.ErrInvalidStatus
	}

	if id == "" {
		return constants.ErrIDMustBeFilled
	}

	err := u.complaintRepo.UpdateStatus(id, status)
	if err != nil {
		return err
	}

	return nil
}

func (u *ComplaintUseCase) Import(file *multipart.FileHeader) error {
	rows, err := u.getRowsFromExcel(file)
	if err != nil {
		return err
	}

	var complaints []entities.Complaint
	var complaintFiles []entities.ComplaintFile
	var process []entities.ComplaintProcess

	// Loop through each row in the sheet
	for i, row := range rows {
		if i == 0 {
			// Skip the header row
			continue
		}

		if len(row) < 9 {
			// Ensure the row has enough columns
			return constants.ErrColumnsDoesntMatch
		}

		userId, err := strconv.Atoi(row[0])
		if err != nil {
			return constants.ErrInvalidIDFormat
		}

		categoryId, err := strconv.Atoi(row[1])
		if err != nil {
			return constants.ErrInvalidCategoryIDFormat
		}

		regencyId := row[2]
		address := row[3]
		description := row[4]
		status := row[5]
		typeComplaint := row[6]
		date, _ := time.Parse("02-01-2006", row[7])
		pathFiles := row[8]

		process = []entities.ComplaintProcess{}
		if status == "Verifikasi" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Pending",
				Message: "Aduan anda sedang dalam proses verifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Verifikasi",
				Message: "Aduan anda telah diverifikasi oleh admin kami",
			})
		} else if status == "On Progress" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Pending",
				Message: "Aduan anda sedang dalam proses verifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Verifikasi",
				Message: "Aduan anda telah diverifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "On Progress",
				Message: "Aduan anda sedang dalam proses penanganan",
			})
		} else if status == "Selesai" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Pending",
				Message: "Aduan anda sedang dalam proses verifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Verifikasi",
				Message: "Aduan anda telah diverifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "On Progress",
				Message: "Aduan anda sedang dalam proses penanganan",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Selesai",
				Message: "Aduan anda telah selesai ditangani",
			})
		} else if status == "Ditolak" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Pending",
				Message: "Aduan anda sedang dalam proses verifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Ditolak",
				Message: "Aduan anda ditolak karena tidak sesuai dengan ketentuan yang berlaku",
			})
		} else if status == "Pending" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "Pending",
				Message: "Aduan anda sedang dalam proses verifikasi oleh admin kami",
			})
		}

		// split pathFiles by comma
		files := strings.Split(pathFiles, ",")
		complaintFiles = []entities.ComplaintFile{}
		for _, file := range files {
			complaintFile := entities.ComplaintFile{
				Path: file,
			}

			complaintFiles = append(complaintFiles, complaintFile)
		}

		complaint := entities.Complaint{
			ID:          utils.GenerateID("C-", 10),
			UserID:      userId,
			CategoryID:  categoryId,
			RegencyID:   regencyId,
			Address:     address,
			Description: description,
			Status:      status,
			Type:        typeComplaint,
			Date:        date,
			Files:       complaintFiles,
			Process:     process,
		}

		complaints = append(complaints, complaint)
	}

	// Import the complaints
	err = u.complaintRepo.Import(complaints)
	if err != nil {
		return err
	}

	return nil
}

func (u *ComplaintUseCase) IncreaseTotalLikes(id string) error {
	err := u.complaintRepo.IncreaseTotalLikes(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *ComplaintUseCase) DecreaseTotalLikes(id string) error {
	err := u.complaintRepo.DecreaseTotalLikes(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *ComplaintUseCase) GetComplaintIDsByUserID(userID int) ([]string, error) {
	complaintIDs, err := u.complaintRepo.GetComplaintIDsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return complaintIDs, nil
}
