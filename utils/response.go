package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(msg error, internalErr int) fiber.Map {
	var reason string
	if msg == nil {
		reason = "Internal Server Error [500]"
	} else {
		reason = fmt.Sprintf("%v [%d]", msg.Error(), internalErr)
	}

	return fiber.Map{
		"status":  "error",
		"success": false,
		"reason":  reason,
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
