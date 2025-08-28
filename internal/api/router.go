package api

import (
	"log"
	"test/internal/api/handlers"
	"test/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(cv service.Cv, addr string) {
	r := gin.Default()
	h := handlers.NewHandler(cv)

	r.Use(cors.Default())

	api := r.Group("/api")

	{
		api.GET("/profile", h.GetProfile)
		api.GET("/experience", h.GetExperience)
		api.GET("/education", h.GetEducation)
		api.GET("/languages", h.GetLanguages)
		api.GET("/projects", h.GetProjects)
	}

	r.NoRoute(h.NoRoute)

	if err := r.Run(addr); err != nil {
		log.Fatalf("cannot start api server: %v", err)
	}
}
