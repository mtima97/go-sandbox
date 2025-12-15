package api

import (
	handlers "test/internal/api/handlers/v2"
	"test/internal/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register(h handlers.CvHandler, c config.Cfg) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: c.AllowedOrigins,
		AllowMethods: []string{"GET", "OPTIONS"},
		AllowHeaders: []string{"Accept", "Content-Type"},
		MaxAge:       12 * time.Hour,
	}))

	v2 := r.Group("/api/v2")
	{
		v2.GET("/cv", h.GetCV)
	}

	r.NoRoute(h.Default)

	return r
}
