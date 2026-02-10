package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ishansaini/Projects/routes"
	"github.com/joho/godotenv"
)

func main() {

	//Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//Validate required env vars
	if os.Getenv("SECRET_KEY") == "" {
		log.Fatal("SECRET_KEY not set")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access grante for api-2"})
	})

	router.Run(":" + port)
}
