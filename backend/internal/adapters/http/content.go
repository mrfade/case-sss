package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrfade/case-sss/internal/core/ports"
	"github.com/mrfade/case-sss/pkg/request"
)

type ContentHandler struct {
	service ports.ContentService
}

func NewContentHandler(service ports.ContentService) *ContentHandler {
	return &ContentHandler{
		service,
	}
}

func (handler *ContentHandler) FindAll(ctx *gin.Context) {
	meta := &request.Meta{
		Filterable: []string{"type"},
		Sortable:   []string{"score"},
		Searchable: []string{"title"},
	}

	req := request.CaptureGinRequest(ctx)
	request.FilterUnsupportedFields(req, meta)

	contents, totalRecords, err := handler.service.FindAll(ctx, req)
	if err != nil {
		Error(ctx, err)
		return
	}

	pagination := NewPagination(totalRecords, req.PageNumber, req.PageSize)
	Success(ctx, NewPaginatedResponse(contents, meta, pagination))
}
