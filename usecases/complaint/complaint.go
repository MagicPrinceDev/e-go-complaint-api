package complaint

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"e-complaint-api/utils"
	"strings"
)

type ComplaintUseCase struct {
	repository entities.ComplaintRepositoryInterface
}

func NewComplaintUseCase(repository entities.ComplaintRepositoryInterface) *ComplaintUseCase {
	return &ComplaintUseCase{
		repository: repository,
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

	complaints, err := u.repository.GetPaginated(limit, page, search, filter, sortBy, sortType)
	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return complaints, nil
}

func (u *ComplaintUseCase) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var pagination entities.Pagination
	metaData, err := u.repository.GetMetaData(limit, page, search, filter)

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
	complaint, err := u.repository.GetByID(id)
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

	err := u.repository.Create(complaint)
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
		err := u.repository.AdminDelete(id)
		if err != nil {
			return err
		}
	} else {
		err := u.repository.Delete(id, userId)
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

	complaint, err := u.repository.Update(complaint)
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
