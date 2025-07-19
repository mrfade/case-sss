package request

import "slices"

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

func FilterUnsupportedFields(request *Request, meta *Meta) {
	if request == nil || meta == nil {
		return
	}

	for field := range request.Filters {
		if !slices.Contains(meta.Filterable, field) {
			delete(request.Filters, field)
		}
	}

	for field := range request.Sorts {
		if !slices.Contains(meta.Sortable, field) {
			delete(request.Sorts, field)
		}
	}
}
