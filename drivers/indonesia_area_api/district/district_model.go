package district

type District struct {
	Data []struct {
		Code        string `json:"code"`
		Name        string `json:"name"`
		RegencyCode string `json:"regencyCode"`
	} `json:"data"`
}
