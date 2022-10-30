package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(msg error, internalErr error) fiber.Map {
	if msg == nil {
		msg = errors.New("Internal Server Error")
	}
	return fiber.Map{
		"status":  "error",
		"success": false,
		"reason":  fmt.Sprintf("%v %v", msg.Error(), internalErr.Error()),
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
