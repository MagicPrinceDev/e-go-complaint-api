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
