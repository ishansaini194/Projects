package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishansaini194/Projects/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn // database connection
	var leads []Lead      // slice to store results

	result := db.Find(&leads) // execute query

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	db := database.DBConn // db connection
	id := c.Params("id")  // get id from URL

	var lead Lead // struct to store record

	result := db.First(&lead, id) // SELECT * FROM leads WHERE id = ?

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Lead not found",
		})
	}

	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	var lead Lead

	// Parse request body into lead struct
	parseErr := c.BodyParser(&lead)
	if parseErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Insert record into database
	createResult := db.Create(&lead)
	if createResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": createResult.Error.Error(),
		})
	}

	return c.Status(201).JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	var lead Lead

	// Check if lead exists
	findResult := db.First(&lead, id)
	if findResult.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Lead not found",
		})
	}

	// Delete the lead
	deleteResult := db.Delete(&lead)
	if deleteResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": deleteResult.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Lead deleted successfully",
	})
}
