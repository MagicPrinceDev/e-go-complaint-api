package village

import (
	"e-complaint-api/entities"
	"encoding/json"
	"net/http"
)

type VillageAPI struct {
	APIURL string
}

func NewVillageAPI() *VillageAPI {
	return &VillageAPI{
		APIURL: "https://idn-area.up.railway.app/villages?page=1&limit=100&sortBy=code",
	}
}

func (r *VillageAPI) GetVillagesDataFromAPI(district_ids []string) ([]entities.Village, error) {
	villages := []entities.Village{}
	for _, district_id := range district_ids {
		response, err := http.Get(r.APIURL + "&districtCode=" + district_id)
		if err != nil {
			return villages, err
		}
		defer response.Body.Close()

		var dataResponse Village
		err = json.NewDecoder(response.Body).Decode(&dataResponse)
		if err != nil {
			return villages, err
		}

		for _, reg := range dataResponse.Data {
			villages = append(villages, entities.Village{ID: reg.Code, Name: reg.Name, DistrictID: reg.DistrictCode})
		}
	}

	return villages, nil
}
