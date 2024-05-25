package complaint_file

import (
	"e-complaint-api/entities"
	"mime/multipart"
)

type ComplaintFileUseCase struct {
	repository entities.ComplaintFileRepositoryInterface
	gcs_api    entities.ComplaintFileGCSAPIInterface
}

func NewComplaintFileUseCase(repository entities.ComplaintFileRepositoryInterface, gcs_api entities.ComplaintFileGCSAPIInterface) *ComplaintFileUseCase {
	return &ComplaintFileUseCase{
		repository: repository,
		gcs_api:    gcs_api,
	}
}

func (u *ComplaintFileUseCase) Create(files []*multipart.FileHeader, complaintID string) ([]entities.ComplaintFile, error) {
	filepaths, err_upload := u.gcs_api.Upload(files)
	if err_upload != nil {
		return []entities.ComplaintFile{}, err_upload
	}

	var complaintFiles []*entities.ComplaintFile
	for _, filepath := range filepaths {
		complaintFile := &entities.ComplaintFile{
			ComplaintID: complaintID,
			Path:        filepath,
		}
		complaintFiles = append(complaintFiles, complaintFile)
	}

	err_create := u.repository.Create(complaintFiles)
	if err_create != nil {
		return []entities.ComplaintFile{}, err_create
	}

	var convertedComplaintFiles []entities.ComplaintFile
	for _, cf := range complaintFiles {
		convertedComplaintFiles = append(convertedComplaintFiles, *cf)
	}

	return convertedComplaintFiles, nil
}
