package request

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CaptureGinRequest(c *gin.Context) *Request {
	pageNumber := 1
	pageSize := 10

	page, _ := c.GetQueryMap("page")
	if p, ok := page["number"]; ok {
		intPageNumber, err := strconv.Atoi(p)
		if err == nil {
			pageNumber = intPageNumber
		}
	}
	if p, ok := page["size"]; ok {
		intPageSize, err := strconv.Atoi(p)
		if err == nil {
			pageSize = intPageSize
		}
	}

	sorts := map[string]string{}
	sortString := c.Query("sort")
	sortSlice := strings.Split(strings.TrimSpace(sortString), ",")
	for _, sort := range sortSlice {
		if len(sort) == 0 {
			continue
		}

		// if it has a "-" prefix, it means it is a descending sort
		if strings.HasPrefix(sort, "-") {
			sorts[strings.TrimPrefix(sort, "-")] = "desc"
		} else {
			sorts[sort] = "asc"
		}
	}

	search := c.Query("search")

	filters := map[string]string{}
	filterMap := c.QueryMap("filter")
	for key, value := range filterMap {
		if len(value) > 0 {
			filters[key] = value
		}
	}

	return &Request{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		Sorts:      sorts,
		Filters:    filters,
		Search:     search,
	}
}
