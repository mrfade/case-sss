package request

type Pagination struct {
	PageNumber   int  `json:"pageNumber"`
	PageSize     int  `json:"pageSize"`
	TotalPages   int  `json:"totalPages"`
	TotalRecords int  `json:"totalRecords"`
	HasNext      bool `json:"hasNext"`
	HasPrev      bool `json:"hasPrev"`
}

type Meta struct {
	Searchable []string `json:"searchable,omitempty"`
	Filterable []string `json:"filterable,omitempty"`
	Sortable   []string `json:"sortable,omitempty"`
}

type Request struct {
	PageNumber int
	PageSize   int
	Sorts      map[string]string
	Filters    map[string]string
	Search     string
}
