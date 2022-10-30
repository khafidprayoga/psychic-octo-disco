package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(msg error, internalErr int) fiber.Map {
	// For test
	status := "Not Found"
	var reason string
	if msg == nil {
		reason = "Internal Server Error [500]"
	} else {
		reason = fmt.Sprintf("%v [%d]", msg.Error(), internalErr)
	}

	// For test
	if internalErr == HTTPRequestErr {
		status = "Bad Request"
	}

	return fiber.Map{
		"status":  status,
		"success": false,
		"reason":  reason,
	}
}

func SuccessResponse(msg string, data any) fiber.Map {
	return fiber.Map{
		"status":  "Success",
		"success": true,
		"message": msg,
		"data":    data,
	}
}
