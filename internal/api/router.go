package api

import (
	handlers "test/internal/api/handlers/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register(h handlers.CvHandler) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	v2 := r.Group("/api/v2")
	{
		v2.GET("/cv", h.GetCV)
	}

	r.NoRoute(h.Default)

	return r
}
