package complaint_process

import "e-complaint-api/entities"

type ComplaintProcessUseCase struct {
	repository entities.ComplaintProcessRepositoryInterface
}

func NewComplaintProcessUseCase(repository entities.ComplaintProcessRepositoryInterface) *ComplaintProcessUseCase {
	return &ComplaintProcessUseCase{
		repository: repository,
	}
}

func (u *ComplaintProcessUseCase) Create(complaintProcess *entities.ComplaintProcess) (entities.ComplaintProcess, error) {
	err := u.repository.Create(complaintProcess)
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
