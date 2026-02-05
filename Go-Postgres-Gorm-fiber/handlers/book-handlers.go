package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/ishansaini194/Projects/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// CREATE BOOK
func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := models.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{"message": "invalid request body"})
	}

	if err := r.DB.Create(&book).Error; err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{"message": "could not create book"})
	}

	return c.Status(http.StatusCreated).
		JSON(fiber.Map{"message": "book created successfully", "data": book})
}

// GET ALL BOOKS
func (r *Repository) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if err := r.DB.Find(&books).Error; err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{"message": "could not fetch books"})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{"data": books})
}

// GET BOOK BY ID
func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{"message": "id is required"})
	}

	book := models.Book{}
	if err := r.DB.First(&book, id).Error; err != nil {
		return c.Status(http.StatusNotFound).
			JSON(fiber.Map{"message": "book not found"})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{"data": book})
}

// DELETE BOOK
func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{"message": "id is required"})
	}

	result := r.DB.Delete(&models.Book{}, id)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{"message": "could not delete book"})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{"message": "book deleted successfully"})
}
