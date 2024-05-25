package google_cloud_storage

import (
	"context"
	"e-complaint-api/constants"
	"e-complaint-api/utils"
	"io"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const (
	projectID  = "alterra-academy-420809"
	bucketName = "e-complaint-assets"
)

type FileHandlingAPI struct {
	APIKey     string
	FolderPath string
}

func NewFileHandlingAPI(APIKey string, folderPath string) *FileHandlingAPI {
	return &FileHandlingAPI{
		APIKey:     APIKey,
		FolderPath: folderPath,
	}
}

func (f *FileHandlingAPI) Upload(files []*multipart.FileHeader) ([]string, error) {
	var credentials = f.APIKey

	var filePaths []string
	for _, fileHeader := range files {
		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Load GCP credentials securely (consider using KMS or secrets manager)
		ctx := context.Background()
		client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(credentials)))
		if err != nil {
			return nil, err
		}
		defer client.Close()

		// Hashing nama file menggunakan SHA256
		hashedFilename := utils.HashFileName(fileHeader.Filename)

		dstPath := f.FolderPath + hashedFilename // Menggunakan nama file yang dihash
		dst := client.Bucket(bucketName).Object(dstPath).NewWriter(ctx)
		filePaths = append(filePaths, dstPath)
		defer dst.Close()

		// Salin isi file dari source ke destination di GCS
		if _, err = io.Copy(dst, file); err != nil {
			return nil, err
		}
	}

	return filePaths, nil
}

func (f *FileHandlingAPI) Delete(filePaths []string) error {
	var credentials = os.Getenv("GCS_CREDENTIALS")

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(credentials)))
	if err != nil {
		return constants.ErrFailedToCreateClientGCS
	}
	defer client.Close()

	for _, path := range filePaths {
		obj := client.Bucket(bucketName).Object(path)
		if err := obj.Delete(ctx); err != nil {
			return constants.ErrFailedToDeleteObject
		}
	}

	return nil
}
