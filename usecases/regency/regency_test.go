package regency

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RegencyRepositoryMock struct {
	mock.Mock
}

func (m *RegencyRepositoryMock) GetAll() ([]entities.Regency, error) {
	args := m.Called()
	return args.Get(0).([]entities.Regency), args.Error(1)
}

func TestRegencyUseCase_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := new(RegencyRepositoryMock)
		uc := NewRegencyUseCase(repo)

		regencies := []entities.Regency{
			{
				ID:   "1901",
				Name: "Kabupaten Bogor",
			},
			{
				ID:   "1902",
				Name: "Kabupaten Bekasi",
			},
		}

		repo.On("GetAll").Return(regencies, nil)

		result, err := uc.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, regencies, result)
	})

	t.Run("failed internal server error", func(t *testing.T) {
		repo := new(RegencyRepositoryMock)
		uc := NewRegencyUseCase(repo)

		repo.On("GetAll").Return([]entities.Regency{}, errors.New("unexpected error"))

		result, err := uc.GetAll()

		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Empty(t, result)
	})
}
