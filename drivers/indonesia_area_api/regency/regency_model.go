package regency

type Regency struct {
	Data []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"data"`
}
