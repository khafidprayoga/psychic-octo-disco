package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetAppRouteList(app *fiber.App) string {
	var metadata strings.Builder
	listRoutes := app.GetRoutes(true)

	for _, endpoint := range listRoutes {
		if endpoint.Name != "" {
			routeData := fmt.Sprintf(
				"%v %v  => %v. handled on memory address %p\n",
				endpoint.Method,
				endpoint.Path,
				endpoint.Name,
				endpoint.Handlers,
			)
			metadata.WriteString(routeData)
		}
	}

	return metadata.String()
}
