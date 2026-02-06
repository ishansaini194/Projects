package errors

import "github.com/gofiber/fiber/v2"

type APIError struct {
	Status  int
	Message string
}

func (e APIError) Error() string {
	return e.Message
}

func New(status int, msg string) APIError {
	return APIError{
		Status:  status,
		Message: msg,
	}
}

func Handle(c *fiber.Ctx, err error) error {
	if apiErr, ok := err.(APIError); ok {
		return c.Status(apiErr.Status).JSON(fiber.Map{
			"error": apiErr.Message,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error": "Internal server error",
	})
}
