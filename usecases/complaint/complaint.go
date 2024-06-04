package complaint

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ComplaintUseCase struct {
	complaintRepo     entities.ComplaintRepositoryInterface
	complaintFileRepo entities.ComplaintFileRepositoryInterface
}

func NewComplaintUseCase(complaintRepo entities.ComplaintRepositoryInterface, complaintFileRepo entities.ComplaintFileRepositoryInterface) *ComplaintUseCase {
	return &ComplaintUseCase{
		complaintRepo:     complaintRepo,
		complaintFileRepo: complaintFileRepo,
	}
}

func (u *ComplaintUseCase) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	if limit == 0 || page == 0 {
		return nil, constants.ErrLimitAndPageMustBeFilled
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

	pagination.FirstPage = 1
	pagination.LastPage = (metaData.TotalData + limit - 1) / limit
	pagination.CurrentPage = page
	if pagination.CurrentPage == pagination.LastPage {
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

func (u *ComplaintUseCase) Create(complaint *entities.Complaint) (entities.Complaint, error) {
	if complaint.CategoryID == 0 || complaint.UserID == 0 || complaint.RegencyID == "" || complaint.Description == "" || complaint.Address == "" || complaint.Type == "" {
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
	if complaint.CategoryID == 0 || complaint.UserID == 0 || complaint.RegencyID == "" || complaint.Description == "" || complaint.Address == "" || complaint.Type == "" {
		return entities.Complaint{}, constants.ErrAllFieldsMustBeFilled
	}

	complaint, err := u.complaintRepo.Update(complaint)
	if err != nil {
		if strings.HasSuffix(err.Error(), "REFERENCES `regencies` (`id`))") {
			return entities.Complaint{}, constants.ErrRegencyNotFound
		} else if strings.HasSuffix(err.Error(), "REFERENCES `categories` (`id`))") {
			return entities.Complaint{}, constants.ErrCategoryNotFound
		} else {
			return entities.Complaint{}, constants.ErrInternalServerError
		}
	}

	return complaint, nil
}

func (u *ComplaintUseCase) UpdateStatus(id string, status string) error {
	if status != "pending" && status != "verifikasi" && status != "on progress" && status != "selesai" && status != "ditolak" {
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
	// Open the file from the multipart.FileHeader
	f, err := file.Open()
	if err != nil {
		return constants.ErrInternalServerError
	}
	defer f.Close()

	// Create a temporary file to copy the contents
	tempFile, err := ioutil.TempFile("", "uploaded-*.xlsx")
	if err != nil {
		return constants.ErrInternalServerError
	}
	defer os.Remove(tempFile.Name()) // Clean up the temp file
	defer tempFile.Close()

	// Copy the file contents to the temporary file
	if _, err := io.Copy(tempFile, f); err != nil {
		return constants.ErrInternalServerError
	}

	// Open the temporary file using excelize
	excelFile, err := excelize.OpenFile(tempFile.Name())
	if err != nil {
		return constants.ErrInternalServerError
	}
	defer excelFile.Close()

	// Get all rows in the "Sheet1"
	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		return constants.ErrInternalServerError
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

		if len(row) < 6 {
			// Ensure the row has enough columns
			return constants.ErrInternalServerError
		}

		userId, err := strconv.Atoi(row[0])
		if err != nil {
			return constants.ErrInternalServerError
		}

		categoryId, err := strconv.Atoi(row[1])
		if err != nil {
			return constants.ErrInternalServerError
		}

		regencyId := row[2]
		address := row[3]
		description := row[4]
		status := row[5]
		typeComplaint := row[6]
		pathFiles := row[7]

		process = []entities.ComplaintProcess{}
		if status == "verifikasi" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "verifikasi",
				Message: "Aduan anda telah diverifikasi oleh admin kami",
			})
		} else if status == "on progress" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "verifikasi",
				Message: "Aduan anda telah diverifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "on progress",
				Message: "Aduan anda sedang dalam proses penanganan",
			})
		} else if status == "selesai" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "verifikasi",
				Message: "Aduan anda telah diverifikasi oleh admin kami",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "on progress",
				Message: "Aduan anda sedang dalam proses penanganan",
			})
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "selesai",
				Message: "Aduan anda telah selesai ditangani",
			})
		} else if status == "ditolak" {
			process = append(process, entities.ComplaintProcess{
				AdminID: 1,
				Status:  "ditolak",
				Message: "Aduan anda ditolak karena tidak sesuai dengan ketentuan yang berlaku",
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
