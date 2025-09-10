package api

import (
	"test/internal/api/handlers"
	"test/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(cv service.Cv) *gin.Engine {
	r := gin.Default()
	h := handlers.NewHandler(cv)

	r.Use(cors.Default())

	api := r.Group("/api")

	{
		api.GET("/profile", h.GetProfile)
	}

	r.NoRoute(h.NoRoute)

	return r
}
