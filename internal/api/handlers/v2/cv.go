package v2

import (
	"context"
	"net/http"
	"test/internal/domain"
	"test/internal/models/responses"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetCV(ctx context.Context, language string) (responses.CV, error)
}

type CvHandler struct {
	svc Service
}

func NewCvHandler(svc Service) CvHandler {
	return CvHandler{svc: svc}
}

func (h CvHandler) GetCV(ctx *gin.Context) {
	language := ctx.DefaultQuery("lang", domain.LangEn)

	page, err := h.svc.GetCV(ctx.Request.Context(), language)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": page})
}

func (h CvHandler) Default(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}
