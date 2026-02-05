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
		return RespondError(c, http.StatusUnprocessableEntity, "invalid request body")
	}

	if err := r.DB.Create(&book).Error; err != nil {
		return RespondError(c, http.StatusInternalServerError, "could not create book")
	}

	return RespondSuccess(c, http.StatusCreated, "book created successfully", book)
}

// GET ALL BOOKS
func (r *Repository) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if err := r.DB.Find(&books).Error; err != nil {
		return RespondError(c, http.StatusInternalServerError, "failed to fetch books")
	}

	return RespondSuccess(c, http.StatusOK, "", books)
}

// GET BOOK BY ID
func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return RespondError(c, http.StatusBadRequest, "id is required")
	}

	book := models.Book{}
	if err := r.DB.First(&book, id).Error; err != nil {
		return RespondError(c, http.StatusNotFound, "book not found")
	}

	return RespondSuccess(c, http.StatusOK, "", book)
}

// DELETE BOOK
func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return RespondError(c, http.StatusBadRequest, "id is required")
	}

	result := r.DB.Delete(&models.Book{}, id)
	if result.Error != nil {
		return RespondError(c, http.StatusInternalServerError, "could not delete book")
	}

	return RespondSuccess(c, http.StatusOK, "book deleted successfully", nil)
}
