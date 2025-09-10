package api

import (
	"test/internal/api/handlers"
	"test/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(store store.Store) *gin.Engine {
	r := gin.Default()
	h := handlers.NewHandler(store)

	api := r.Group("/api")

	{
		api.GET("/profile", h.GetProfile)
	}

	r.NoRoute(h.NoRoute)

	return r
}
