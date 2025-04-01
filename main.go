package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yourusername/license-api/internal/api"

	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.StaticFile("/docs/openapi.json", "./docs/openapi.json")

	r.Static("/swagger-ui", "./swagger-ui/dist")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger-ui/")
	})

	handler := &api.HandlerImpl{}
	api.RegisterHandlersWithOptions(r, handler, api.GinServerOptions{
		BaseURL: "/api",
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
