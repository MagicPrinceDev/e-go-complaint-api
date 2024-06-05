package news_file

import (
	"e-complaint-api/entities"
	"mime/multipart"
)

type NewsFileUseCase struct {
	repository entities.NewsFileRepositoryInterface
	gcs_api    entities.NewsFileGCSAPIInterface
}

func NewNewsFileUseCase(repository entities.NewsFileRepositoryInterface, gcs_api entities.NewsFileGCSAPIInterface) *NewsFileUseCase {
	return &NewsFileUseCase{
		repository: repository,
		gcs_api:    gcs_api,
	}
}

func (u *NewsFileUseCase) Create(files []*multipart.FileHeader, newsID int) ([]entities.NewsFile, error) {
	filepaths, err_upload := u.gcs_api.Upload(files)
	if err_upload != nil {
		return []entities.NewsFile{}, err_upload
	}

	var newsFiles []*entities.NewsFile
	for _, filepath := range filepaths {
		newsFile := &entities.NewsFile{
			NewsID: newsID,
			Path:   filepath,
		}
		newsFiles = append(newsFiles, newsFile)
	}

	err_create := u.repository.Create(newsFiles)
	if err_create != nil {
		return []entities.NewsFile{}, err_create
	}

	var convertedNewsFiles []entities.NewsFile
	for _, nf := range newsFiles {
		convertedNewsFiles = append(convertedNewsFiles, *nf)
	}

	return convertedNewsFiles, nil
}

func (u *NewsFileUseCase) DeleteByNewsID(newsID int) error {
	err_delete := u.repository.DeleteByNewsID(newsID)
	if err_delete != nil {
		return err_delete
	}

	return nil
}
