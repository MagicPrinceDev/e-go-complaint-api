package complaint_process

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
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

	if complaintProcess.Status != "verifikasi" && complaintProcess.Status != "on progress" && complaintProcess.Status != "selesai" && complaintProcess.Status != "ditolak" {
		return entities.ComplaintProcess{}, constants.ErrInvalidStatus
	}

	status, err := u.complaintRepository.GetStatus(complaintProcess.ComplaintID)
	if err != nil {
		return entities.ComplaintProcess{}, err
	}

	if complaintProcess.Status == "pending" {
		if status == "on progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		} else if complaintProcess.Status == "selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		}
	} else if complaintProcess.Status == "verifikasi" {
		if status == "verifikasi" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyVerified
		} else if status == "ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "on progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyOnProgress
		}
	} else if complaintProcess.Status == "on progress" {
		if status == "on progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyOnProgress
		} else if status == "ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "pending" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		}
	} else if complaintProcess.Status == "selesai" {
		if status == "selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "pending" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotVerified
		} else if status == "verifikasi" {
			return entities.ComplaintProcess{}, constants.ErrComplaintNotOnProgress
		}
	} else if complaintProcess.Status == "ditolak" {
		if status == "ditolak" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyRejected
		} else if status == "selesai" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyFinished
		} else if status == "verifikasi" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyVerified
		} else if status == "on progress" {
			return entities.ComplaintProcess{}, constants.ErrComplaintAlreadyOnProgress
		}
	}

	err = u.repository.Create(complaintProcess)
	if err != nil {
		return entities.ComplaintProcess{}, err
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

	if status == "verifikasi" {
		status = "pending"
	} else if status == "on progress" {
		status = "verifikasi"
	} else if status == "selesai" {
		status = "on progress"
	} else if status == "ditolak" {
		status = "pending"
	}

	return status, nil
}
