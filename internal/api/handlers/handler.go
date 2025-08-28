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

	profile, err := h.servc.Profile(ctx, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, genError(err))
		return
	}

	ctx.JSON(http.StatusOK, response(profile))
}

func (h Handler) GetExperience(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", domain.LangEn)

	res, err := h.servc.Experience(ctx, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, genError(err))
		return
	}

	ctx.JSON(http.StatusOK, response(res))
}

func (h Handler) GetEducation(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", domain.LangEn)

	res, err := h.servc.Education(ctx, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, genError(err))
		return
	}

	ctx.JSON(http.StatusOK, response(res))
}

func (h Handler) GetLanguages(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", domain.LangEn)

	res, err := h.servc.Languages(ctx, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, genError(err))
		return
	}

	ctx.JSON(http.StatusOK, response(res))
}

func (h Handler) GetProjects(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", domain.LangEn)

	res, err := h.servc.Projects(ctx, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, genError(err))
		return
	}

	ctx.JSON(http.StatusOK, response(res))
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
