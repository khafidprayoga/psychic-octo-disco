package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponse(msg string) fiber.Map {
	if msg == "" {
		msg = "Internal Server Error"
	}
	return fiber.Map{
		"status":  "error",
		"success": false,
		"reason":  msg,
		"data":    nil,
	}
}

func SuccessResponse(msg string, data any) fiber.Map {
	return fiber.Map{
		"status":  "success",
		"success": true,
		"message": msg,
		"data":    data,
	}
}
