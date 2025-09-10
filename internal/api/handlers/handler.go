package handlers

import (
	"fmt"
	"net/http"
	"test/internal/domain"
	"test/internal/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store store.Store
}

func NewHandler(store store.Store) Handler {
	return Handler{store: store}
}

func (h Handler) GetProfile(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", domain.LangEn)

	raw, err := h.store.GetProfile(ctx.Request.Context(), lang)
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
