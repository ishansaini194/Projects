package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishansaini194/Projects/handlers"
)

func SetupRoutes(app *fiber.App, r *handlers.Repository) {
	api := app.Group("/api")

	api.Post("/books", r.CreateBook)
	api.Get("/books", r.GetBooks)
	api.Get("/books/:id", r.GetBookByID)
	api.Delete("/books/:id", r.DeleteBook)
}
