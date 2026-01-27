package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ishansaini194/Projects/database"
	"github.com/ishansaini194/Projects/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("lead.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	fmt.Println("Database Connected")

	err = database.DBConn.AutoMigrate(&lead.Lead{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
