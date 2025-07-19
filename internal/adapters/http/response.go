package http

import (
	"math"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mrfade/case-sss/pkg/request"
)

type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func Success(ctx *gin.Context, data any) {
	value := reflect.ValueOf(data)

	if value.Kind() == reflect.Slice && value.IsNil() {
		// if the data is a nil slice, set it to an empty slice
		data = []any{}
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

type Pagination struct {
	PageNumber   int  `json:"pageNumber"`
	PageSize     int  `json:"pageSize"`
	TotalPages   int  `json:"totalPages"`
	TotalRecords int  `json:"totalRecords"`
	HasNext      bool `json:"hasNext"`
	HasPrev      bool `json:"hasPrev"`
}

func NewPagination(totalRecords int64, pageNumber int, pageSize int) *Pagination {
	// calculate the total pages
	totalPages := math.Ceil(float64(totalRecords) / float64(pageSize))

	// return the paginated response
	pagination := &Pagination{
		PageNumber:   pageNumber,
		PageSize:     pageSize,
		TotalPages:   int(totalPages),
		TotalRecords: int(totalRecords),
		HasNext:      pageNumber < int(totalPages),
		HasPrev:      pageNumber > 1,
	}

	return pagination
}

type PaginatedResponse struct {
	Meta       *request.Meta `json:"meta"`
	Pagination *Pagination   `json:"pagination"`
	Items      any           `json:"items"`
}

func NewPaginatedResponse(data any, meta *request.Meta, pagination *Pagination) *PaginatedResponse {
	value := reflect.ValueOf(data)

	if value.Kind() == reflect.Slice && value.IsNil() {
		// if the data is a nil slice, set it to an empty slice
		data = []any{}
	}

	return &PaginatedResponse{
		Meta:       meta,
		Pagination: pagination,
		Items:      data,
	}
}
