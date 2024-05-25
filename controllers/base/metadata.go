package base

type Pagination struct {
	TotalDataPerPage int `json:"total_data_per_page"`
	FirstPage        int `json:"first_page"`
	LastPage         int `json:"last_page"`
	CurrentPage      int `json:"current_page"`
	NextPage         int `json:"next_page"`
	PrevPage         int `json:"prev_page"`
}

type Metadata struct {
	TotalData  int `json:"total_data"`
	Pagination Pagination
}

func NewMetadata(totalData int, totalDataPerPage int, firstPage int, lastPage int, currentPage int, nextPage int, prevPage int) *Metadata {
	return &Metadata{
		TotalData: totalData,
		Pagination: Pagination{
			TotalDataPerPage: totalDataPerPage,
			FirstPage:        firstPage,
			LastPage:         lastPage,
			CurrentPage:      currentPage,
			NextPage:         nextPage,
			PrevPage:         prevPage,
		},
	}
}
