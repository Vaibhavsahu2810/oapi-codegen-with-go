package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/license-api/internal/api"
)

func main() {
	router := gin.Default()

	_, err := api.GetSwagger()
	if err != nil {
		panic("Failed to load OpenAPI spec")
	}
	api.RegisterHandlersWithOptions(router, &api.HandlerImpl{}, api.GinServerOptions{
		BaseURL: "/api",
	})

	router.Run(":8080")
}
