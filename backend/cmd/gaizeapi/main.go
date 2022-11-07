package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gaizeapi/api/handler"
	"gaizeapi/internal/config"
)

func main() {
	// Ensure necessary directories exist
	config.EnsureDirectories()

	r := gin.Default()

	// CORS setup to allow requests from the frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5174"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	r.POST("/data", handler.ReceiveVideo)
	r.GET("/result", handler.SendVideo)

	r.Run(":8080")
}
