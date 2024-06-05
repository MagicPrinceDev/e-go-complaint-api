package regency

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
)

type RegencyUseCase struct {
	repository entities.RegencyRepositoryInterface
}

func NewRegencyUseCase(repo entities.RegencyRepositoryInterface) *RegencyUseCase {
	return &RegencyUseCase{repository: repo}
}

func (uc *RegencyUseCase) GetAll() ([]entities.Regency, error) {
	regencies, err := uc.repository.GetAll()
	if err != nil {
		return []entities.Regency{}, constants.ErrInternalServerError
	}

	return regencies, nil
}
