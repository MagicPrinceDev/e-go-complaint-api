package complaint_process

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"strings"
)

type ComplaintProcessUseCase struct {
	repository          entities.ComplaintProcessRepositoryInterface
	complaintRepository entities.ComplaintRepositoryInterface
}

func NewComplaintProcessUseCase(repository entities.ComplaintProcessRepositoryInterface, complaintRepository entities.ComplaintRepositoryInterface) *ComplaintProcessUseCase {
	return &ComplaintProcessUseCase{
		repository:          repository,
		complaintRepository: complaintRepository,
	}
}

func (u *ComplaintProcessUseCase) Create(complaintProcess *entities.ComplaintProcess) (entities.ComplaintProcess, error) {
	if complaintProcess.Message == "" || complaintProcess.Status == "" {
		return entities.ComplaintProcess{}, constants.ErrAllFieldsMustBeFilled
	}

	if complaintProcess.Status != "Pending" && complaintProcess.Status != "Verifikasi" && complaintProcess.Status != "On Progress" && complaintProcess.Status != "Selesai" && complaintProcess.Status != "Ditolak" {
		return entities.ComplaintProcess{}, constants.ErrInvalidStatus
	}

	status, err := u.complaintRepository.GetStatus(complaintProcess.ComplaintID)
	if err != nil {
		return entities.ComplaintProcess{}, constants.ErrInternalServerError
	}

	if complaintProcess.Status == "Pending" {
		if status == "On Progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		} else if status == "Selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		}
	} else if complaintProcess.Status == "Verifikasi" {
		if status == "Verifikasi" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyVerified
		} else if status == "Ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "Selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "On Progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyOnProgress
		}
	} else if complaintProcess.Status == "On Progress" {
		if status == "On Progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyOnProgress
		} else if status == "Ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "Selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "Pending" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		}
	} else if complaintProcess.Status == "Selesai" {
		if status == "Selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "Ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "Pending" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		} else if status == "Verifikasi" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotOnProgress
		}
	} else if complaintProcess.Status == "Ditolak" {
		if status == "Ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "Selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "Verifikasi" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyVerified
		} else if status == "On Progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyOnProgress
		}
	}

	err = u.repository.Create(complaintProcess)
	if err != nil {
		if strings.Contains(err.Error(), "REFERENCES `complaints` (`id`)") {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotFound
		} else {
			return entities.ComplaintProcess{}, constants.ErrInternalServerError
		}
	}

	return *complaintProcess, nil
}

func (u *ComplaintProcessUseCase) GetByComplaintID(complaintID string) ([]entities.ComplaintProcess, error) {
	complaintProcesses, err := u.repository.GetByComplaintID(complaintID)
	if err != nil {
		return nil, err
	}

	return complaintProcesses, nil
}

func (u *ComplaintProcessUseCase) Update(complaintProcess *entities.ComplaintProcess) (entities.ComplaintProcess, error) {
	if complaintProcess.Message == "" {
		return entities.ComplaintProcess{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Update(complaintProcess)
	if err != nil {
		return entities.ComplaintProcess{}, err
	}

	return *complaintProcess, nil
}

func (u *ComplaintProcessUseCase) Delete(complaintID string, complaintProcessID int) (string, error) {
	if complaintID == "" || complaintProcessID == 0 {
		return "", constants.ErrInvalidIDFormat
	}

	status, err := u.repository.Delete(complaintID, complaintProcessID)
	if err != nil {
		return "", err
	}

	if status == "Verifikasi" {
		status = "Pending"
	} else if status == "On Progress" {
		status = "Verifikasi"
	} else if status == "Selesai" {
		status = "On Progress"
	} else if status == "Ditolak" {
		status = "Pending"
	}

	return status, nil
}
