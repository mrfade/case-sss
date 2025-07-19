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
	req := request.CaptureGinRequest(ctx)

	meta := request.Meta{
		Filterable: []string{"type"},
		Sortable:   []string{"score"},
	}

	contents, totalRecords, err := handler.service.FindAll(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	pagination := NewPagination(totalRecords, req.PageNumber, req.PageSize)
	Success(ctx, NewPaginatedResponse(contents, meta, pagination))
}
