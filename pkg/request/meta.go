package request

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
