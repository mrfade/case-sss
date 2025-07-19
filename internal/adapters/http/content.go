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
	request := request.CaptureGinRequest(ctx)

	contents, err := handler.service.FindAll(ctx, request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, contents)
}
