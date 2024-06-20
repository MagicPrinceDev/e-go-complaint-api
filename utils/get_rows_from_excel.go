package utils

import (
	"e-complaint-api/constants"
	"io"
	"mime/multipart"
	"os"

	"github.com/xuri/excelize/v2"
)

func GetRowsFromExcel(file *multipart.FileHeader) ([][]string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, constants.ErrInternalServerError
	}
	defer f.Close()

	tempFile, err := os.CreateTemp("", "uploaded-*.xlsx")
	if err != nil {
		return nil, constants.ErrInternalServerError
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, f); err != nil {
		return nil, constants.ErrInternalServerError
	}

	excelFile, err := excelize.OpenFile(tempFile.Name())
	if err != nil {
		return nil, constants.ErrInternalServerError
	}
	defer excelFile.Close()

	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return rows, nil
}
