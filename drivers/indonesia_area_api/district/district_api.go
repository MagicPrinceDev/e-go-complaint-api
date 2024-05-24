package district

import (
	"e-complaint-api/entities"
	"encoding/json"
	"net/http"
)

type DistrictAPI struct {
	APIURL string
}

func NewDistrictAPI() *DistrictAPI {
	return &DistrictAPI{
		APIURL: "https://idn-area.up.railway.app/districts?page=1&limit=100&sortBy=code",
	}
}

func (r *DistrictAPI) GetDistrictsDataFromAPI(regencies_id []string) ([]entities.District, error) {
	districts := []entities.District{}
	for _, regency_id := range regencies_id {
		response, err := http.Get(r.APIURL + "&regencyCode=" + regency_id)
		if err != nil {
			return districts, err
		}
		defer response.Body.Close()

		var dataResponse District
		err = json.NewDecoder(response.Body).Decode(&dataResponse)
		if err != nil {
			return districts, err
		}

		for _, reg := range dataResponse.Data {
			districts = append(districts, entities.District{ID: reg.Code, Name: reg.Name, RegencyID: reg.RegencyCode})
		}
	}

	return districts, nil
}
