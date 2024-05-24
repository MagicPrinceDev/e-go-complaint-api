package regency

import (
	"e-complaint-api/entities"
	"encoding/json"
	"net/http"
)

type RegencyAPI struct {
	APIURL string
}

func NewRegencyAPI() *RegencyAPI {
	return &RegencyAPI{
		APIURL: "https://idn-area.up.railway.app/regencies?provinceCode=36&limit=100&sortBy=code",
	}
}

func (r *RegencyAPI) GetRegenciesDataFromAPI() ([]entities.Regency, error) {
	regencies := []entities.Regency{}
	response, err := http.Get(r.APIURL)
	if err != nil {
		return regencies, err
	}
	defer response.Body.Close()

	var dataResponse Regency
	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		return regencies, err
	}

	for _, reg := range dataResponse.Data {
		regencies = append(regencies, entities.Regency{ID: reg.Code, Name: reg.Name})
	}

	return regencies, nil
}
