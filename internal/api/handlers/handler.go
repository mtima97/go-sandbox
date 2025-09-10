package handlers

import (
	"fmt"
	"net/http"
	"test/internal/domain"
	"test/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	servc service.Cv
}

func NewHandler(servc service.Cv) Handler {
	return Handler{servc: servc}
}

func (h Handler) GetProfile(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", domain.LangEn)

	raw, err := h.servc.Profile(ctx, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, genError(err))
		return
	}

	ctx.JSON(http.StatusOK, response(raw))
}

func (h Handler) NoRoute(ctx *gin.Context) {
	err := fmt.Errorf("no route")

	ctx.JSON(http.StatusNotFound, genError(err))
}

func genError(err error) gin.H {
	return gin.H{
		"error": fmt.Sprintf("error: %v", err),
	}
}

func response(data any) gin.H {
	return gin.H{
		"data": data,
	}
}
