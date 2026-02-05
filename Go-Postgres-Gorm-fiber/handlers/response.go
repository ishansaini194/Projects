package handlers

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondError(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(ErrorResponse{
		Error: msg,
	})
}

func RespondSuccess(c *fiber.Ctx, status int, msg string, data interface{}) error {
	return c.Status(status).JSON(SuccessResponse{
		Message: msg,
		Data:    data,
	})
}
