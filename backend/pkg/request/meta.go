package request

import (
	"fmt"
	"slices"
)

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
	Searchs    map[string]string
}

func (r *Request) String() string {
	return fmt.Sprintf("%d-%d-%v-%v-%v", r.PageNumber, r.PageSize, r.Sorts, r.Filters, r.Searchs)
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

	for field := range request.Searchs {
		if !slices.Contains(meta.Searchable, field) {
			delete(request.Searchs, field)
		}
	}
}
