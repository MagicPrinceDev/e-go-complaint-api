package village

type Village struct {
	Data []struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		DistrictCode string `json:"districtCode"`
	} `json:"data"`
}
