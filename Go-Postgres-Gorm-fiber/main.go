package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/ishansaini194/Projects/handlers"
	"github.com/ishansaini194/Projects/models"
	"github.com/ishansaini194/Projects/routes"
	"github.com/ishansaini194/Projects/storage"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file")
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not connect to database")
	}

	if err := models.MigrateBooks(db); err != nil {
		log.Fatal("migration failed")
	}

	repo := &handlers.Repository{
		DB: db,
	}

	app := fiber.New()

	routes.SetupRoutes(app, repo)

	log.Fatal(app.Listen(":8080"))
}
