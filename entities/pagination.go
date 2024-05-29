package entities

type Pagination struct {
	TotalDataPerPage int
	FirstPage        int
	LastPage         int
	CurrentPage      int
	NextPage         int
	PrevPage         int
}

type Metadata struct {
	TotalData  int
	Pagination Pagination
}
