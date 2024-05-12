package main

import (
	"architecture/initializers"
	io "architecture/ws/services/IO"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	// r := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())
	io.BulkWriteInMemoryRoutes(router)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
